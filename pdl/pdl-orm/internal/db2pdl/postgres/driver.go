package postgres

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
		portValue = "5432"
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", connection.Host, portValue, connection.User, connection.Password, connection.Database)
	database, openErr := sql.Open("postgres", dsn)
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
	query := "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = current_schema() ORDER BY tablename"
	rows, queryErr := connection.QueryContext(ctx, query)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()
	result := make([]string, 0)
	for rows.Next() {
		var tableName sql.NullString
		if scanErr := rows.Scan(&tableName); scanErr != nil {
			return nil, scanErr
		}
		if tableName.Valid {
			result = append(result, tableName.String)
		}
	}
	return result, nil
}

func (driver *Driver) BuildTableData(ctx context.Context, connection *sql.DB, tableName string, baseName string) (shared.TableData, error) {
	primaryKeySet, keyErr := driver.fetchPrimaryKeys(ctx, connection, tableName)
	if keyErr != nil {
		return shared.TableData{}, keyErr
	}
	fields := make([]shared.FieldInfo, 0)
	var primaryKeyCamel string
	var primaryKeyPascal string
	var primaryKeyOriginal string
	goImports := map[string]struct{}{
		"github.com/kapablanka/pdl/pdl/infra/go": {},
	}
	query := "SELECT column_name, data_type, is_nullable, column_default, is_identity FROM information_schema.columns WHERE table_schema = current_schema() AND table_name = $1 ORDER BY ordinal_position"
	rows, queryErr := connection.QueryContext(ctx, query, tableName)
	if queryErr != nil {
		return shared.TableData{}, queryErr
	}
	defer rows.Close()
	for rows.Next() {
		var columnName sql.NullString
		var dataType sql.NullString
		var isNullable sql.NullString
		var columnDefault sql.NullString
		var identityFlag sql.NullString
		if scanErr := rows.Scan(&columnName, &dataType, &isNullable, &columnDefault, &identityFlag); scanErr != nil {
			return shared.TableData{}, scanErr
		}
		originalName := columnName.String
		if shared.ContainsCaseSensitive(driver.config.ExcludedColumns, originalName) {
			continue
		}
		camelCase := shared.ToCamelCase(originalName)
		pascalCase := shared.ToPascalCase(originalName)
		cleanType := shared.ClearSQLType(dataType.String)
		if NeedsTimeImport(cleanType) {
			goImports["time"] = struct{}{}
		}
		field := shared.FieldInfo{
			FieldName:  camelCase,
			CamelCase:  camelCase,
			Original:   originalName,
			SnakeCase:  originalName,
			PascalCase: pascalCase,
			DbType:     cleanType,
		}
		_, isPrimary := primaryKeySet[strings.ToLower(originalName)]
		isNullableColumn := strings.EqualFold(isNullable.String, "YES")
		defaultText := strings.ToLower(columnDefault.String)
		isIdentity := strings.EqualFold(identityFlag.String, "YES")
		isAutoIncrement := strings.Contains(defaultText, "nextval(") || isIdentity
		field.IsPrimaryKey = isPrimary
		field.IsNullable = isNullableColumn
		field.IsAutoIncrement = isAutoIncrement
		if isPrimary {
			if primaryKeyOriginal == "" {
				primaryKeyCamel = camelCase
				primaryKeyPascal = pascalCase
				primaryKeyOriginal = originalName
			}
		}
		fields = append(fields, field)
	}
	shared.SortFieldInfos(fields)
	return shared.ComposeTableData(baseName, tableName, driver.config, driver.defaultNamespace, driver.phpNamespace, fields, goImports, primaryKeyCamel, primaryKeyPascal, primaryKeyOriginal), nil
}

func (driver *Driver) fetchPrimaryKeys(ctx context.Context, connection *sql.DB, tableName string) (map[string]struct{}, error) {
	result := make(map[string]struct{})
	query := "SELECT kc.column_name FROM information_schema.table_constraints tc JOIN information_schema.key_column_usage kc ON kc.constraint_name = tc.constraint_name AND kc.constraint_schema = tc.constraint_schema AND kc.table_name = tc.table_name WHERE tc.constraint_type = 'PRIMARY KEY' AND kc.table_schema = current_schema() AND kc.table_name = $1 ORDER BY kc.ordinal_position"
	rows, queryErr := connection.QueryContext(ctx, query, tableName)
	if queryErr != nil {
		return result, queryErr
	}
	defer rows.Close()
	for rows.Next() {
		var columnName sql.NullString
		if scanErr := rows.Scan(&columnName); scanErr != nil {
			return result, scanErr
		}
		if columnName.Valid {
			result[strings.ToLower(columnName.String)] = struct{}{}
		}
	}
	return result, nil
}
