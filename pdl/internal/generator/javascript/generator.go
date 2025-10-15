package javascript

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kapablanka/pdl/pdl/internal/ast"
	"github.com/kapablanka/pdl/pdl/internal/generator"
)

const generatorName = "javascript"

var javascriptTypeMap = map[string]string{
	generator.TypeString:   "String",
	generator.TypeBool:     "Boolean",
	generator.TypeInt:      "Number",
	generator.TypeUInt:     "Number",
	generator.TypeDouble:   "Number",
	generator.TypeArray:    "Array",
	generator.TypeFunction: "Function",
	generator.TypeObject:   "Object",
	generator.TypeVoid:     "void",
}

type Generator struct {
	generator.Base
}

func init() {
	generator.Register(generatorName, func(ctx context.Context, document ast.Document, options generator.Options) generator.Generator {
		instance := &Generator{Base: generator.NewBase(javascriptTypeMap)}
		instance.Initialize(ctx, document, options)
		return instance
	})
}

func (Generator) Name() string {
	return generatorName
}

func (jsGen *Generator) Generate() error {
	namespace := jsGen.Document().Namespace
	outputDir := jsGen.OutputDir()
	handler := func(classNode ast.Class) error {
		return jsGen.generateClass(namespace, outputDir, classNode)
	}
	return jsGen.ForEachClass(handler)
}

func (jsGen *Generator) generateClass(namespace ast.Namespace, outputDir string, classNode ast.Class) error {
	var result error
	classPathSegments := append([]string{}, namespace.Name.Segments...)
	classPathSegments = append(classPathSegments, classNode.Name+".js")
	targetFile := filepath.Join(outputDir, filepath.Join(classPathSegments...))
	content := renderClass(namespace, classNode)
	jsGen.WriteFile(targetFile, []byte(content))
	return result
}

func renderClass(namespace ast.Namespace, classNode ast.Class) string {
	qualifiedNamespace := namespace.Name.QualifiedName
	classQualified := classQualifiedName(qualifiedNamespace, classNode.Name)
	parentQualified := parentQualifiedName(qualifiedNamespace, classNode)
	builder := new(strings.Builder)
	builder.WriteString("/**\n* PDL Compiler generated code\n")
	builder.WriteString("* class " + classQualified + "\n")
	builder.WriteString("*/\n\n\n")
	builder.WriteString("const inherit = require('pdl/infrastructure/languages/js/inheritance').inherit;\n\n")
	builder.WriteString(renderDocBlock(classNode.Name, "null", classNode, true))
	builder.WriteString(fmt.Sprintf("const %s = function () {\n\n", classNode.Name))
	builder.WriteString(fmt.Sprintf("    this.__type = '%s';\n", classQualified))
	builder.WriteString("        \n        \n")
	builder.WriteString("};\n\n")
	builder.WriteString(renderDocBlock(classQualified, parentQualified, classNode, false))
	builder.WriteString(classQualified + " = " + classNode.Name + ";\n")
	builder.WriteString(renderPropertyAttributes(classQualified, classNode))
	builder.WriteString(renderTypedef(classQualified, classNode.Name))
	builder.WriteString(renderInheritance(namespace, classNode.Name, parentQualified))
	builder.WriteString(fmt.Sprintf("export default %s;\n", classNode.Name))
	return builder.String()
}

func renderDocBlock(displayName string, extends string, classNode ast.Class, includeExtends bool) string {
	builder := new(strings.Builder)
	builder.WriteString("/**\n")
	builder.WriteString(fmt.Sprintf(" * @class {%s}\n", displayName))
	if includeExtends {
		builder.WriteString(fmt.Sprintf(" * @extends {%s}\n", extends))
	} else if extends != "" && extends != "null" {
		builder.WriteString(fmt.Sprintf(" * @extends {%s}\n", extends))
	}
	builder.WriteString(" *\n")
	for _, member := range classNode.Members {
		if member.Kind != "property" && member.Kind != "shortProperty" {
			continue
		}
		docType := propertyDocType(member)
		builder.WriteString(fmt.Sprintf(" * @property {%s} %s\n", docType, member.Name))
	}
	builder.WriteString(" */\n")
	return builder.String()
}

func renderPropertyAttributes(classQualified string, classNode ast.Class) string {
	propertyLines := make([]string, 0)
	for _, member := range classNode.Members {
		if len(member.Attributes) == 0 {
			continue
		}
		lines := make([]string, 0)
		for _, attribute := range member.Attributes {
			lines = append(lines, renderAttribute(attribute))
		}
		propertyLines = append(propertyLines, fmt.Sprintf("propertyAttributes.%s = [\n%s\n];", member.Name, strings.Join(lines, "\n")))
	}
	builder := new(strings.Builder)
	builder.WriteString("let propertyAttributes = {};\n")
	if len(propertyLines) > 0 {
		for _, line := range propertyLines {
			builder.WriteString(line + "\n")
		}
	}
	builder.WriteString(classQualified + ".__propertyAttributes = propertyAttributes;\n\n")
	return builder.String()
}

func renderAttribute(attribute ast.Attribute) string {
	builder := new(strings.Builder)
	builder.WriteString("    {\n")
	builder.WriteString(fmt.Sprintf("        name: '%s',\n", attribute.Name.QualifiedName))
	builder.WriteString("        values: {\n")
	for index, value := range attribute.Params.Required {
		key := fmt.Sprintf("default%d", index+1)
		builder.WriteString(fmt.Sprintf("        %s: \"%v\"\n", key, value))
	}
	builder.WriteString("        }\n")
	builder.WriteString("    }")
	return builder.String()
}

func renderTypedef(classQualified string, className string) string {
	return fmt.Sprintf("/** @typedef {%s|%s} %s */\n\n", classQualified, className, className)
}

func renderInheritance(namespace ast.Namespace, className string, parentQualified string) string {
	namespaceExpr := strings.Join(namespace.Name.Segments, ".")
	if namespaceExpr == "" {
		namespaceExpr = "this"
	}
	inheritParent := "null"
	if parentQualified != "" {
		inheritParent = parentQualified
	}
	return fmt.Sprintf("/** inheritance */\ninherit( %s, '%s', %s );\n\n", namespaceExpr, className, inheritParent)
}

func classQualifiedName(namespaceQualified string, className string) string {
	if namespaceQualified == "" {
		return className
	}
	return namespaceQualified + "." + className
}

func parentQualifiedName(namespaceQualified string, classNode ast.Class) string {
	if classNode.Parent == nil {
		return "null"
	}
	parent := classNode.Parent.QualifiedName
	if strings.Contains(parent, ".") {
		return parent
	}
	if namespaceQualified == "" {
		return parent
	}
	return namespaceQualified + "." + parent
}

func propertyDocType(member ast.Member) string {
	for _, attribute := range member.Attributes {
		if strings.EqualFold(attribute.Name.QualifiedName, "jstype") && len(attribute.Params.Required) > 0 {
			if value, ok := attribute.Params.Required[0].(string); ok {
				return value
			}
		}
	}
	propertyType, err := member.PropertyType()
	if err != nil {
		return "any"
	}
	rawType := propertyType.Type.QualifiedName
	baseType := mapTypeName(rawType)
	if len(propertyType.ArrayNotation) > 0 {
		baseType += "[]"
	}
	return baseType
}

func mapTypeName(name string) string {
	switch strings.ToLower(name) {
	case "string":
		return "String"
	case "bool", "boolean":
		return "Boolean"
	case "double", "float", "decimal":
		return "Number"
	case "int", "int32", "int64":
		return "Number"
	case "object":
		return "Object"
	default:
		return name
	}
}

func isBuiltInType(name string) bool {
	switch name {
	case "String", "Boolean", "Number", "Object":
		return true
	default:
		return false
	}
}
