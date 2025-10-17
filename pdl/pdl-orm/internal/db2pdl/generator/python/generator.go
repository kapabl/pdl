package python

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
	config   shared.DB2PDLConfig
}

func New(renderer gen.Renderer, writer gen.CodeWriter, config shared.DB2PDLConfig) Generator {
	return Generator{
		renderer: renderer,
		write:    writer,
		config:   config,
	}
}

func (generator Generator) Generate(table shared.TableData) error {
	renderTable, imports := generator.injectAttributes(table)
	payload := struct {
		Table   shared.TableData
		Imports pythonImports
	}{
		Table:   renderTable,
		Imports: imports,
	}
	modulePath := strings.ReplaceAll(renderTable.PythonPackage, ".", string(filepath.Separator))
	targetName := renderTable.RowClass
	if modulePath != "" {
		targetName = filepath.Join(modulePath, targetName)
	}
	source, renderErr := generator.renderer.Render(filepath.Join("python", "row"), payload)
	if renderErr != nil {
		return renderErr
	}
	return generator.write("py", targetName, source)
}

func (generator Generator) injectAttributes(table shared.TableData) (shared.TableData, pythonImports) {
	result := gen.CloneTable(table)
	imports := pythonImports{}
	for index := range result.FieldsInfo {
		field := &result.FieldsInfo[index]
		field.PythonType = resolvePythonType(result.DatabaseDriver, field.DbType)
		imports.include(field.PythonType)
	}
	return result, imports
}

func (generator Generator) InjectAttributes(table shared.TableData) shared.TableData {
	result, _ := generator.injectAttributes(table)
	return result
}

type pythonImports struct {
	NeedsDatetime bool
	NeedsDecimal  bool
	NeedsUUID     bool
}

func (flags *pythonImports) include(pythonType string) {
	if strings.HasPrefix(pythonType, "datetime.") {
		flags.NeedsDatetime = true
	}
	if pythonType == "Decimal" {
		flags.NeedsDecimal = true
	}
	if pythonType == "uuid.UUID" {
		flags.NeedsUUID = true
	}
}

func resolvePythonType(driver string, dbType string) string {
	switch strings.ToLower(driver) {
	case "mysql":
		return mysqlPythonType(dbType)
	case "postgres", "postgresql":
		return postgresPythonType(dbType)
	default:
		return "str"
	}
}

func mysqlPythonType(dbType string) string {
	clean := strings.ToLower(dbType)
	if value, found := mysqlTypes[clean]; found {
		return value
	}
	return "str"
}

func postgresPythonType(dbType string) string {
	clean := strings.ToLower(dbType)
	if value, found := postgresTypes[clean]; found {
		return value
	}
	return "str"
}

var mysqlTypes = map[string]string{
	"int":        "int",
	"tinyint":    "int",
	"smallint":   "int",
	"mediumint":  "int",
	"numeric":    "Decimal",
	"bigint":     "int",
	"float":      "float",
	"double":     "float",
	"decimal":    "Decimal",
	"bit":        "int",
	"char":       "str",
	"varchar":    "str",
	"tinytext":   "str",
	"text":       "str",
	"mediumtext": "str",
	"longtext":   "str",
	"binary":     "bytes",
	"varbinary":  "bytes",
	"tinyblob":   "bytes",
	"blob":       "bytes",
	"mediumblob": "bytes",
	"longblob":   "bytes",
	"enum":       "str",
	"date":       "datetime.date",
	"datetime":   "datetime.datetime",
	"time":       "datetime.time",
	"timestamp":  "datetime.datetime",
	"year":       "int",
}

var postgresTypes = map[string]string{
	"boolean":                     "bool",
	"bool":                        "bool",
	"smallint":                    "int",
	"integer":                     "int",
	"int":                         "int",
	"bigint":                      "int",
	"serial":                      "int",
	"bigserial":                   "int",
	"numeric":                     "Decimal",
	"decimal":                     "Decimal",
	"real":                        "float",
	"double precision":            "float",
	"money":                       "Decimal",
	"character varying":           "str",
	"varchar":                     "str",
	"character":                   "str",
	"char":                        "str",
	"text":                        "str",
	"citext":                      "str",
	"uuid":                        "uuid.UUID",
	"json":                        "str",
	"jsonb":                       "str",
	"date":                        "datetime.date",
	"timestamp":                   "datetime.datetime",
	"timestamp without time zone": "datetime.datetime",
	"timestamp with time zone":    "datetime.datetime",
	"time":                        "datetime.time",
	"time without time zone":      "datetime.time",
	"time with time zone":         "datetime.time",
	"bytea":                       "bytes",
}
