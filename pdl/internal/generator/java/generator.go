package java

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/kapablanka/pdl/pdl/internal/ast"
	"github.com/kapablanka/pdl/pdl/internal/generator"
)

const generatorName = "java"

var javaIntrinsicTypes = map[string]string{
	generator.TypeString:   "String",
	generator.TypeBool:     "Boolean",
	generator.TypeVoid:     "Void",
	generator.TypeInt:      "Integer",
	generator.TypeUInt:     "Long",
	generator.TypeDouble:   "Double",
	generator.TypeArray:    "Object",
	generator.TypeFunction: "Object",
	generator.TypeObject:   "Object",
}

type Generator struct {
	generator.Base
}

func init() {
	generator.Register(generatorName, func(ctx context.Context, document ast.Document, options generator.Options) generator.Generator {
		instance := &Generator{Base: generator.NewBase(javaIntrinsicTypes)}
		instance.Initialize(ctx, document, options)
		return instance
	})
}

func (Generator) Name() string {
	result := generatorName
	return result
}

func (javaGen *Generator) Generate() error {
	document := javaGen.Document()
	targetDir := javaNamespacePath(javaGen.OutputDir(), document.Namespace.Name.Segments)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}
	for _, classNode := range document.Namespace.Classes {
		if ctx := javaGen.Context(); ctx != nil {
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		ctx := classContext{
			namespace: document.Namespace,
			classNode: classNode,
			targetDir: targetDir,
		}
		if writeErr := writeJavaClass(ctx); writeErr != nil {
			return writeErr
		}
	}
	return nil
}

type classContext struct {
	namespace ast.Namespace
	classNode ast.Class
	targetDir string
}

type classArtifacts struct {
	PackageLine string
	Imports     []string
	Constants   []constEntry
	Fields      []fieldEntry
	Methods     []methodEntry
	Extends     string
}

type constEntry struct {
	Name  string
	Type  string
	Value string
}

type fieldEntry struct {
	FieldName string
	FieldType string
	Property  string
}

type methodEntry struct {
	Signature string
	Body      []string
}

func writeJavaClass(ctx classContext) error {
	var result error
	artifacts, buildErr := buildArtifacts(ctx)
	if buildErr != nil {
		result = buildErr
		return result
	}
	content := renderClassSource(ctx, artifacts)
	filePath := filepath.Join(ctx.targetDir, ctx.classNode.Name+".java")
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		result = err
		return result
	}
	return result
}

func buildArtifacts(ctx classContext) (classArtifacts, error) {
	var result classArtifacts
	result.PackageLine = javaPackageLine(ctx.namespace)
	constants := collectConstants(ctx.classNode)
	fields, fieldErr := collectFields(ctx.classNode)
	if fieldErr != nil {
		return result, fieldErr
	}
	methods := collectAccessors(fields)
	imports := collectImports(constants, fields)
	parent := javaExtends(ctx.classNode)
	result = classArtifacts{
		PackageLine: result.PackageLine,
		Imports:     imports,
		Constants:   constants,
		Fields:      fields,
		Methods:     methods,
		Extends:     parent,
	}
	return result, nil
}

func renderClassSource(ctx classContext, artifacts classArtifacts) string {
	var result string
	builder := new(strings.Builder)
	builder.WriteString(artifacts.PackageLine)
	if len(artifacts.Imports) > 0 {
		for _, imp := range artifacts.Imports {
			builder.WriteString("import " + imp + ";\n")
		}
		builder.WriteString("\n")
	}
	builder.WriteString("public class " + ctx.classNode.Name)
	if artifacts.Extends != "" {
		builder.WriteString(" extends " + artifacts.Extends)
	}
	builder.WriteString(" {\n")
	writeConstants(builder, artifacts.Constants)
	writeFields(builder, artifacts.Fields)
	writeConstructor(builder, ctx.classNode.Name)
	writeMethods(builder, artifacts.Methods)
	builder.WriteString("}\n")
	result = builder.String()
	return result
}

func javaPackageLine(namespace ast.Namespace) string {
	var result string
	if len(namespace.Name.Segments) == 0 {
		return result
	}
	segments := make([]string, 0, len(namespace.Name.Segments))
	for _, segment := range namespace.Name.Segments {
		segments = append(segments, strings.ToLower(segment))
	}
	result = "package " + strings.Join(segments, ".") + ";\n\n"
	return result
}

func javaNamespacePath(root string, segments []string) string {
	parts := make([]string, 0, len(segments)+1)
	parts = append(parts, root)
	for _, segment := range segments {
		parts = append(parts, strings.ToLower(segment))
	}
	result := filepath.Join(parts...)
	return result
}

func collectConstants(classNode ast.Class) []constEntry {
	result := make([]constEntry, 0)
	for _, member := range classNode.Members {
		if member.Kind != "const" {
			continue
		}
		typeInfo, err := member.ConstType()
		if err != nil {
			continue
		}
		value := javaLiteral(member.Value)
		if value == "" {
			continue
		}
		entry := constEntry{
			Name:  generator.SnakeCaseUpper(member.Name),
			Type:  javaTypeName(typeInfo),
			Value: value,
		}
		result = append(result, entry)
	}
	sort.SliceStable(result, func(left int, right int) bool {
		return result[left].Name < result[right].Name
	})
	return result
}

func collectFields(classNode ast.Class) ([]fieldEntry, error) {
	result := make([]fieldEntry, 0)
	for _, member := range classNode.Members {
		if member.Kind != "property" && member.Kind != "shortProperty" {
			continue
		}
		propType, err := member.PropertyType()
		if err != nil {
			return result, err
		}
		fieldType := javaPropertyType(propType)
		entry := fieldEntry{
			FieldName: generator.CamelCase(member.Name),
			FieldType: fieldType,
			Property:  member.Name,
		}
		result = append(result, entry)
	}
	return result, nil
}

func collectAccessors(fields []fieldEntry) []methodEntry {
	result := make([]methodEntry, 0, len(fields)*2)
	for _, field := range fields {
		getter := methodEntry{
			Signature: "public " + field.FieldType + " get" + generator.Capitalize(field.Property) + "()",
			Body:      []string{"return this." + field.FieldName + ";"},
		}
		setter := methodEntry{
			Signature: "public void set" + generator.Capitalize(field.Property) + "(" + field.FieldType + " value)",
			Body: []string{
				"this." + field.FieldName + " = value;",
			},
		}
		result = append(result, getter, setter)
	}
	return result
}

func collectImports(constants []constEntry, fields []fieldEntry) []string {
	needsList := false
	for _, field := range fields {
		if strings.HasPrefix(field.FieldType, "List<") {
			needsList = true
			break
		}
	}
	imports := make([]string, 0)
	if needsList {
		imports = append(imports, "java.util.List")
	}
	sort.Strings(imports)
	return imports
}

func javaExtends(classNode ast.Class) string {
	var result string
	if classNode.Parent == nil {
		return result
	}
	result = generator.SimpleName(classNode.Parent.QualifiedName)
	return result
}

func writeConstants(builder *strings.Builder, constants []constEntry) {
	if len(constants) == 0 {
		return
	}
	for _, constant := range constants {
		builder.WriteString("    public static final " + constant.Type + " " + constant.Name + " = " + constant.Value + ";\n")
	}
	builder.WriteString("\n")
}

func writeFields(builder *strings.Builder, fields []fieldEntry) {
	for _, field := range fields {
		builder.WriteString("    private " + field.FieldType + " " + field.FieldName + ";\n")
	}
	if len(fields) > 0 {
		builder.WriteString("\n")
	}
}

func writeConstructor(builder *strings.Builder, className string) {
	builder.WriteString("    public " + className + "() {\n")
	builder.WriteString("    }\n\n")
}

func writeMethods(builder *strings.Builder, methods []methodEntry) {
	for index, method := range methods {
		builder.WriteString("    " + method.Signature + " {\n")
		for _, line := range method.Body {
			builder.WriteString("        " + line + "\n")
		}
		builder.WriteString("    }\n")
		if index+1 < len(methods) {
			builder.WriteString("\n")
		}
	}
}

func javaPropertyType(propType ast.PropertyType) string {
	result := javaTypeName(propType.Type)
	if len(propType.ArrayNotation) == 0 {
		return result
	}
	result = "List<" + result + ">"
	return result
}

func javaTypeName(identifier ast.Identifier) string {
	name := identifier.QualifiedName
	if intrinsic, exists := javaIntrinsicTypes[name]; exists {
		return intrinsic
	}
	if len(identifier.Segments) == 0 {
		return generator.Capitalize(identifier.QualifiedName)
	}
	return generator.SimpleName(identifier.QualifiedName)
}

func javaLiteral(value interface{}) string {
	var result string
	switch typed := value.(type) {
	case string:
		result = strconv.Quote(typed)
	case bool:
		if typed {
			result = "true"
		} else {
			result = "false"
		}
	case float64:
		result = strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", typed), "0"), ".")
	default:
		result = fmt.Sprint(value)
	}
	return result
}
