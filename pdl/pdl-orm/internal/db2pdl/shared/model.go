package shared

type FieldInfo struct {
	FieldName       string `json:"fieldName"`
	CamelCase       string `json:"camelCase"`
	Original        string `json:"original"`
	SnakeCase       string `json:"snakeCase"`
	PascalCase      string `json:"pascalCase"`
	TsType          string `json:"tsType"`
	DbType          string `json:"dbType"`
	PhpType         string `json:"phpType"`
	PdlType         string `json:"pdlType"`
	PhpAttributes   string `json:"phpAttributes"`
	PdlAttributes   string `json:"pdlAttributes"`
	GoType          string `json:"goType"`
	IsPrimaryKey    bool   `json:"isPrimaryKey"`
	IsNullable      bool   `json:"isNullable"`
	IsAutoIncrement bool   `json:"isAutoIncrement"`
}

type TableData struct {
	Name                   string      `json:"name"`
	TableName              string      `json:"tableName"`
	DbName                 string      `json:"dbName"`
	PdlRowClass            string      `json:"pdlRowClass"`
	RowClass               string      `json:"rowClass"`
	GoStruct               string      `json:"goStruct"`
	RowSetClass            string      `json:"rowSetClass"`
	ColumnsDefinitionClass string      `json:"columnsDefinitionClass"`
	WhereClass             string      `json:"whereClass"`
	OrderByClass           string      `json:"orderByClass"`
	FieldListClass         string      `json:"fieldListClass"`
	ColumnsListTraits      string      `json:"columnsListTraits"`
	CsharpRowSetClass      string      `json:"csharpRowSetClass"`
	PhpUseNamespaces       string      `json:"phpUseNamespaces"`
	PhpEntitiesNamespace   string      `json:"phpEntitiesNamespace"`
	PdlEntitiesNamespace   string      `json:"pdlEntitiesNamespace"`
	PdlUseNamespaces       []string    `json:"pdlUseNamespaces"`
	PrimaryKeyPascalCase   string      `json:"primaryKeyPascalCase"`
	PrimaryKeyCamelCase    string      `json:"primaryKeyCamelCase"`
	PrimaryKeyOriginal     string      `json:"primaryKeyOriginal"`
	FieldsInfo             []FieldInfo `json:"fieldsInfo"`
	GoImports              []string    `json:"goImports"`
}
