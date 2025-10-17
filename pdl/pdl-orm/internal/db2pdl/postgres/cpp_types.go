package postgres

var cppTypes = map[string]string{
	"boolean":                     "bool",
	"bool":                        "bool",
	"smallint":                    "std::int16_t",
	"integer":                     "std::int32_t",
	"int":                         "std::int32_t",
	"bigint":                      "std::int64_t",
	"serial":                      "std::int32_t",
	"bigserial":                   "std::int64_t",
	"numeric":                     "double",
	"decimal":                     "double",
	"real":                        "float",
	"double precision":            "double",
	"money":                       "double",
	"character varying":           "std::string",
	"varchar":                     "std::string",
	"character":                   "std::string",
	"char":                        "std::string",
	"text":                        "std::string",
	"citext":                      "std::string",
	"uuid":                        "std::string",
	"json":                        "std::string",
	"jsonb":                       "std::string",
	"date":                        "std::string",
	"timestamp":                   "std::string",
	"timestamp without time zone": "std::string",
	"timestamp with time zone":    "std::string",
	"time":                        "std::string",
	"time without time zone":      "std::string",
	"time with time zone":         "std::string",
	"interval":                    "std::string",
	"bytea":                       "std::vector<std::uint8_t>",
	"inet":                        "std::string",
	"vector":                      "std::vector<double>",
}

func CppType(dbType string) string {
	if value, found := cppTypes[dbType]; found {
		return value
	}
	return "std::string"
}
