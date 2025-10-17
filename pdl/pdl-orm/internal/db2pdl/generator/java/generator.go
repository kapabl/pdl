package java

import (
	"path/filepath"
	"sort"
	"strings"

	gen "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

var _ gen.Generator = (*Generator)(nil)

type Generator struct {
	renderer gen.Renderer
	write    gen.CodeWriter
}

func New(renderer gen.Renderer, writer gen.CodeWriter) Generator {
	return Generator{
		renderer: renderer,
		write:    writer,
	}
}

func (generator Generator) Generate(table shared.TableData) error {
	renderTable := generator.InjectAttributes(table)
	imports := collectImports(renderTable)
	payload := struct {
		Table   shared.TableData
		Imports []string
	}{
		Table:   renderTable,
		Imports: imports,
	}
	source, renderErr := generator.renderer.Render(filepath.Join("java", "row"), payload)
	if renderErr != nil {
		return renderErr
	}
	relative := strings.ReplaceAll(renderTable.JavaPackage, ".", string(filepath.Separator))
	targetName := renderTable.RowClass
	if relative != "" {
		targetName = filepath.Join(relative, targetName)
	}
	return generator.write("java", targetName, source)
}

func (generator Generator) InjectAttributes(table shared.TableData) shared.TableData {
	result := gen.CloneTable(table)
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		field.JavaType = resolveJavaType(result.DatabaseDriver, field.DbType)
	}
	return result
}

func collectImports(table shared.TableData) []string {
	imports := map[string]struct{}{
		"io.pdl.infrastructure.data.DBStore":               {},
		"io.pdl.infrastructure.data.Operator":              {},
		"io.pdl.infrastructure.data.QueryBuilder":          {},
		"io.pdl.infrastructure.data.Row":                   {},
		"io.pdl.infrastructure.data.RowExecutor":           {},
		"io.pdl.infrastructure.data.annotations.PdlColumn": {},
		"java.util.ArrayList":                              {},
		"java.util.List":                                   {},
		"java.util.Map":                                    {},
	}
	for _, field := range table.FieldsInfo {
		switch field.JavaType {
		case "LocalDate":
			imports["java.time.LocalDate"] = struct{}{}
		case "LocalDateTime":
			imports["java.time.LocalDateTime"] = struct{}{}
		case "LocalTime":
			imports["java.time.LocalTime"] = struct{}{}
		}
	}
	values := make([]string, 0, len(imports))
	for entry := range imports {
		values = append(values, entry)
	}
	sort.Strings(values)
	return values
}

func resolveJavaType(driver string, dbType string) string {
	var mapper map[string]string
	switch strings.ToLower(driver) {
	case "mysql":
		mapper = mysqlJavaTypes
	case "postgres", "postgresql":
		mapper = postgresJavaTypes
	default:
		return "String"
	}
	if value, found := mapper[dbType]; found {
		return value
	}
	return "String"
}

var mysqlJavaTypes = map[string]string{
	"int":        "Integer",
	"tinyint":    "Integer",
	"smallint":   "Integer",
	"mediumint":  "Integer",
	"numeric":    "Double",
	"bigint":     "Long",
	"float":      "Double",
	"double":     "Double",
	"decimal":    "Double",
	"bit":        "Integer",
	"char":       "String",
	"varchar":    "String",
	"tinytext":   "String",
	"text":       "String",
	"mediumtext": "String",
	"longtext":   "String",
	"binary":     "byte[]",
	"varbinary":  "byte[]",
	"tinyblob":   "byte[]",
	"blob":       "byte[]",
	"mediumblob": "byte[]",
	"longblob":   "byte[]",
	"enum":       "String",
	"date":       "LocalDate",
	"datetime":   "LocalDateTime",
	"time":       "LocalTime",
	"timestamp":  "LocalDateTime",
	"year":       "Integer",
}

var postgresJavaTypes = map[string]string{
	"boolean":                     "Boolean",
	"bool":                        "Boolean",
	"smallint":                    "Integer",
	"integer":                     "Integer",
	"int":                         "Integer",
	"bigint":                      "Long",
	"serial":                      "Integer",
	"bigserial":                   "Long",
	"numeric":                     "Double",
	"decimal":                     "Double",
	"real":                        "Double",
	"double precision":            "Double",
	"money":                       "Double",
	"character varying":           "String",
	"varchar":                     "String",
	"character":                   "String",
	"char":                        "String",
	"text":                        "String",
	"citext":                      "String",
	"uuid":                        "String",
	"json":                        "String",
	"jsonb":                       "String",
	"date":                        "LocalDate",
	"timestamp":                   "LocalDateTime",
	"timestamp without time zone": "LocalDateTime",
	"timestamp with time zone":    "LocalDateTime",
	"time":                        "LocalTime",
	"time without time zone":      "LocalTime",
	"time with time zone":         "LocalTime",
	"interval":                    "String",
	"bytea":                       "byte[]",
	"inet":                        "String",
	"vector":                      "double[]",
}
