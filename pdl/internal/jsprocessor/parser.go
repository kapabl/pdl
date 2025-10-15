package jsprocessor

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

var intrinsicTypeMap = map[string]string{
	"boolean":    "boolean",
	"number":     "number",
	"int":        "number",
	"string":     "string",
	"function":   "Function",
	"object":     "any",
	"array":      "any[]",
	"[]":         "any[]",
	"boolean[]":  "boolean[]",
	"number[]":   "number[]",
	"int[]":      "number[]",
	"string[]":   "string[]",
	"function[]": "Function[]",
	"object[]":   "any[]",
}

type parserContext struct {
	current *ClassData
}

func parseClass(root string, filename string) (ClassData, error) {
	result := ClassData{
		JSFullFilename: filename,
		Imports:        make([]string, 0),
		Properties:     make([]Property, 0),
	}
	file, err := os.Open(filename)
	if err != nil {
		return result, err
	}
	defer file.Close()
	ctx := parserContext{current: &result}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "@class") {
			node := parseClassDeclaration(trimmed)
			ctx.current.ClassNode = node
		} else if strings.Contains(trimmed, "@extend") {
			parent := parseParentClass(trimmed)
			parent = ctx.addImport(parent)
			ctx.current.ParentClassNode = parent
		} else if strings.Contains(trimmed, "@property") {
			prop := ctx.parseProperty(trimmed)
			ctx.current.Properties = append(ctx.current.Properties, prop)
		}
	}
	if ctx.current.ClassNode.FileInfo.Name == "" {
		ctx.current.ClassNode = buildClassNodeFromPath(root, filename)
	}
	rel, relErr := filepath.Rel(root, filename)
	if relErr != nil {
		rel = filename
	}
	rel = filepath.ToSlash(rel)
	dir := filepath.Dir(rel)
	if dir == "." {
		dir = ""
	}
	name := strings.TrimSuffix(filepath.Base(rel), filepath.Ext(rel))
	ctx.current.ClassNode.Path = strings.TrimSuffix(rel, filepath.Ext(rel))
	ctx.current.ClassNode.FileInfo = FileInfo{Name: name, Dir: dir}
	return result, scanner.Err()
}

func parseClassDeclaration(line string) ClassNode {
	full := extractBraceContent(line)
	node := parseClassName(full)
	return node
}

func parseParentClass(line string) ClassNode {
	full := extractBraceContent(line)
	node := parseClassName(full)
	return node
}

func (ctx *parserContext) parseProperty(line string) Property {
	content := strings.TrimSpace(strings.TrimPrefix(line, "*"))
	content = strings.TrimSpace(strings.TrimPrefix(content, "@property"))
	typeSegment := extractBraceContent(content)
	tsType := jsTypeToTsType(typeSegment)
	withoutType := strings.Replace(content, "{"+typeSegment+"}", "", 1)
	tokens := strings.Fields(strings.TrimSpace(withoutType))
	var name string
	if len(tokens) > 0 {
		name = strings.Trim(tokens[len(tokens)-1], ";")
	}
	classNode := parseClassName(tsType)
	classNode = ctx.addImport(classNode)
	property := Property{Name: name, ClassNode: classNode}
	return property
}

func (ctx *parserContext) addImport(classNode ClassNode) ClassNode {
	if !classNode.IsClass || classNode.IsFloatingClass || ctx.current.ClassNode.FileInfo.Name == "" {
		return classNode
	}
	if classNode.IsIntrinsic {
		return classNode
	}
	if classNode.FullName == ctx.current.ClassNode.FullName {
		return classNode
	}
	classNode.IsImported = true
	if classNode.Namespace == ctx.current.ClassNode.Namespace {
		importStatement := "import {" + classNode.Name + "} from './" + classNode.Name + "';"
		ctx.appendImport(importStatement)
		return classNode
	}
	namespaceParts := strings.Split(classNode.Namespace, ".")
	innerNamespace := namespaceParts[len(namespaceParts)-1]
	classNode.UsageName = innerNamespace + "." + classNode.UsageName
	baseDir := ctx.current.ClassNode.FileInfo.Dir
	if baseDir == "" {
		baseDir = "."
	}
	targetDir := classNode.FileInfo.Dir
	if targetDir == "" {
		targetDir = "."
	}
	relative, relErr := filepath.Rel(baseDir, targetDir)
	if relErr != nil {
		relative = targetDir
	}
	relative = filepath.ToSlash(relative)
	if !strings.HasPrefix(relative, "./") && !strings.HasPrefix(relative, "../") {
		if relative == "." {
			relative = "./"
		} else {
			relative = "./" + relative
		}
	}
	importStatement := "import {" + innerNamespace + "} from '" + relative + "';"
	ctx.appendImport(importStatement)
	return classNode
}

func (ctx *parserContext) appendImport(statement string) {
	for _, existing := range ctx.current.Imports {
		if existing == statement {
			return
		}
	}
	ctx.current.Imports = append(ctx.current.Imports, statement)
}

func parseClassName(full string) ClassNode {
	node := ClassNode{}
	cleaned := strings.TrimSpace(full)
	if cleaned == "" {
		return node
	}
	parts := strings.Split(cleaned, ".")
	node.Path = strings.Join(parts, "/")
	if len(parts) > 1 {
		node.Namespace = strings.Join(parts[:len(parts)-1], ".")
	}
	namePart := parts[len(parts)-1]
	arraySplit := strings.Split(namePart, "[")
	baseName := arraySplit[0]
	node.IsArray = len(arraySplit) > 1
	node.Name = baseName
	node.FullName = cleaned
	node.IsIntrinsic = isIntrinsicType(baseName)
	node.IsClass = node.Namespace != "" || !node.IsIntrinsic
	node.IsFloatingClass = node.IsClass && node.Namespace == ""
	node.UsageName = baseName
	if node.IsArray {
		node.UsageName = node.UsageName + "[]"
	}
	dir := filepath.Dir(node.Path)
	if dir == "." {
		dir = ""
	}
	node.FileInfo = FileInfo{Name: baseName, Dir: dir}
	return node
}

func buildClassNodeFromPath(root string, filename string) ClassNode {
	rel, relErr := filepath.Rel(root, filename)
	if relErr != nil {
		rel = filename
	}
	rel = filepath.ToSlash(rel)
	withoutExt := strings.TrimSuffix(rel, filepath.Ext(rel))
	parts := strings.Split(withoutExt, "/")
	namespace := ""
	if len(parts) > 1 {
		namespace = strings.Join(parts[:len(parts)-1], ".")
	}
	name := parts[len(parts)-1]
	full := name
	if namespace != "" {
		full = namespace + "." + name
	}
	node := parseClassName(full)
	return node
}

func extractBraceContent(line string) string {
	start := strings.Index(line, "{")
	if start == -1 {
		return ""
	}
	sub := line[start+1:]
	end := strings.Index(sub, "}")
	if end == -1 {
		return ""
	}
	content := strings.TrimSpace(sub[:end])
	return content
}

func isIntrinsicType(name string) bool {
	_, found := intrinsicTypeMap[strings.ToLower(name)]
	return found
}

func jsTypeToTsType(name string) string {
	value, found := intrinsicTypeMap[strings.ToLower(name)]
	if !found {
		return name
	}
	return value
}
