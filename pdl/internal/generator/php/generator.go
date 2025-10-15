package php

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gobeam/stringy"

	"github.com/kapablanka/pdl/pdl/internal/ast"
	"github.com/kapablanka/pdl/pdl/internal/generator"
)

const generatorName = "php"

var phpTypeMap = map[string]string{
	generator.TypeString:   "string",
	generator.TypeBool:     "bool",
	generator.TypeInt:      "int",
	generator.TypeUInt:     "int",
	generator.TypeDouble:   "float",
	generator.TypeArray:    "array",
	generator.TypeFunction: "callable",
	generator.TypeObject:   "object",
	generator.TypeVoid:     "void",
}

type Generator struct {
	generator.Base
}

type classContext struct {
	namespace ast.Namespace
	classNode ast.Class
}

type renderArtifacts struct {
	name       string
	parent     string
	doc        string
	attributes string
	hasAttrs   bool
}

func init() {
	generator.Register(generatorName,
		func(ctx context.Context, document ast.Document, options generator.Options) generator.Generator {
			instance := &Generator{Base: generator.NewBase(phpTypeMap)}
			instance.Initialize(ctx, document, options)
			return instance
		})
}

func (Generator) Name() string {
	return generatorName
}

func (phpGen *Generator) Generate() error {
	document := phpGen.Document()
	if !generator.ShouldGenerateNamespace(document.Namespace, "com") {
		return nil
	}
	outputDir := phpGen.OutputDir()
	namespace := document.Namespace
	handler := func(classNode ast.Class) error {
		return phpGen.generateClass(namespace, outputDir, classNode)
	}
	return phpGen.ForEachClass(handler)
}

func (phpGen *Generator) generateClass(namespace ast.Namespace, outputDir string, classNode ast.Class) error {
	var result error
	ctx := classContext{namespace: namespace, classNode: classNode}
	classSource, buildErr := renderClass(ctx)
	if buildErr != nil {
		result = buildErr
		return result
	}
	targetPath := phpTargetPath(outputDir, namespace, classNode.Name)
	phpGen.WriteFile(targetPath, []byte(classSource))
	return result
}

func renderClass(ctx classContext) (string, error) {
	var result string
	docBlock, docErr := buildDocBlock(ctx)
	if docErr != nil {
		return result, docErr
	}
	parentClause := phpParentClause(ctx)
	propertyAttributes, hasAttributes, attrErr := buildPropertyAttributes(ctx.classNode)
	if attrErr != nil {
		return result, attrErr
	}
	header := buildHeader(ctx)
	data := renderArtifacts{
		name:       ctx.classNode.Name,
		parent:     parentClause,
		doc:        docBlock,
		attributes: propertyAttributes,
		hasAttrs:   hasAttributes,
	}
	body := buildClassBody(data)
	result = header + body
	return result, nil
}

func buildHeader(ctx classContext) string {
	var result string
	header := new(strings.Builder)
	header.WriteString("<?php\n")
	header.WriteString("/**\n* PDL Compiler generated code\n")
	header.WriteString("* class " + phpHeaderName(ctx.namespace, ctx.classNode.Name) + "\n")
	header.WriteString("*/\n\n")
	namespaceQualified := phpNamespaceQualified(ctx.namespace)
	if namespaceQualified != "" {
		header.WriteString("namespace " + namespaceQualified + ";\n\n\n")
	}
	result = header.String()
	return result
}

func buildClassBody(data renderArtifacts) string {
	var result string
	body := new(strings.Builder)
	body.WriteString(data.doc)
	body.WriteString("class " + data.name)
	if data.parent != "" {
		body.WriteString(" " + data.parent)
	}
	body.WriteString(" \n{\n\n")
	body.WriteString("    public function __construct()\n")
	body.WriteString("    {\n")
	if data.hasAttrs {
		body.WriteString(data.attributes + "\n")
	}
	body.WriteString("    }\n\n")
	body.WriteString("}\n")
	result = body.String()
	return result
}

func buildDocBlock(ctx classContext) (string, error) {
	var result string
	doc := new(strings.Builder)
	doc.WriteString("/**\n")
	doc.WriteString(" * @class " + ctx.classNode.Name + "\n")
	doc.WriteString(" *\n")
	namespaceQualified := phpNamespaceQualified(ctx.namespace)
	if namespaceQualified != "" {
		doc.WriteString(" * @package " + namespaceQualified + "\n")
		doc.WriteString(" *\n")
	}
	parentDoc := phpParentDoc(ctx)
	if parentDoc != "" {
		doc.WriteString(" * " + parentDoc + "\n")
		doc.WriteString(" *\n")
	}
	propertyLines, propertyErr := buildPropertyDocLines(ctx.classNode)
	if propertyErr != nil {
		return result, propertyErr
	}
	for _, line := range propertyLines {
		doc.WriteString(line + "\n")
	}
	doc.WriteString(" *\n")
	doc.WriteString(" */\n")
	result = doc.String()
	return result, nil
}

func buildPropertyDocLines(classNode ast.Class) ([]string, error) {
	result := make([]string, 0)
	for _, member := range classNode.Members {
		if !generator.IsPropertyMember(member) {
			continue
		}
		propertyType, typeErr := member.PropertyType()
		if typeErr != nil {
			return result, typeErr
		}
		docType := phpDocType(propertyType)
		result = append(result, " * @property "+docType+" $"+member.Name)
	}
	return result, nil
}

func buildPropertyAttributes(classNode ast.Class) (string, bool, error) {
	var result string
	members := collectMembersWithAttributes(classNode)
	if len(members) == 0 {
		return result, false, nil
	}
	builder := new(strings.Builder)
	builder.WriteString("        $this->_propertyAttributes = [\n")
	for index, member := range members {
		builder.WriteString("            '" + member.Name + "' => [ ")
		for attrIndex, attribute := range member.Attributes {
			if attrIndex > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString("'" + attribute.Name.QualifiedName + "' => [\n")
			valueLines := phpAttributeValueLines(attribute)
			for _, line := range valueLines {
				builder.WriteString("                " + line + "\n")
			}
			builder.WriteString("            ]")
		}
		builder.WriteString(" ]")
		if index < len(members)-1 {
			builder.WriteString(",")
		}
		builder.WriteString("\n")
	}
	builder.WriteString("        ];")
	result = builder.String()
	return result, true, nil
}

func collectMembersWithAttributes(classNode ast.Class) []ast.Member {
	result := make([]ast.Member, 0)
	for _, member := range classNode.Members {
		if len(member.Attributes) == 0 {
			continue
		}
		result = append(result, member)
	}
	return result
}

func phpAttributeValueLines(attribute ast.Attribute) []string {
	result := make([]string, 0)
	for index, value := range attribute.Params.Required {
		key := "default" + strconv.Itoa(index+1)
		result = append(result, "'"+key+"' => "+phpLiteral(value))
	}
	for _, named := range attribute.Params.Optional {
		result = append(result, "'"+named.Name+"' => "+phpLiteral(named.Value))
	}
	return result
}

func phpDocType(propertyType ast.PropertyType) string {
	var result string
	baseType := phpTypeName(propertyType.Type)
	if len(propertyType.ArrayNotation) == 0 {
		result = baseType
		return result
	}
	builder := new(strings.Builder)
	builder.WriteString(baseType)
	for range propertyType.ArrayNotation {
		builder.WriteString("[]")
	}
	result = builder.String()
	return result
}

func phpLiteral(value interface{}) string {
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
		result = strconv.Quote(fmt.Sprint(value))
	}
	return result
}

func phpNamespaceQualified(namespace ast.Namespace) string {
	var result string
	if len(namespace.Name.Segments) == 0 {
		return result
	}
	segments := make([]string, 0, len(namespace.Name.Segments))
	for _, segment := range namespace.Name.Segments {
		segments = append(segments, stringy.New(segment).PascalCase().Get())
	}
	result = strings.Join(segments, "\\")
	return result
}

func phpHeaderName(namespace ast.Namespace, className string) string {
	var result string
	segments := make([]string, 0, len(namespace.Name.Segments))
	for _, segment := range namespace.Name.Segments {
		segments = append(segments, stringy.New(segment).PascalCase().Get())
	}
	if len(segments) > 0 {
		segments = append(segments, className)
		result = strings.Join(segments, ".")
		return result
	}
	result = className
	return result
}

func phpParentClause(ctx classContext) string {
	var result string
	if ctx.classNode.Parent == nil {
		return result
	}
	result = "extends " + phpIdentifierReference(ctx.namespace, *ctx.classNode.Parent)
	return result
}

func phpParentDoc(ctx classContext) string {
	var result string
	if ctx.classNode.Parent == nil {
		return result
	}
	parent := phpIdentifierReference(ctx.namespace, *ctx.classNode.Parent)
	result = "@extends " + parent
	return result
}

func phpIdentifierReference(namespace ast.Namespace, identifier ast.Identifier) string {
	var result string
	if len(identifier.Segments) == 0 {
		result = identifier.QualifiedName
		return result
	}
	if strings.Contains(identifier.QualifiedName, ".") {
		result = strings.ReplaceAll(identifier.QualifiedName, ".", "\\")
		return result
	}
	if len(identifier.Segments) == 1 {
		result = stringy.New(identifier.Segments[0]).PascalCase().Get()
		return result
	}
	segments := make([]string, 0, len(identifier.Segments))
	for _, segment := range identifier.Segments {
		segments = append(segments, stringy.New(segment).PascalCase().Get())
	}
	result = strings.Join(segments, "\\")
	return result
}

func phpTypeName(identifier ast.Identifier) string {
	pdlType := identifier.QualifiedName
	if phpType, exists := phpTypeMap[pdlType]; exists {
		return phpType
	}
	if len(identifier.Segments) == 0 {
		return pdlType
	}
	segments := make([]string, 0, len(identifier.Segments))
	for _, segment := range identifier.Segments {
		segments = append(segments, stringy.New(segment).PascalCase().Get())
	}
	return strings.Join(segments, "\\")
}

func phpTargetPath(baseDir string, namespace ast.Namespace, className string) string {
	var result string
	namespacePath := phpNamespacePath(namespace)
	targetRoot := filepath.Join(baseDir, "src")
	if namespacePath != "" {
		targetRoot = filepath.Join(targetRoot, namespacePath)
	}
	result = filepath.Join(targetRoot, className+".php")
	return result
}

func phpNamespacePath(namespace ast.Namespace) string {
	var result string
	if len(namespace.Name.Segments) == 0 {
		return result
	}
	segments := make([]string, 0, len(namespace.Name.Segments))
	for _, segment := range namespace.Name.Segments {
		segments = append(segments, stringy.New(segment).PascalCase().Get())
	}
	result = filepath.Join(segments...)
	return result
}
