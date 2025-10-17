package postgres

var rustTypes = map[string]string{
	"boolean":                     "bool",
	"bool":                        "bool",
	"smallint":                    "i16",
	"integer":                     "i32",
	"int":                         "i32",
	"bigint":                      "i64",
	"serial":                      "i32",
	"bigserial":                   "i64",
	"numeric":                     "f64",
	"decimal":                     "f64",
	"real":                        "f32",
	"double precision":            "f64",
	"money":                       "f64",
	"character varying":           "String",
	"varchar":                     "String",
	"character":                   "String",
	"char":                        "String",
	"text":                        "String",
	"citext":                      "String",
	"uuid":                        "String",
	"json":                        "String",
	"jsonb":                       "String",
	"date":                        "String",
	"timestamp":                   "String",
	"timestamp without time zone": "String",
	"timestamp with time zone":    "String",
	"time":                        "String",
	"time without time zone":      "String",
	"time with time zone":         "String",
	"interval":                    "String",
	"bytea":                       "Vec<u8>",
	"inet":                        "String",
	"vector":                      "Vec<f64>",
}

func RustType(dbType string) string {
	if value, found := rustTypes[dbType]; found {
		return value
	}
	return "String"
}
