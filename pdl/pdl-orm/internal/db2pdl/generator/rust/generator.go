package rust

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

type rustField struct {
	shared.FieldInfo
	RustName string
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
		Table        shared.TableData
		Fields       []rustField
		Imports      []string
		AccessorName string
	}{
		Table:        renderTable,
		Fields:       fields,
		Imports:      collectImports(),
		AccessorName: accessorConstName(renderTable.Name),
	}
	source, renderErr := generator.renderer.Render(filepath.Join("rust", "row"), payload)
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
		field.RustType = resolveRustType(result.DatabaseDriver, field.DbType)
	}
	return result
}

func resolveRustType(driver string, dbType string) string {
	switch strings.ToLower(driver) {
	case "mysql":
		return mysql.RustType(dbType)
	case "postgres", "postgresql":
		return postgres.RustType(dbType)
	default:
		return "String"
	}
}

func buildFields(table shared.TableData) []rustField {
	result := make([]rustField, len(table.FieldsInfo))
	for index, field := range table.FieldsInfo {
		name := shared.ToSnakeCaseFromPascal(field.PascalCase)
		if name == "" {
			name = strings.ToLower(field.Original)
		}
		result[index] = rustField{
			FieldInfo: field,
			RustName:  name,
		}
	}
	return result
}

func collectImports() []string {
	values := []string{
		"pdl_infrastructure::data::{DbError, DBStore, Operator, OrderDirection, QueryBuilder, Record, Row, RowExecutor}",
	}
	sort.Strings(values)
	return values
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
	fileName := shared.ToSnakeCaseFromPascal(table.RowClass)
	if fileName == "" {
		fileName = strings.ToLower(table.RowClass)
	}
	fileName = fileName + ".rs"
	if basePath == "" {
		return filepath.Join("rust", fileName)
	}
	return filepath.Join("rust", basePath, fileName)
}
