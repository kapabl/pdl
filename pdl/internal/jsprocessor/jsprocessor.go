package jsprocessor

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/kapablanka/pdl/pdl/internal/config"
	"github.com/kapablanka/pdl/pdl/internal/templates"
	"github.com/kapablanka/pdl/pdl/internal/utils"
)

type Processor struct {
	cfg            config.RootConfig
	jsCfg          config.JSConfig
	printer        utils.VerbosePrinter
	engine         *templates.Engine
	namespaceTree  map[string]map[string]interface{}
	usedNamespaces map[string]bool
	globalIndex    []GlobalIndexClass
}

func Run(cfg config.RootConfig, printer utils.VerbosePrinter) error {
	baseData := map[string]interface{}{
		"companyName": cfg.CompanyName,
		"project":     cfg.Project,
		"version":     cfg.Version,
	}
	engine := templates.NewEngine(baseData)
	processor := Processor{
		cfg:            cfg,
		jsCfg:          cfg.JS,
		printer:        printer,
		engine:         engine,
		namespaceTree:  make(map[string]map[string]interface{}),
		usedNamespaces: make(map[string]bool),
		globalIndex:    make([]GlobalIndexClass, 0),
	}
	return processor.processAll()
}

func (processor *Processor) processAll() error {
	dirs := processor.jsCfg.Dirs
	for _, dir := range dirs {
		abs, absErr := filepath.Abs(dir)
		if absErr != nil {
			return absErr
		}
		if _, statErr := os.Stat(abs); statErr != nil {
			continue
		}
		classes, err := processor.processDir(abs, abs)
		if err != nil {
			return err
		}
		if len(classes) == 0 {
			continue
		}
		if processor.jsCfg.Namespaces.Enabled {
			for _, classData := range classes {
				processor.addNamespace(classData.ClassNode.Namespace)
			}
		}
		if processor.shouldGenerateGlobalIndex() {
			alias := processor.allocateGlobalNamespace(classes[0].ClassNode.Namespace)
			entry := GlobalIndexClass{Namespace: alias, Classes: classes}
			processor.globalIndex = append(processor.globalIndex, entry)
		}
	}
	if processor.jsCfg.Namespaces.Enabled {
		if err := processor.generateNamespaceFiles(); err != nil {
			return err
		}
	}
	if processor.shouldGenerateGlobalIndex() {
		if err := processor.generateGlobalIndexFile(); err != nil {
			return err
		}
	}
	return nil
}

func (processor *Processor) processDir(root string, dir string) ([]ClassData, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	classes := make([]ClassData, 0)
	subDirs := make([]string, 0)
	for _, entry := range entries {
		entryPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			subDirs = append(subDirs, entryPath)
			continue
		}
		if strings.HasSuffix(strings.ToLower(entry.Name()), ".js") && strings.ToLower(entry.Name()) != "index.js" {
			classData, err := parseClass(root, entryPath)
			if err != nil {
				return nil, err
			}
			if classData.IsValidClass() {
				classData.Last = false
				classes = append(classes, classData)
			}
		}
	}
	sort.SliceStable(classes, func(left int, right int) bool {
		return classes[left].ClassNode.FullName < classes[right].ClassNode.FullName
	})
	if len(classes) > 0 {
		classes[len(classes)-1].Last = true
		if processor.jsCfg.Index.Enabled {
			if err := processor.generateIndexFile(dir, classes); err != nil {
				return nil, err
			}
		}
		if processor.jsCfg.Typescript.Generate {
			if err := processor.generateTypeScript(dir, classes); err != nil {
				return nil, err
			}
		}
	}
	for _, sub := range subDirs {
		if _, statErr := os.Stat(sub); statErr != nil {
			continue
		}
		_, subErr := processor.processDir(root, sub)
		if subErr != nil {
			return nil, subErr
		}
	}
	return classes, nil
}

func (processor *Processor) generateIndexFile(dir string, classes []ClassData) error {
	payload := map[string]interface{}{
		"classes": classes,
	}
	content, err := processor.engine.Render(filepath.Join("js", "index"), payload)
	if err != nil {
		return err
	}
	target := filepath.Join(dir, processor.jsCfg.Index.Filename)
	if err := utils.EnsureDir(dir); err != nil {
		return err
	}
	return os.WriteFile(target, []byte(content), 0o644)
}

func (processor *Processor) generateTypeScript(dir string, classes []ClassData) error {
	tsCfg := processor.jsCfg.Typescript
	if !tsCfg.Generate {
		return nil
	}
	if tsCfg.OutputDir == "" {
		return nil
	}
	namespacePath := strings.ReplaceAll(classes[0].ClassNode.Namespace, ".", string(os.PathSeparator))
	targetDir := tsCfg.OutputDir
	if namespacePath != "" {
		targetDir = filepath.Join(tsCfg.OutputDir, namespacePath)
	}
	if err := utils.EnsureDir(targetDir); err != nil {
		return err
	}
	classMap := make(map[string]ClassData)
	for _, classData := range classes {
		payload := classData
		content, err := processor.engine.Render(filepath.Join("ts", "singleClass"), payload)
		if err != nil {
			return err
		}
		fileName := classData.ClassNode.FileInfo.Name + ".ts"
		if writeErr := os.WriteFile(filepath.Join(targetDir, fileName), []byte(content), 0o644); writeErr != nil {
			return writeErr
		}
		classMap[classData.ClassNode.FileInfo.Name] = classData
	}
	if len(classMap) == 0 {
		return nil
	}
	if tsCfg.GenerateIndex {
		barrelPayload := map[string]interface{}{
			"classes": classMap,
		}
		barrelContent, err := processor.engine.Render(filepath.Join("ts", "barrel"), barrelPayload)
		if err != nil {
			return err
		}
		if writeErr := os.WriteFile(filepath.Join(targetDir, tsCfg.BarrelFilename), []byte(barrelContent), 0o644); writeErr != nil {
			return writeErr
		}
		innerNamespace := filepath.Base(namespacePath)
		if innerNamespace == "." || innerNamespace == "" {
			innerNamespace = "Global"
		}
		indexPayload := map[string]interface{}{
			"innerNamespace": innerNamespace,
			"barrelFilename": strings.TrimSuffix(tsCfg.BarrelFilename, filepath.Ext(tsCfg.BarrelFilename)),
		}
		indexContent, err := processor.engine.Render(filepath.Join("ts", "index"), indexPayload)
		if err != nil {
			return err
		}
		if writeErr := os.WriteFile(filepath.Join(targetDir, tsCfg.IndexFilename), []byte(indexContent), 0o644); writeErr != nil {
			return writeErr
		}
	}
	return nil
}

func (processor *Processor) shouldGenerateGlobalIndex() bool {
	return processor.jsCfg.Index.Enabled && processor.jsCfg.GlobalIndex.Enabled
}

func (processor *Processor) allocateGlobalNamespace(namespace string) string {
	parts := strings.Split(namespace, ".")
	if len(parts) == 0 {
		parts = []string{"Global"}
	}
	depth := processor.jsCfg.GlobalIndex.Namespaces.Depth
	if depth <= 0 || depth > len(parts) {
		depth = len(parts)
	}
	selected := parts[len(parts)-depth:]
	candidate := ""
	for _, part := range selected {
		candidate += capitalize(part)
	}
	if candidate == "" {
		candidate = "Global"
	}
	original := candidate
	counter := 1
	for processor.usedNamespaces[candidate] {
		candidate = original + strconv.Itoa(counter)
		counter++
	}
	processor.usedNamespaces[candidate] = true
	return candidate
}

func (processor *Processor) addNamespace(namespace string) {
	if namespace == "" {
		return
	}
	parts := strings.Split(namespace, ".")
	if len(parts) == 0 {
		return
	}
	root := parts[0]
	remainder := parts[1:]
	node, ok := processor.namespaceTree[root]
	if !ok {
		node = make(map[string]interface{})
		processor.namespaceTree[root] = node
	}
	current := node
	for _, part := range remainder {
		child, exists := current[part]
		if !exists {
			nested := make(map[string]interface{})
			current[part] = nested
			current = nested
			continue
		}
		nested, ok := child.(map[string]interface{})
		if !ok {
			nested = make(map[string]interface{})
			current[part] = nested
		}
		current = nested
	}
}

func (processor *Processor) generateNamespaceFiles() error {
	if len(processor.namespaceTree) == 0 {
		return nil
	}
	outputDir := processor.jsCfg.Namespaces.OutputDir
	if outputDir == "" {
		return nil
	}
	if err := utils.EnsureDir(outputDir); err != nil {
		return err
	}
	for root, tree := range processor.namespaceTree {
		filename := strings.ReplaceAll(processor.jsCfg.Namespaces.Filename, "[root]", root)
		payload := map[string]interface{}{
			"root": root,
			"tree": processor.formatNamespaceTree(tree, 0),
		}
		content, err := processor.engine.Render(filepath.Join("js", "namespace"), payload)
		if err != nil {
			return err
		}
		target := filepath.Join(outputDir, filename)
		if err := os.WriteFile(target, []byte(content), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func capitalize(value string) string {
	if value == "" {
		return value
	}
	lower := strings.ToLower(value)
	result := strings.ToUpper(lower[:1]) + lower[1:]
	return result
}

func (processor *Processor) formatNamespaceTree(node map[string]interface{}, depth int) string {
	indent := strings.Repeat("    ", depth)
	builder := strings.Builder{}
	builder.WriteString("{\n")
	keys := make([]string, 0, len(node))
	for key := range node {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for index, key := range keys {
		child, _ := node[key].(map[string]interface{})
		builder.WriteString(indent)
		builder.WriteString("    ")
		builder.WriteString(key)
		builder.WriteString(": ")
		builder.WriteString(processor.formatNamespaceTree(child, depth+1))
		if index < len(keys)-1 {
			builder.WriteString(",")
		}
		builder.WriteString("\n")
	}
	builder.WriteString(indent)
	builder.WriteString("}")
	return builder.String()
}

func (processor *Processor) generateGlobalIndexFile() error {
	payload := map[string]interface{}{
		"globalIndexClasses": processor.globalIndex,
	}
	content, err := processor.engine.Render(filepath.Join("js", "global-index"), payload)
	if err != nil {
		return err
	}
	dir := processor.jsCfg.GlobalIndex.OutputDir
	if dir == "" {
		dir = processor.jsCfg.Index.OutputDir
	}
	if dir == "" {
		dir = "."
	}
	if err := utils.EnsureDir(dir); err != nil {
		return err
	}
	target := filepath.Join(dir, processor.jsCfg.GlobalIndex.Filename)
	if err := os.WriteFile(target, []byte(content), 0o644); err != nil {
		return err
	}
	return nil
}
