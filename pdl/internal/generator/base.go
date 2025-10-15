package generator

import (
    "context"
    "os"
    "path/filepath"
    "strings"

    "github.com/gobeam/stringy"

    "github.com/kapablanka/pdl/pdl/internal/ast"
)


type Base struct {
    typeMap     map[string]string
    context     context.Context
    document    ast.Document
    outputDir   string
    templateDir string
    configPath  string
}


func NewBase(typeMap map[string]string) Base {
	return Base{typeMap: typeMap}
}


func (base *Base) Initialize(ctx context.Context, document ast.Document, options Options) {
    base.context = ctx
    base.document = document
    base.outputDir = strings.TrimSpace(options.OutputDir)
    base.templateDir = strings.TrimSpace(options.TemplateDir)
    base.configPath = strings.TrimSpace(options.ConfigPath)
}


func (base Base) Context() context.Context {
	return base.context
}

func (base Base) Document() ast.Document {
	return base.document
}

func (base Base) OutputDir() string {
	return base.outputDir
}

func (base Base) TemplateDir() string {
	return base.templateDir
}


func (base Base) ConfigPath() string {
	return base.configPath
}


func (Base) EnsureDir(targetDir string) error {
	if targetDir == "" {
		return nil
	}
	return os.MkdirAll(targetDir, 0o755)
}


func (Base) WriteFile(targetPath string, contents []byte) {
	parentDir := filepath.Dir(targetPath)
	if err := os.MkdirAll(parentDir, 0o755); err != nil {
		panic(err)
	}
	if err := os.WriteFile(targetPath, contents, 0o644); err != nil {
		panic(err)
	}
}

func (base Base) ForEachClass(handler func(ast.Class) error) error {
	if handler == nil {
		return nil
	}
	for _, classNode := range base.document.Namespace.Classes {
		if base.context != nil {
			if err := base.context.Err(); err != nil {
				return err
			}
		}
		if err := handler(classNode); err != nil {
			return err
		}
	}
	return nil
}

func ShouldGenerateNamespace(namespace ast.Namespace, expectedRoot string) bool {
	if expectedRoot == "" {
		return false
	}
	if len(namespace.Name.Segments) == 0 {
		return false
	}
	return strings.EqualFold(namespace.Name.Segments[0], expectedRoot)
}

func IsPropertyMember(member ast.Member) bool {
	return member.Kind == "property" || member.Kind == "shortProperty"
}

func CamelCase(value string) string {
	return stringy.New(value).CamelCase().Get()
}


func SnakeCaseUpper(value string) string {
	return strings.ToUpper(stringy.New(value).SnakeCase().Get())
}


func Capitalize(value string) string {
	return stringy.New(value).UcFirst()
}

func SimpleName(qualified string) string {
	segments := strings.Split(qualified, ".")
	return Capitalize(segments[len(segments)-1])
}

func (base Base) FromPdlType(dataType string) (string, bool) {
	result, ok := base.typeMap[dataType]
	return result, ok
}
