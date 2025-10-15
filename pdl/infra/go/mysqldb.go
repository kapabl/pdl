package pdlgo

import (
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type MySQLDB struct {
	connection *sql.DB
}

func NewMySQLDB(connection *sql.DB) *MySQLDB {
	result := &MySQLDB{connection: connection}
	return result
}

func (database *MySQLDB) Insert(table string, primaryKey string, values map[string]any) (map[string]any, error) {
	if len(values) == 0 {
		return nil, errors.New("pdlgo: insert requires at least one column")
	}
	columnNames := make([]string, 0, len(values))
	for columnName := range values {
		columnNames = append(columnNames, columnName)
	}
	sort.Strings(columnNames)
	columnFragments := make([]string, 0, len(columnNames))
	placeholderFragments := make([]string, 0, len(columnNames))
	argumentValues := make([]any, 0, len(columnNames))
	for _, columnName := range columnNames {
		columnFragments = append(columnFragments, mysqlQuoteIdentifier(columnName))
		placeholderFragments = append(placeholderFragments, "?")
		argumentValues = append(argumentValues, values[columnName])
	}
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("INSERT INTO ")
	queryBuilder.WriteString(mysqlQuoteIdentifier(table))
	queryBuilder.WriteString(" (")
	queryBuilder.WriteString(strings.Join(columnFragments, ", "))
	queryBuilder.WriteString(") VALUES (")
	queryBuilder.WriteString(strings.Join(placeholderFragments, ", "))
	queryBuilder.WriteString(")")
	insertResult, executionError := database.connection.Exec(queryBuilder.String(), argumentValues...)
	if executionError != nil {
		return nil, executionError
	}
	insertedValues := make(map[string]any)
	if primaryKey != "" {
		if _, hasPrimaryValue := values[primaryKey]; !hasPrimaryValue {
			if lastID, idErr := insertResult.LastInsertId(); idErr == nil && lastID > 0 {
				insertedValues[primaryKey] = lastID
			}
		}
	}
	return insertedValues, nil
}

func (database *MySQLDB) Update(table string, primaryKey string, values map[string]any) error {
	primaryKeyValue, hasPrimaryKey := values[primaryKey]
	if !hasPrimaryKey {
		result := fmt.Errorf("pdlgo: primary key %s missing", primaryKey)
		return result
	}
	columnNames := make([]string, 0, len(values))
	for columnName := range values {
		columnNames = append(columnNames, columnName)
	}
	sort.Strings(columnNames)
	setFragments := make([]string, 0, len(columnNames))
	argumentValues := make([]any, 0, len(columnNames))
	for _, columnName := range columnNames {
		if columnName == primaryKey {
			continue
		}
		setFragments = append(setFragments, fmt.Sprintf("%s = ?", mysqlQuoteIdentifier(columnName)))
		argumentValues = append(argumentValues, values[columnName])
	}
	if len(setFragments) == 0 {
		result := errors.New("pdlgo: update requires at least one column besides the primary key")
		return result
	}
	argumentValues = append(argumentValues, primaryKeyValue)
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("UPDATE ")
	queryBuilder.WriteString(mysqlQuoteIdentifier(table))
	queryBuilder.WriteString(" SET ")
	queryBuilder.WriteString(strings.Join(setFragments, ", "))
	queryBuilder.WriteString(" WHERE ")
	queryBuilder.WriteString(mysqlQuoteIdentifier(primaryKey))
	queryBuilder.WriteString(" = ?")
	_, executionError := database.connection.Exec(queryBuilder.String(), argumentValues...)
	var result error
	result = executionError
	return result
}

func (database *MySQLDB) Delete(table string, primaryKey string, values map[string]any) error {
	primaryKeyValue, hasPrimaryKey := values[primaryKey]
	if !hasPrimaryKey {
		result := fmt.Errorf("pdlgo: primary key %s missing", primaryKey)
		return result
	}
	queryText := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", mysqlQuoteIdentifier(table), mysqlQuoteIdentifier(primaryKey))
	_, executionError := database.connection.Exec(queryText, primaryKeyValue)
	var result error
	result = executionError
	return result
}

func (database *MySQLDB) Select(table string, filters []Filter, projection []string, orderings []OrderClause, limit *int, offset *int) ([]map[string]any, error) {
	result := make([]map[string]any, 0)
	selectColumns := "*"
	if len(projection) > 0 {
		quotedColumns := make([]string, 0, len(projection))
		for columnIndex := range projection {
			quotedColumns = append(quotedColumns, mysqlQuoteIdentifier(projection[columnIndex]))
		}
		selectColumns = strings.Join(quotedColumns, ", ")
	}
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("SELECT ")
	queryBuilder.WriteString(selectColumns)
	queryBuilder.WriteString(" FROM ")
	queryBuilder.WriteString(mysqlQuoteIdentifier(table))
	argumentValues := make([]any, 0)
	if len(filters) > 0 {
		queryBuilder.WriteString(" WHERE ")
		for filterIndex, filter := range filters {
			if filterIndex > 0 {
				queryBuilder.WriteString(" AND ")
			}
			clauseText, clauseArguments, clauseError := buildMySQLFilterClause(filter)
			if clauseError != nil {
				return result, clauseError
			}
			queryBuilder.WriteString(clauseText)
			argumentValues = append(argumentValues, clauseArguments...)
		}
	}
	if len(orderings) > 0 {
		queryBuilder.WriteString(" ORDER BY ")
		orderFragments := make([]string, 0, len(orderings))
		for _, clause := range orderings {
			fragment := fmt.Sprintf("%s %s", mysqlQuoteIdentifier(clause.Column), clause.Direction)
			orderFragments = append(orderFragments, fragment)
		}
		queryBuilder.WriteString(strings.Join(orderFragments, ", "))
	}
	if limit != nil {
		queryBuilder.WriteString(" LIMIT ?")
		argumentValues = append(argumentValues, *limit)
		if offset != nil {
			queryBuilder.WriteString(" OFFSET ?")
			argumentValues = append(argumentValues, *offset)
		}
	} else if offset != nil {
		queryBuilder.WriteString(" LIMIT 18446744073709551615 OFFSET ?")
		argumentValues = append(argumentValues, *offset)
	}
	rows, queryError := database.connection.Query(queryBuilder.String(), argumentValues...)
	if queryError != nil {
		return result, queryError
	}
	defer rows.Close()
	columnNames, columnError := rows.Columns()
	if columnError != nil {
		return result, columnError
	}
	for rows.Next() {
		rawValues := make([]any, len(columnNames))
		scanTargets := make([]any, len(columnNames))
		for columnIndex := range rawValues {
			scanTargets[columnIndex] = &rawValues[columnIndex]
		}
		scanError := rows.Scan(scanTargets...)
		if scanError != nil {
			return result, scanError
		}
		rowMap := make(map[string]any, len(columnNames))
		for columnIndex, columnName := range columnNames {
			value := rawValues[columnIndex]
			byteValue, isByteSlice := value.([]byte)
			if isByteSlice {
				rowMap[columnName] = string(byteValue)
				continue
			}
			rowMap[columnName] = value
		}
		result = append(result, rowMap)
	}
	iterationError := rows.Err()
	if iterationError != nil {
		return result, iterationError
	}
	return result, nil
}

func buildMySQLFilterClause(filter Filter) (string, []any, error) {
	var result string
	switch filter.Op {
	case OpEq:
		result = fmt.Sprintf("%s = ?", mysqlQuoteIdentifier(filter.Field))
		arguments := []any{filter.Value}
		return result, arguments, nil
	case OpNeq:
		result = fmt.Sprintf("%s <> ?", mysqlQuoteIdentifier(filter.Field))
		arguments := []any{filter.Value}
		return result, arguments, nil
	case OpIn:
		argumentSlice, sliceError := toInterfaceSlice(filter.Value)
		if sliceError != nil {
			return result, nil, sliceError
		}
		if len(argumentSlice) == 0 {
			return result, nil, errors.New("pdlgo: IN filter requires at least one value")
		}
		placeholderFragments := make([]string, 0, len(argumentSlice))
		for range argumentSlice {
			placeholderFragments = append(placeholderFragments, "?")
		}
		result = fmt.Sprintf("%s IN (%s)", mysqlQuoteIdentifier(filter.Field), strings.Join(placeholderFragments, ", "))
		return result, argumentSlice, nil
	default:
		return result, nil, fmt.Errorf("pdlgo: unsupported operator %s", filter.Op)
	}
}

func toInterfaceSlice(value any) ([]any, error) {
	switch typedValue := value.(type) {
	case []any:
		result := make([]any, len(typedValue))
		copy(result, typedValue)
		return result, nil
	case []string:
		result := make([]any, len(typedValue))
		for valueIndex, stringValue := range typedValue {
			result[valueIndex] = stringValue
		}
		return result, nil
	case []int:
		result := make([]any, len(typedValue))
		for valueIndex, intValue := range typedValue {
			result[valueIndex] = intValue
		}
		return result, nil
	default:
		return nil, errors.New("pdlgo: unsupported slice type for IN filter")
	}
}

func mysqlQuoteIdentifier(identifier string) string {
	escaped := strings.ReplaceAll(identifier, "`", "``")
	result := "`" + escaped + "`"
	return result
}
