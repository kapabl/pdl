package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

type Driver struct {
	config           shared.DB2PDLConfig
	defaultNamespace string
	phpNamespace     string
}

func NewDriver(config shared.DB2PDLConfig, defaultNamespace string, phpNamespace string) *Driver {
	return &Driver{config: config, defaultNamespace: defaultNamespace, phpNamespace: phpNamespace}
}

func (driver *Driver) Open(ctx context.Context) (*sql.DB, error) {
	connection := driver.config.Connection
	portValue := connection.Port
	if portValue == "" {
		portValue = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", connection.User, connection.Password, connection.Host, portValue, connection.Database)
	database, openErr := sql.Open("mysql", dsn)
	if openErr != nil {
		return nil, openErr
	}
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if pingErr := database.PingContext(pingCtx); pingErr != nil {
		database.Close()
		return nil, pingErr
	}
	return database, nil
}

func (driver *Driver) ListTables(ctx context.Context, connection *sql.DB) ([]string, error) {
	query := "SHOW TABLES FROM `" + driver.config.Connection.Database + "`"
	rows, queryErr := connection.QueryContext(ctx, query)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()
	columns, colErr := rows.Columns()
	if colErr != nil {
		return nil, colErr
	}
	holders := make([]sql.NullString, len(columns))
	receivers := make([]interface{}, len(columns))
	for index := range holders {
		receivers[index] = &holders[index]
	}
	result := make([]string, 0)
	for rows.Next() {
		if scanErr := rows.Scan(receivers...); scanErr != nil {
			return nil, scanErr
		}
		if len(holders) > 0 && holders[0].Valid {
			result = append(result, holders[0].String)
		}
	}
	return result, nil
}

func (driver *Driver) BuildTableData(ctx context.Context, connection *sql.DB, tableName string, baseName string) (shared.TableData, error) {
	fields := make([]shared.FieldInfo, 0)
	var primaryKeyCamel string
	var primaryKeyPascal string
	var primaryKeyOriginal string
	goImports := map[string]struct{}{
		"github.com/kapablanka/pdl/pdl/infra/go": {},
	}
	query := "SHOW COLUMNS FROM `" + tableName + "`"
	rows, queryErr := connection.QueryContext(ctx, query)
	if queryErr != nil {
		return shared.TableData{}, queryErr
	}
	defer rows.Close()
	for rows.Next() {
		var fieldName sql.NullString
		var fieldType sql.NullString
		var fieldNull sql.NullString
		var fieldKey sql.NullString
		var fieldDefault sql.NullString
		var fieldExtra sql.NullString
		scanErr := rows.Scan(&fieldName, &fieldType, &fieldNull, &fieldKey, &fieldDefault, &fieldExtra)
		if scanErr != nil {
			return shared.TableData{}, scanErr
		}
		originalName := fieldName.String
		if shared.ContainsCaseSensitive(driver.config.ExcludedColumns, originalName) {
			continue
		}
		camelCase := shared.ToCamelCase(originalName)
		pascalCase := shared.ToPascalCase(originalName)
		cleanType := shared.ClearSQLType(fieldType.String)
		goType := sqlTypeToLanguage("Go", cleanType)
		if strings.Contains(goType, ".") {
			parts := strings.Split(goType, ".")
			if len(parts) == 2 && parts[0] == "time" {
				goImports["time"] = struct{}{}
			}
		}
		field := shared.FieldInfo{
			FieldName:  camelCase,
			CamelCase:  camelCase,
			Original:   originalName,
			SnakeCase:  originalName,
			PascalCase: pascalCase,
			TsType:     sqlTypeToLanguage("TypeScript", cleanType),
			DbType:     cleanType,
			PhpType:    sqlTypeToLanguage("Php", cleanType),
			PdlType:    sqlTypeToLanguage("Pdl", cleanType),
			GoType:     goType,
		}
		phpAttributes := make([]string, 0)
		pdlAttributes := make([]string, 0)
		isPrimary := fieldKey.String == "PRI"
		isNullable := strings.EqualFold(fieldNull.String, "YES")
		isAutoIncrement := strings.Contains(strings.ToLower(fieldExtra.String), "auto_increment")
		field.IsPrimaryKey = isPrimary
		field.IsNullable = isNullable
		field.IsAutoIncrement = isAutoIncrement
		if isPrimary {
			primaryKeyCamel = camelCase
			primaryKeyPascal = pascalCase
			primaryKeyOriginal = originalName
			pdlAttributes = append(pdlAttributes, shared.GenerateAttribute(field, driver.config.PDL.Attributes.DbID))
			phpAttributes = append(phpAttributes, shared.GenerateAttribute(field, driver.config.PHP.Attributes.DbID))
		}
		if shared.CantConvertBackAndForth(originalName) {
			pdlAttributes = append(pdlAttributes, shared.GenerateAttribute(field, driver.config.PDL.Attributes.ColumnName))
			phpAttributes = append(phpAttributes, shared.GenerateAttribute(field, driver.config.PHP.Attributes.ColumnName))
		}
		if len(phpAttributes) > 0 {
			field.PhpAttributes = strings.Join(phpAttributes, "\n")
		}
		if len(pdlAttributes) > 0 {
			field.PdlAttributes = strings.Join(pdlAttributes, "\n")
		}
		fields = append(fields, field)
	}
	shared.SortFieldInfos(fields)
	return shared.ComposeTableData(baseName, tableName, driver.config, driver.defaultNamespace, driver.phpNamespace, fields, goImports, primaryKeyCamel, primaryKeyPascal, primaryKeyOriginal), nil
}
