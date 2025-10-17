package cpp

import (
	"path/filepath"
	"sort"
	"strings"

	gen "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/mysql"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/postgres"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

var _ gen.Generator = (*Generator)(nil)

type Generator struct {
	renderer gen.Renderer
	write    gen.CodeWriter
}

type cppField struct {
	shared.FieldInfo
	MemberName string
}

func New(renderer gen.Renderer, writer gen.CodeWriter) Generator {
	return Generator{
		renderer: renderer,
		write:    writer,
	}
}

func (generator Generator) Generate(table shared.TableData) error {
	renderTable := generator.InjectAttributes(table)
	fields := buildFields(renderTable)
	payload := struct {
		Table         shared.TableData
		Fields        []cppField
		Includes      []string
		Namespace     string
		AccessorType  string
		AccessorConst string
	}{
		Table:         renderTable,
		Fields:        fields,
		Includes:      collectIncludes(fields),
		Namespace:     namespacePath(renderTable.PdlEntitiesNamespace),
		AccessorType:  accessorTypeName(renderTable.Name),
		AccessorConst: accessorConstName(renderTable.Name),
	}
	source, renderErr := generator.renderer.Render(filepath.Join("cpp", "row"), payload)
	if renderErr != nil {
		return renderErr
	}
	target := targetFilePath(renderTable)
	return generator.write("", target, source)
}

func (generator Generator) InjectAttributes(table shared.TableData) shared.TableData {
	result := gen.CloneTable(table)
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		field.CppType = resolveCppType(result.DatabaseDriver, field.DbType)
	}
	return result
}

func resolveCppType(driver string, dbType string) string {
	switch strings.ToLower(driver) {
	case "mysql":
		return mysql.CppType(dbType)
	case "postgres", "postgresql":
		return postgres.CppType(dbType)
	default:
		return "std::string"
	}
}

func buildFields(table shared.TableData) []cppField {
	result := make([]cppField, len(table.FieldsInfo))
	for index, field := range table.FieldsInfo {
		member := field.PascalCase
		if member == "" {
			member = shared.ToPascalCase(field.Original)
		}
		result[index] = cppField{
			FieldInfo:  field,
			MemberName: member,
		}
	}
	return result
}

func collectIncludes(fields []cppField) []string {
	includes := map[string]struct{}{
		"<pdl/infrastructure/data/DBStore.hpp>":      {},
		"<pdl/infrastructure/data/QueryBuilder.hpp>": {},
		"<pdl/infrastructure/data/Row.hpp>":          {},
		"<pdl/infrastructure/data/RowExecutor.hpp>":  {},
		"<pdl/infrastructure/data/Types.hpp>":        {},
		"<initializer_list>":                         {},
		"<memory>":                                   {},
		"<optional>":                                 {},
		"<string>":                                   {},
		"<vector>":                                   {},
		"<cstdint>":                                  {},
		"<utility>":                                 {},
	}
	for _, field := range fields {
		if strings.HasPrefix(field.CppType, "std::vector") {
			includes["<vector>"] = struct{}{}
		}
	}
	values := make([]string, 0, len(includes))
	for include := range includes {
		values = append(values, include)
	}
	sort.Strings(values)
	return values
}

func namespacePath(namespace string) string {
	if namespace == "" {
		return ""
	}
	parts := strings.Split(namespace, ".")
	for index, part := range parts {
		parts[index] = strings.ToLower(strings.TrimSpace(part))
	}
	return strings.Join(parts, "::")
}

func accessorTypeName(base string) string {
	if base == "" {
		return "Accessor"
	}
	return base + "Accessor"
}

func accessorConstName(base string) string {
	snake := shared.ToSnakeCaseFromPascal(base)
	if snake == "" {
		snake = strings.ToLower(base)
	}
	return strings.ToUpper(strings.ReplaceAll(snake, "-", "_"))
}

func targetFilePath(table shared.TableData) string {
	basePath := shared.NamespaceToGoPath(table.PdlEntitiesNamespace)
	fileName := table.RowClass
	if fileName == "" {
		fileName = shared.ToPascalCase(table.TableName)
	}
	fileName = fileName + ".hpp"
	if basePath == "" {
		return filepath.Join("cpp", fileName)
	}
	return filepath.Join("cpp", basePath, fileName)
}
