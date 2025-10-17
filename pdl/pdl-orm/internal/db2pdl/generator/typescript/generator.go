package typescript

import (
	"path/filepath"
	"strings"

	gen "github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/generator"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

type Generator struct {
	renderer gen.Renderer
	write    gen.CodeWriter
	blocks   []string
}

func New(renderer gen.Renderer, writer gen.CodeWriter) *Generator {
	return &Generator{
		renderer: renderer,
		write:    writer,
		blocks:   make([]string, 0),
	}
}

func (generator *Generator) AddTable(table shared.TableData) error {
	// TODO: add TypeScript-specific attribute handling if required.
	table = generator.InjectAttributes(table)
	source, renderErr := generator.renderer.Render(filepath.Join("ts", "dbBlock"), table)
	if renderErr != nil {
		return renderErr
	}
	generator.blocks = append(generator.blocks, source)
	return nil
}

func (generator *Generator) Flush(outputFile string) error {
	if len(generator.blocks) == 0 {
		return nil
	}
	payload := map[string]interface{}{
		"source": strings.Join(generator.blocks, "\n"),
	}
	source, renderErr := generator.renderer.Render(filepath.Join("ts", "dbModule"), payload)
	if renderErr != nil {
		return renderErr
	}
	return generator.write("ts", outputFile, source)
}

func (generator *Generator) InjectAttributes(table shared.TableData) shared.TableData {
	result := gen.CloneTable(table)
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		field.TsType = resolveTypeScriptType(result.DatabaseDriver, field.DbType)
	}
	return result
}

func resolveTypeScriptType(driver string, dbType string) string {
	var mapper map[string]string
	switch strings.ToLower(driver) {
	case "mysql":
		mapper = mysqlTypeScriptTypes
	case "postgres", "postgresql":
		mapper = postgresTypeScriptTypes
	default:
		return "string"
	}
	if value, found := mapper[dbType]; found {
		return value
	}
	return "string"
}

var mysqlTypeScriptTypes = map[string]string{
	"int":        "number",
	"tinyint":    "number",
	"smallint":   "number",
	"mediumint":  "number",
	"numeric":    "number",
	"bigint":     "number",
	"float":      "number",
	"double":     "number",
	"decimal":    "number",
	"bit":        "number",
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
	"year":       "number",
}

var postgresTypeScriptTypes = map[string]string{
	"boolean":                     "boolean",
	"bool":                        "boolean",
	"smallint":                    "number",
	"integer":                     "number",
	"int":                         "number",
	"bigint":                      "number",
	"serial":                      "number",
	"bigserial":                   "number",
	"numeric":                     "number",
	"decimal":                     "number",
	"real":                        "number",
	"double precision":            "number",
	"money":                       "number",
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
