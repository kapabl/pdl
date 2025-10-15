package cpp

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/gobeam/stringy"

	"github.com/kapablanka/pdl/pdl/internal/ast"
	"github.com/kapablanka/pdl/pdl/internal/generator"
)

const generatorName = "cpp"

var cppTypeMap = map[string]string{
	generator.TypeString:   "std::string",
	generator.TypeBool:     "bool",
	generator.TypeInt:      "int",
	generator.TypeUInt:     "unsigned int",
	generator.TypeDouble:   "double",
	generator.TypeObject:   "std::string",
	generator.TypeFunction: "std::string",
	generator.TypeVoid:     "void",
}

type Generator struct {
	generator.Base
}

type classContext struct {
	namespace ast.Namespace
	classNode ast.Class
	targetDir string
}

type fieldEntry struct {
	Line string
}

type constEntry struct {
	Line string
}

func init() {
	generator.Register(generatorName, func(ctx context.Context, document ast.Document, options generator.Options) generator.Generator {
		instance := &Generator{Base: generator.NewBase(cppTypeMap)}
		instance.Initialize(ctx, document, options)
		return instance
	})
}

func (Generator) Name() string {
	return generatorName
}

func (cGen *Generator) Generate() error {
	document := cGen.Document()
	targetDir := filepath.Join(cGen.OutputDir(), "include")
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}
	handler := func(classNode ast.Class) error {
		ctx := classContext{namespace: document.Namespace, classNode: classNode, targetDir: targetDir}
		return cGen.writeHeader(ctx)
	}
	return cGen.ForEachClass(handler)
}

func (cGen *Generator) writeHeader(ctx classContext) error {
	content, err := cGen.renderHeader(ctx)
	if err != nil {
		return err
	}
	className := stringy.New(ctx.classNode.Name).PascalCase().Get()
	path := filepath.Join(ctx.targetDir, className+".hpp")
	return os.WriteFile(path, []byte(content), 0o644)
}

func (cGen *Generator) renderHeader(ctx classContext) (string, error) {
	builder := new(strings.Builder)
	className := stringy.New(ctx.classNode.Name).PascalCase().Get()
	builder.WriteString("#pragma once\n\n#include <optional>\n#include <string>\n#include <vector>\n\n")
	openNamespaces(builder, ctx.namespace)
	builder.WriteString("struct " + className)
	parent := cppParentClause(ctx.classNode)
	if parent != "" {
		builder.WriteString(" : public " + parent)
	}
	builder.WriteString(" {\n")
	fields, err := cGen.buildFields(ctx.classNode)
	if err != nil {
		return "", err
	}
	for _, entry := range fields {
		builder.WriteString("    " + entry.Line + "\n")
	}
	consts := buildConstants(ctx.classNode)
	if len(consts) > 0 {
		builder.WriteString("\n    struct Constants {\n")
		for _, entry := range consts {
			builder.WriteString("        " + entry.Line + "\n")
		}
		builder.WriteString("    };\n")
	}
	builder.WriteString("};\n")
	closeNamespaces(builder, ctx.namespace)
	return builder.String(), nil
}

func (cGen *Generator) buildFields(classNode ast.Class) ([]fieldEntry, error) {
	result := make([]fieldEntry, 0)
	for _, member := range classNode.Members {
		if !generator.IsPropertyMember(member) {
			continue
		}
		propType, err := member.PropertyType()
		if err != nil {
			return result, err
		}
		line := cGen.cppField(member.Name, propType)
		result = append(result, fieldEntry{Line: line})
	}
	return result, nil
}

func (cGen *Generator) cppField(name string, propType ast.PropertyType) string {
	fieldName := stringy.New(name).CamelCase().Get()
	base := cGen.cppType(propType)
	return "std::optional<" + base + "> " + fieldName + ";"
}

func (cGen *Generator) cppType(propType ast.PropertyType) string {
	base := cGen.cppIdentifierType(propType.Type)
	for range propType.ArrayNotation {
		base = "std::vector<" + base + ">"
	}
	return base
}

func (cGen *Generator) cppIdentifierType(identifier ast.Identifier) string {
	if mapped, ok := cGen.FromPdlType(identifier.QualifiedName); ok {
		return mapped
	}
	if len(identifier.Segments) == 0 {
		return stringy.New(identifier.QualifiedName).PascalCase().Get()
	}
	segments := make([]string, len(identifier.Segments))
	for idx, segment := range identifier.Segments {
		segments[idx] = stringy.New(segment).PascalCase().Get()
	}
	return strings.Join(segments, "::")
}

func buildConstants(classNode ast.Class) []constEntry {
	result := make([]constEntry, 0)
	for _, member := range classNode.Members {
		if member.Kind != "const" {
			continue
		}
		line := cppConstLine(member)
		if line != "" {
			result = append(result, constEntry{Line: line})
		}
	}
	sort.Slice(result, func(i int, j int) bool {
		return result[i].Line < result[j].Line
	})
	return result
}

func cppConstLine(member ast.Member) string {
	name := stringy.New(member.Name).CamelCase().Get()
	switch value := member.Value.(type) {
	case string:
		return "static constexpr const char* " + name + " = " + strconv.Quote(value) + ";"
	case bool:
		if value {
			return "static constexpr bool " + name + " = true;"
		}
		return "static constexpr bool " + name + " = false;"
	case int, int32, int64, uint, uint32, uint64, float32, float64:
		return "static constexpr auto " + name + " = " + strconv.FormatFloat(toFloat(value), 'f', -1, 64) + ";"
	default:
		return ""
	}
}

func toFloat(value interface{}) float64 {
	switch typed := value.(type) {
	case int:
		return float64(typed)
	case int32:
		return float64(typed)
	case int64:
		return float64(typed)
	case uint:
		return float64(typed)
	case uint32:
		return float64(typed)
	case uint64:
		return float64(typed)
	case float32:
		return float64(typed)
	case float64:
		return typed
	default:
		return 0
	}
}

func openNamespaces(builder *strings.Builder, namespace ast.Namespace) {
	for _, segment := range namespace.Name.Segments {
		builder.WriteString("namespace " + stringy.New(segment).CamelCase().Get() + " {\n")
	}
}

func closeNamespaces(builder *strings.Builder, namespace ast.Namespace) {
	for range namespace.Name.Segments {
		builder.WriteString("}\n")
	}
}

func cppParentClause(classNode ast.Class) string {
	if classNode.Parent == nil {
		return ""
	}
	segments := make([]string, len(classNode.Parent.Segments))
	for idx, segment := range classNode.Parent.Segments {
		segments[idx] = stringy.New(segment).PascalCase().Get()
	}
	if len(segments) == 0 {
		return stringy.New(classNode.Parent.QualifiedName).PascalCase().Get()
	}
	return strings.Join(segments, "::")
}
