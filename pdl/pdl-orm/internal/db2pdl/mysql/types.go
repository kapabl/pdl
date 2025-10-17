package mysql

var goTypes = map[string]string{
	"int":        "int",
	"tinyint":    "int",
	"smallint":   "int",
	"mediumint":  "int",
	"numeric":    "int",
	"bigint":     "int64",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"bit":        "int",
	"char":       "string",
	"varchar":    "string",
	"tinytext":   "string",
	"text":       "string",
	"mediumtext": "string",
	"longtext":   "string",
	"binary":     "[]byte",
	"varbinary":  "[]byte",
	"tinyblob":   "[]byte",
	"blob":       "[]byte",
	"mediumblob": "[]byte",
	"longblob":   "[]byte",
	"enum":       "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"time":       "string",
	"timestamp":  "time.Time",
	"year":       "int",
}

var timeImports = map[string]struct{}{
	"date":      {},
	"datetime":  {},
	"timestamp": {},
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
