package postgres

var goTypes = map[string]string{
	"boolean":                     "bool",
	"bool":                        "bool",
	"smallint":                    "int",
	"integer":                     "int",
	"int":                         "int",
	"bigint":                      "int64",
	"serial":                      "int",
	"bigserial":                   "int64",
	"numeric":                     "float64",
	"decimal":                     "float64",
	"real":                        "float64",
	"double precision":            "float64",
	"money":                       "float64",
	"character varying":           "string",
	"varchar":                     "string",
	"character":                   "string",
	"char":                        "string",
	"text":                        "string",
	"citext":                      "string",
	"uuid":                        "string",
	"json":                        "string",
	"jsonb":                       "string",
	"date":                        "time.Time",
	"timestamp":                   "time.Time",
	"timestamp without time zone": "time.Time",
	"timestamp with time zone":    "time.Time",
	"time":                        "string",
	"time without time zone":      "string",
	"time with time zone":         "string",
	"interval":                    "string",
	"bytea":                       "[]byte",
	"inet":                        "string",
	"vector":                      "[]float64",
}

var timeImports = map[string]struct{}{
	"date":                        {},
	"timestamp":                   {},
	"timestamp without time zone": {},
	"timestamp with time zone":    {},
}

func GoType(cleanType string) string {
	if value, found := goTypes[cleanType]; found {
		return value
	}
	return "string"
}

func NeedsTimeImport(cleanType string) bool {
	_, ok := timeImports[cleanType]
	return ok
}
