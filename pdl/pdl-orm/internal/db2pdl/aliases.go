package db2pdl

import sharedpkg "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"

type (
	AttributeTemplates = sharedpkg.AttributeTemplates
	PHPConfig          = sharedpkg.PHPConfig
	PDLConfig          = sharedpkg.PDLConfig
	TypeScriptConfig   = sharedpkg.TypeScriptConfig
	CSConfig           = sharedpkg.CSConfig
	GoConfig           = sharedpkg.GoConfig
	ConnectionConfig   = sharedpkg.ConnectionConfig
	DB2PDLConfig       = sharedpkg.DB2PDLConfig
	RootConfig         = sharedpkg.RootConfig
	FieldInfo          = sharedpkg.FieldInfo
	TableData          = sharedpkg.TableData
	VerbosePrinter     = sharedpkg.VerbosePrinter
)

var (
	LoadConfig     = sharedpkg.LoadConfig
	ValidateConfig = sharedpkg.ValidateConfig
)

func NewVerbosePrinter(enabled bool) sharedpkg.VerbosePrinter {
	return sharedpkg.NewVerbosePrinter(enabled)
}

func CleanOutputDirectory(target string) error {
	return sharedpkg.CleanOutputDirectory(target)
}

func NamespaceToPHP(namespace string) string {
	return sharedpkg.NamespaceToPHP(namespace)
}

func NamespaceToGoPath(namespace string) string {
	return sharedpkg.NamespaceToGoPath(namespace)
}

func NamespaceToGoPackage(namespace string) string {
	return sharedpkg.NamespaceToGoPackage(namespace)
}

func ResolveTemplatesDir(overrideValue string) (string, bool) {
	return sharedpkg.ResolveTemplatesDir(overrideValue)
}

func ContainsCaseSensitive(collection []string, value string) bool {
	return sharedpkg.ContainsCaseSensitive(collection, value)
}

func ToPascalCase(value string) string {
	return sharedpkg.ToPascalCase(value)
}

func Singularize(value string) string {
	return sharedpkg.Singularize(value)
}
