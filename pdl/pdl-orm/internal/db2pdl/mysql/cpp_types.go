package mysql

var cppTypes = map[string]string{
	"int":        "std::int32_t",
	"tinyint":    "std::int32_t",
	"smallint":   "std::int32_t",
	"mediumint":  "std::int32_t",
	"numeric":    "double",
	"bigint":     "std::int64_t",
	"float":      "float",
	"double":     "double",
	"decimal":    "double",
	"bit":        "std::int32_t",
	"char":       "std::string",
	"varchar":    "std::string",
	"tinytext":   "std::string",
	"text":       "std::string",
	"mediumtext": "std::string",
	"longtext":   "std::string",
	"binary":     "std::vector<std::uint8_t>",
	"varbinary":  "std::vector<std::uint8_t>",
	"tinyblob":   "std::vector<std::uint8_t>",
	"blob":       "std::vector<std::uint8_t>",
	"mediumblob": "std::vector<std::uint8_t>",
	"longblob":   "std::vector<std::uint8_t>",
	"enum":       "std::string",
	"date":       "std::string",
	"datetime":   "std::string",
	"time":       "std::string",
	"timestamp":  "std::string",
	"year":       "std::int32_t",
}

func CppType(dbType string) string {
	if value, found := cppTypes[dbType]; found {
		return value
	}
	return "std::string"
}
