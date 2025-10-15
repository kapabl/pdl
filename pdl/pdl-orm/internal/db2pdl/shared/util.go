package shared

import (
	"errors"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type VerbosePrinter struct {
	enabled bool
}

func NewVerbosePrinter(enabled bool) VerbosePrinter {
	return VerbosePrinter{enabled: enabled}
}

func (printer VerbosePrinter) Println(message string) {
	if printer.enabled {
		os.Stdout.WriteString(message + "\n")
	}
}

func CleanOutputDirectory(target string) error {
	if target == "" {
		return errors.New("output directory not specified")
	}
	info, statErr := os.Stat(target)
	if statErr == nil && !info.IsDir() {
		return errors.New("output path exists and is not a directory")
	}
	if statErr == nil {
		oldName := target + "_old"
		suffix := 1
		for {
			renameErr := os.Rename(target, oldName)
			if renameErr == nil {
				break
			}
			if !os.IsExist(renameErr) {
				break
			}
			oldName = target + "_old" + strconv.Itoa(suffix)
			suffix++
		}
		os.RemoveAll(oldName)
	}
	return os.MkdirAll(target, 0o755)
}

func NamespaceToPHP(namespace string) string {
	parts := strings.Split(namespace, ".")
	for index := range parts {
		if len(parts[index]) == 0 {
			continue
		}
		parts[index] = strings.ToUpper(parts[index][:1]) + parts[index][1:]
	}
	return strings.Join(parts, "\\")
}

func NamespaceToGoPath(namespace string) string {
	parts := strings.Split(namespace, ".")
	filtered := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		filtered = append(filtered, strings.ToLower(part))
	}
	return strings.Join(filtered, "/")
}

func NamespaceToGoPackage(namespace string) string {
	path := NamespaceToGoPath(namespace)
	if path == "" {
		return ""
	}
	segments := strings.Split(path, "/")
	return segments[len(segments)-1]
}

func MergeImports(values map[string]struct{}) []string {
	result := make([]string, 0, len(values))
	for value := range values {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func ClearSQLType(sqlType string) string {
	result := strings.Split(sqlType, "(")[0]
	result = strings.TrimSpace(result)
	return strings.ToLower(result)
}

func CantConvertBackAndForth(fieldName string) bool {
	digitPattern := regexp.MustCompile(`\d+`)
	matches := digitPattern.FindAllString(fieldName, -1)
	return matches != nil
}

func GenerateAttribute(field FieldInfo, template string) string {
	result := strings.ReplaceAll(template, "{$columnName}", field.CamelCase)
	result = strings.ReplaceAll(result, "{$column_name}", field.SnakeCase)
	if strings.Contains(result, "{}") {
		result = strings.ReplaceAll(result, "{}", field.SnakeCase)
	}
	return result
}

func ContainsCaseSensitive(collection []string, value string) bool {
	for _, item := range collection {
		if item == value {
			return true
		}
	}
	return false
}

func ToCamelCase(source string) string {
	segments := strings.Split(source, "_")
	for index := range segments {
		segments[index] = strings.ToLower(segments[index])
	}
	for index := range segments {
		if index == 0 || len(segments[index]) == 0 {
			continue
		}
		segments[index] = strings.ToUpper(segments[index][:1]) + segments[index][1:]
	}
	return strings.Join(segments, "")
}

func ToPascalCase(source string) string {
	camel := ToCamelCase(source)
	if len(camel) == 0 {
		return camel
	}
	return strings.ToUpper(camel[:1]) + camel[1:]
}

func SortFieldInfos(fields []FieldInfo) {
	sort.SliceStable(fields, func(left int, right int) bool {
		return fields[left].Original < fields[right].Original
	})
}

func ToSnakeCaseFromPascal(source string) string {
	if source == "" {
		return source
	}
	runes := []rune(source)
	builder := strings.Builder{}
	for index, r := range runes {
		if index > 0 && r >= 'A' && r <= 'Z' {
			builder.WriteRune('_')
		}
		builder.WriteRune(r)
	}
	return strings.ToLower(builder.String())
}

func ResolveTemplatesDir(overrideValue string) (string, bool) {
	if overrideValue != "" && overrideValue != "default" {
		return overrideValue, true
	}
	return "", false
}

func Singularize(source string) string {
	if strings.HasSuffix(source, "ies") {
		return strings.TrimSuffix(source, "ies") + "y"
	}
	if strings.HasSuffix(source, "ses") {
		return strings.TrimSuffix(source, "es")
	}
	if strings.HasSuffix(source, "s") {
		return strings.TrimSuffix(source, "s")
	}
	return source
}

func ComposeTableData(baseName string, tableName string, config DB2PDLConfig, defaultNamespace string, phpNamespace string, fields []FieldInfo, goImports map[string]struct{}, primaryKeyCamel string, primaryKeyPascal string, primaryKeyOriginal string) TableData {
	result := TableData{
		Name:                   baseName,
		TableName:              tableName,
		DbName:                 config.Connection.Database,
		PdlRowClass:            baseName + "Row",
		RowClass:               baseName + "Row",
		GoStruct:               baseName,
		RowSetClass:            baseName + "RowSet",
		ColumnsDefinitionClass: baseName + "Columns",
		WhereClass:             baseName + "Where",
		OrderByClass:           baseName + "OrderBy",
		FieldListClass:         baseName + "FieldList",
		ColumnsListTraits:      baseName + "ColumnsTraits",
		CsharpRowSetClass:      baseName + "RowSet",
		PhpUseNamespaces:       defaultNamespace,
		PhpEntitiesNamespace:   phpNamespace,
		PdlEntitiesNamespace:   config.PDL.EntitiesNamespace,
		PdlUseNamespaces:       config.PDL.UseNamespaces,
		PrimaryKeyCamelCase:    primaryKeyCamel,
		PrimaryKeyPascalCase:   primaryKeyPascal,
		PrimaryKeyOriginal:     primaryKeyOriginal,
		FieldsInfo:             fields,
	}
	result.GoImports = MergeImports(goImports)
	if result.PdlEntitiesNamespace == "" {
		result.PdlEntitiesNamespace = strings.ReplaceAll(defaultNamespace, "/", ".")
	}
	return result
}
