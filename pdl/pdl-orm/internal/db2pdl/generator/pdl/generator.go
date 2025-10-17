package pdl

import (
	"path/filepath"
	"strings"

	gen "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

var _ gen.Generator = (*Generator)(nil)

const (
	dbIDAttribute       = "[io.pdl.infrastructure.data.attributes.IsDbId]"
	columnNameAttribute = "[io.pdl.infrastructure.data.attributes.ColumnName(\"{$column_name}\")]"
)

type Generator struct {
	renderer gen.Renderer
	write    gen.CodeWriter
}

func New(renderer gen.Renderer, writer gen.CodeWriter) *Generator {
	return &Generator{
		renderer: renderer,
		write:    writer,
	}
}

func (generator *Generator) Generate(table shared.TableData) error {
	renderTable := generator.InjectAttributes(table)
	source, renderErr := generator.renderer.Render(filepath.Join("pdl", "pdlRowClass"), renderTable)
	if renderErr != nil {
		return renderErr
	}
	return generator.write("pdl", renderTable.PdlRowClass, source)
}

func (generator *Generator) InjectAttributes(table shared.TableData) shared.TableData {
	table = generator.applyTypes(table)
	config := gen.AttributeConfig{
		PrimaryTemplate: dbIDAttribute,
		ColumnTemplate:  columnNameAttribute,
	}
	return gen.ApplyAttributes(table, config, func(field *shared.FieldInfo, tags string) {
		field.PdlAttributes = tags
	})
}

func (generator *Generator) applyTypes(table shared.TableData) shared.TableData {
	result := gen.CloneTable(table)
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		field.PdlType = resolvePdlType(result.DatabaseDriver, field.DbType)
	}
	return result
}

func resolvePdlType(driver string, dbType string) string {
	var mapper map[string]string
	switch strings.ToLower(driver) {
	case "mysql":
		mapper = mysqlPdlTypes
	case "postgres", "postgresql":
		mapper = postgresPdlTypes
	default:
		return "string"
	}
	if value, found := mapper[dbType]; found {
		return value
	}
	return "string"
}

var mysqlPdlTypes = map[string]string{
	"int":        "int",
	"tinyint":    "int",
	"smallint":   "int",
	"mediumint":  "int",
	"numeric":    "int",
	"bigint":     "int",
	"float":      "double",
	"double":     "double",
	"decimal":    "double",
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

var postgresPdlTypes = map[string]string{
	"boolean":                     "bool",
	"bool":                        "bool",
	"smallint":                    "int",
	"integer":                     "int",
	"int":                         "int",
	"bigint":                      "int",
	"serial":                      "int",
	"bigserial":                   "int",
	"numeric":                     "double",
	"decimal":                     "double",
	"real":                        "double",
	"double precision":            "double",
	"money":                       "double",
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
