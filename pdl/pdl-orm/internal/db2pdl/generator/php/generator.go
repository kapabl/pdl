package php

import (
	"path/filepath"
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
	source, renderErr := generator.renderer.Render(filepath.Join("php", "row"), renderTable)
	if renderErr != nil {
		return renderErr
	}
	return generator.write("php", renderTable.RowClass, source)
}

func (generator Generator) InjectAttributes(table shared.TableData) shared.TableData {
	return generator.applyTypes(table)
}

func (generator Generator) applyTypes(table shared.TableData) shared.TableData {
	result := gen.CloneTable(table)
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		field.PhpType = resolvePhpType(result.DatabaseDriver, field.DbType)
	}
	return result
}

func resolvePhpType(driver string, dbType string) string {
	var mapper map[string]string
	switch strings.ToLower(driver) {
	case "mysql":
		mapper = mysqlPhpTypes
	case "postgres", "postgresql":
		mapper = postgresPhpTypes
	default:
		return "string"
	}
	if value, found := mapper[dbType]; found {
		return value
	}
	return "string"
}

var mysqlPhpTypes = map[string]string{
	"int":        "int",
	"tinyint":    "int",
	"smallint":   "int",
	"mediumint":  "int",
	"numeric":    "int",
	"bigint":     "int",
	"float":      "float",
	"double":     "float",
	"decimal":    "float",
	"bit":        "int",
	"char":       "string",
	"varchar":    "string",
	"tinytext":   "string",
	"text":       "string",
	"mediumtext": "string",
	"longtext":   "string",
	"binary":     "string",
	"varbinary":  "string",
	"tinyblob":   "string",
	"blob":       "string",
	"mediumblob": "string",
	"longblob":   "string",
	"enum":       "string",
	"date":       "string",
	"datetime":   "string",
	"time":       "string",
	"timestamp":  "string",
	"year":       "int",
}

var postgresPhpTypes = map[string]string{
	"boolean":                     "bool",
	"bool":                        "bool",
	"smallint":                    "int",
	"integer":                     "int",
	"int":                         "int",
	"bigint":                      "int",
	"serial":                      "int",
	"bigserial":                   "int",
	"numeric":                     "float",
	"decimal":                     "float",
	"real":                        "float",
	"double precision":            "float",
	"money":                       "float",
	"character varying":           "string",
	"varchar":                     "string",
	"character":                   "string",
	"char":                        "string",
	"text":                        "string",
	"citext":                      "string",
	"uuid":                        "string",
	"json":                        "string",
	"jsonb":                       "string",
	"date":                        "string",
	"timestamp":                   "string",
	"timestamp without time zone": "string",
	"timestamp with time zone":    "string",
	"time":                        "string",
	"time without time zone":      "string",
	"time with time zone":         "string",
	"interval":                    "string",
	"bytea":                       "string",
	"inet":                        "string",
	"vector":                      "string",
}
