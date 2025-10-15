package kotlin

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/gobeam/stringy"

	"github.com/kapablanka/pdl/pdl/internal/ast"
	"github.com/kapablanka/pdl/pdl/internal/generator"
)

const generatorName = "kotlin"

var kotlinTypeMap = map[string]string{
	generator.TypeString:   "String",
	generator.TypeBool:     "Boolean",
	generator.TypeInt:      "Int",
	generator.TypeUInt:     "Long",
	generator.TypeDouble:   "Double",
	generator.TypeObject:   "Any",
	generator.TypeFunction: "() -> Unit",
	generator.TypeVoid:     "Unit",
}

type Generator struct {
	generator.Base
}

type classContext struct {
	namespace ast.Namespace
	classNode ast.Class
	targetDir string
}

type propertyEntry struct {
	Line string
}

type constEntry struct {
	Line string
}

func init() {
	generator.Register(generatorName, func(ctx context.Context, document ast.Document, options generator.Options) generator.Generator {
		instance := &Generator{Base: generator.NewBase(kotlinTypeMap)}
		instance.Initialize(ctx, document, options)
		return instance
	})
}

func (Generator) Name() string {
	return generatorName
}

func (kGen *Generator) Generate() error {
	document := kGen.Document()
	targetDir := kotlinNamespacePath(kGen.OutputDir(), document.Namespace.Name.Segments)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}
	handler := func(classNode ast.Class) error {
		ctx := classContext{namespace: document.Namespace, classNode: classNode, targetDir: targetDir}
		return kGen.writeClass(ctx)
	}
	return kGen.ForEachClass(handler)
}

func kotlinNamespacePath(root string, segments []string) string {
	parts := make([]string, 0, len(segments)+1)
	parts = append(parts, root)
	for _, segment := range segments {
		parts = append(parts, strings.ToLower(segment))
	}
	return filepath.Join(parts...)
}

func (kGen *Generator) writeClass(ctx classContext) error {
	content, err := kGen.renderClass(ctx)
	if err != nil {
		return err
	}
	className := stringy.New(ctx.classNode.Name).PascalCase().Get()
	path := filepath.Join(ctx.targetDir, className+".kt")
	return os.WriteFile(path, []byte(content), 0o644)
}

func (kGen *Generator) renderClass(ctx classContext) (string, error) {
	builder := new(strings.Builder)
	builder.WriteString(kotlinPackageLine(ctx.namespace))
	props, perr := kGen.buildProperties(ctx.classNode)
	if perr != nil {
		return "", perr
	}
	consts := buildConstants(ctx.classNode)
	kGen.writeClassHeader(builder, ctx.classNode, props)
	kGen.writePropertyBlock(builder, props)
	writeConstBlock(builder, consts)
	builder.WriteString("}\n")
	return builder.String(), nil
}

func kotlinPackageLine(namespace ast.Namespace) string {
	if len(namespace.Name.Segments) == 0 {
		return ""
	}
	segments := make([]string, len(namespace.Name.Segments))
	for idx, segment := range namespace.Name.Segments {
		segments[idx] = strings.ToLower(segment)
	}
	return "package " + strings.Join(segments, ".") + "\n\n"
}

func (kGen *Generator) buildProperties(classNode ast.Class) ([]propertyEntry, error) {
	result := make([]propertyEntry, 0)
	for _, member := range classNode.Members {
		if !generator.IsPropertyMember(member) {
			continue
		}
		propType, err := member.PropertyType()
		if err != nil {
			return result, err
		}
		line := kGen.kotlinPropertyLine(member.Name, propType)
		result = append(result, propertyEntry{Line: line})
	}
	return result, nil
}

func (kGen *Generator) kotlinPropertyLine(name string, propType ast.PropertyType) string {
	param := stringy.New(name).CamelCase().Get()
	typeName := kGen.kotlinType(propType)
	if strings.HasPrefix(typeName, "List<") {
		return "var " + param + ": " + typeName + " = emptyList()"
	}
	return "var " + param + ": " + typeName + "? = null"
}

func (kGen *Generator) kotlinType(propType ast.PropertyType) string {
	base := kGen.kotlinIdentifierType(propType.Type)
	for range propType.ArrayNotation {
		base = "List<" + base + ">"
	}
	return base
}

func (kGen *Generator) kotlinIdentifierType(identifier ast.Identifier) string {
	if mapped, ok := kGen.FromPdlType(identifier.QualifiedName); ok {
		return mapped
	}
	if len(identifier.Segments) == 0 {
		return stringy.New(identifier.QualifiedName).PascalCase().Get()
	}
	segments := make([]string, len(identifier.Segments))
	for idx, segment := range identifier.Segments {
		segments[idx] = stringy.New(segment).PascalCase().Get()
	}
	return strings.Join(segments, ".")
}

func buildConstants(classNode ast.Class) []constEntry {
	result := make([]constEntry, 0)
	for _, member := range classNode.Members {
		if member.Kind != "const" {
			continue
		}
		line := kotlinConstLine(member)
		if line != "" {
			result = append(result, constEntry{Line: line})
		}
	}
	sort.Slice(result, func(i int, j int) bool {
		return result[i].Line < result[j].Line
	})
	return result
}

func kotlinConstLine(member ast.Member) string {
	name := strings.ToUpper(stringy.New(member.Name).SnakeCase().Get())
	switch value := member.Value.(type) {
	case string:
		return "const val " + name + " = " + strconv.Quote(value)
	case bool:
		if value {
			return "const val " + name + " = true"
		}
		return "const val " + name + " = false"
	case int, int32, int64, uint, uint32, uint64, float32, float64:
		return "const val " + name + " = " + fmt.Sprint(value)
	default:
		return ""
	}
}

func (kGen *Generator) writeClassHeader(builder *strings.Builder, classNode ast.Class, props []propertyEntry) {
	className := stringy.New(classNode.Name).PascalCase().Get()
	builder.WriteString("data class " + className)
	parent := kotlinParentClause(classNode)
	kGen.writeConstructor(builder, props, parent)
}

func (kGen *Generator) writeConstructor(builder *strings.Builder, props []propertyEntry, parent string) {
	if len(props) == 0 {
		builder.WriteString("()")
	} else {
		builder.WriteString("(\n")
		for idx, entry := range props {
			builder.WriteString("    " + entry.Line)
			if idx+1 < len(props) {
				builder.WriteString(",\n")
			} else {
				builder.WriteString("\n")
			}
		}
		builder.WriteString(")")
	}
	if parent != "" {
		builder.WriteString(" : " + parent)
	}
	builder.WriteString(" {\n")
}

func kotlinParentClause(classNode ast.Class) string {
	if classNode.Parent == nil {
		return ""
	}
	segments := make([]string, len(classNode.Parent.Segments))
	for idx, segment := range classNode.Parent.Segments {
		segments[idx] = stringy.New(segment).PascalCase().Get()
	}
	if len(segments) == 0 {
		return stringy.New(classNode.Parent.QualifiedName).PascalCase().Get() + "()"
	}
	return strings.Join(segments, ".") + "()"
}

func (kGen *Generator) writePropertyBlock(builder *strings.Builder, props []propertyEntry) {
	if len(props) == 0 {
		return
	}
	builder.WriteString("\n")
}

func writeConstBlock(builder *strings.Builder, consts []constEntry) {
	if len(consts) == 0 {
		return
	}
	builder.WriteString("    companion object {\n")
	for _, entry := range consts {
		builder.WriteString("        " + entry.Line + "\n")
	}
	builder.WriteString("    }\n")
}
