package rust

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

const generatorName = "rust"

var rustTypeMap = map[string]string{
	generator.TypeString:   "String",
	generator.TypeBool:     "bool",
	generator.TypeInt:      "i32",
	generator.TypeUInt:     "u32",
	generator.TypeDouble:   "f64",
	generator.TypeObject:   "String",
	generator.TypeFunction: "String",
	generator.TypeVoid:     "()",
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
		instance := &Generator{Base: generator.NewBase(rustTypeMap)}
		instance.Initialize(ctx, document, options)
		return instance
	})
}

func (Generator) Name() string {
	return generatorName
}

func (rGen *Generator) Generate() error {
	document := rGen.Document()
	targetDir := rustNamespacePath(rGen.OutputDir(), document.Namespace.Name.Segments)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}
	handler := func(classNode ast.Class) error {
		ctx := classContext{namespace: document.Namespace, classNode: classNode, targetDir: targetDir}
		return rGen.writeStruct(ctx)
	}
	return rGen.ForEachClass(handler)
}

func rustNamespacePath(root string, segments []string) string {
	parts := make([]string, 0, len(segments)+1)
	parts = append(parts, root)
	for _, segment := range segments {
		parts = append(parts, strings.ToLower(segment))
	}
	return filepath.Join(parts...)
}

func (rGen *Generator) writeStruct(ctx classContext) error {
	content, err := rGen.renderStruct(ctx)
	if err != nil {
		return err
	}
	className := stringy.New(ctx.classNode.Name).PascalCase().Get()
	path := filepath.Join(ctx.targetDir, strings.ToLower(className)+".rs")
	return os.WriteFile(path, []byte(content), 0o644)
}

func (rGen *Generator) renderStruct(ctx classContext) (string, error) {
	builder := new(strings.Builder)
	builder.WriteString("#[derive(Debug, Clone, Default)]\n")
	structName := stringy.New(ctx.classNode.Name).PascalCase().Get()
	builder.WriteString("pub struct " + structName + " {\n")
	fields, err := rGen.buildFields(ctx.classNode)
	if err != nil {
		return "", err
	}
	for _, entry := range fields {
		builder.WriteString("    " + entry.Line + "\n")
	}
	builder.WriteString("}\n\n")
	consts := buildConstants(ctx.classNode)
	if len(consts) > 0 {
		builder.WriteString("impl " + structName + " {\n")
		for _, entry := range consts {
			builder.WriteString("    " + entry.Line + "\n")
		}
		builder.WriteString("}\n")
	}
	return builder.String(), nil
}

func (rGen *Generator) buildFields(classNode ast.Class) ([]fieldEntry, error) {
	result := make([]fieldEntry, 0)
	for _, member := range classNode.Members {
		if !generator.IsPropertyMember(member) {
			continue
		}
		propType, err := member.PropertyType()
		if err != nil {
			return result, err
		}
		definition := rGen.rustField(member.Name, propType)
		result = append(result, fieldEntry{Line: definition})
	}
	return result, nil
}

func (rGen *Generator) rustField(name string, propType ast.PropertyType) string {
	fieldName := stringy.New(name).SnakeCase().Get()
	base := rGen.rustType(propType)
	return "pub " + fieldName + ": Option<" + base + ">,"
}

func (rGen *Generator) rustType(propType ast.PropertyType) string {
	base := rGen.rustIdentifierType(propType.Type)
	for range propType.ArrayNotation {
		base = "Vec<" + base + ">"
	}
	return base
}

func (rGen *Generator) rustIdentifierType(identifier ast.Identifier) string {
	if mapped, ok := rGen.FromPdlType(identifier.QualifiedName); ok {
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
		line := rustConstLine(member)
		if line != "" {
			result = append(result, constEntry{Line: line})
		}
	}
	sort.Slice(result, func(i int, j int) bool {
		return result[i].Line < result[j].Line
	})
	return result
}

func rustConstLine(member ast.Member) string {
	name := strings.ToUpper(stringy.New(member.Name).SnakeCase().Get())
	switch value := member.Value.(type) {
	case string:
		return "pub const " + name + ": &str = " + strconv.Quote(value) + ";"
	case bool:
		if value {
			return "pub const " + name + ": bool = true;"
		}
		return "pub const " + name + ": bool = false;"
	case int, int32, int64:
		return "pub const " + name + ": i64 = " + fmt.Sprint(value) + ";"
	case uint, uint32, uint64:
		return "pub const " + name + ": u64 = " + fmt.Sprint(value) + ";"
	case float32, float64:
		return "pub const " + name + ": f64 = " + fmt.Sprint(value) + ";"
	default:
		return ""
	}
}
