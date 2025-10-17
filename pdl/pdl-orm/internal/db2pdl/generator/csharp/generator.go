package csharp

import (
	"path/filepath"
	"strings"

	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

var _ generator.Generator = (*Generator)(nil)

type Generator struct {
	renderer generator.Renderer
	write    generator.CodeWriter
}

func New(renderer generator.Renderer, writer generator.CodeWriter) Generator {
	return Generator{
		renderer: renderer,
		write:    writer,
	}
}

func (generator Generator) Generate(table shared.TableData) error {
	renderTable := generator.InjectAttributes(table)
	source, renderErr := generator.renderer.Render(filepath.Join("cs", "row"), renderTable)
	if renderErr != nil {
		return renderErr
	}
	return generator.write("cs", renderTable.RowClass, source)
}

func (generator Generator) InjectAttributes(table shared.TableData) shared.TableData {
	result := gen.CloneTable(table)
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		field.CSharpType = resolveCSharpType(result.DatabaseDriver, field.DbType)
	}
	return result
}

func resolveCSharpType(driver string, dbType string) string {
	var mapper map[string]string
	switch strings.ToLower(driver) {
	case "mysql":
		mapper = mysqlTypes
	case "postgres", "postgresql":
		mapper = postgresTypes
	default:
		return "string?"
	}
	if value, ok := mapper[dbType]; ok {
		return value
	}
	return "string?"
}

var mysqlTypes = map[string]string{
	"int":        "int?",
	"tinyint":    "int?",
	"smallint":   "short?",
	"mediumint":  "int?",
	"numeric":    "double?",
	"bigint":     "long?",
	"float":      "double?",
	"double":     "double?",
	"decimal":    "double?",
	"bit":        "bool?",
	"char":       "string?",
	"varchar":    "string?",
	"tinytext":   "string?",
	"text":       "string?",
	"mediumtext": "string?",
	"longtext":   "string?",
	"binary":     "byte[]?",
	"varbinary":  "byte[]?",
	"tinyblob":   "byte[]?",
	"blob":       "byte[]?",
	"mediumblob": "byte[]?",
	"longblob":   "byte[]?",
	"enum":       "string?",
	"date":       "DateTime?",
	"datetime":   "DateTime?",
	"time":       "TimeSpan?",
	"timestamp":  "DateTime?",
	"year":       "int?",
}

var postgresTypes = map[string]string{
	"boolean":                     "bool?",
	"bool":                        "bool?",
	"smallint":                    "short?",
	"integer":                     "int?",
	"int":                         "int?",
	"bigint":                      "long?",
	"serial":                      "int?",
	"bigserial":                   "long?",
	"numeric":                     "double?",
	"decimal":                     "double?",
	"real":                        "double?",
	"double precision":            "double?",
	"money":                       "double?",
	"character varying":           "string?",
	"varchar":                     "string?",
	"character":                   "string?",
	"char":                        "string?",
	"text":                        "string?",
	"citext":                      "string?",
	"uuid":                        "string?",
	"json":                        "string?",
	"jsonb":                       "string?",
	"date":                        "DateTime?",
	"timestamp":                   "DateTime?",
	"timestamp without time zone": "DateTime?",
	"timestamp with time zone":    "DateTime?",
	"time":                        "TimeSpan?",
	"time without time zone":      "TimeSpan?",
	"time with time zone":         "TimeSpan?",
	"interval":                    "string?",
	"bytea":                       "byte[]?",
	"inet":                        "string?",
	"vector":                      "double[]?",
}
