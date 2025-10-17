package mysql

var rustTypes = map[string]string{
	"int":        "i32",
	"tinyint":    "i32",
	"smallint":   "i32",
	"mediumint":  "i32",
	"numeric":    "f64",
	"bigint":     "i64",
	"float":      "f32",
	"double":     "f64",
	"decimal":    "f64",
	"bit":        "bool",
	"char":       "String",
	"varchar":    "String",
	"tinytext":   "String",
	"text":       "String",
	"mediumtext": "String",
	"longtext":   "String",
	"binary":     "Vec<u8>",
	"varbinary":  "Vec<u8>",
	"tinyblob":   "Vec<u8>",
	"blob":       "Vec<u8>",
	"mediumblob": "Vec<u8>",
	"longblob":   "Vec<u8>",
	"enum":       "String",
	"date":       "String",
	"datetime":   "String",
	"time":       "String",
	"timestamp":  "String",
	"year":       "i32",
}

func RustType(dbType string) string {
	if value, found := rustTypes[dbType]; found {
		return value
	}
	return "String"
}
