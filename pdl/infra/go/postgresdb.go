package pdlgo

import (
    "database/sql"
    "errors"
    "fmt"
    "sort"
    "strings"
)

type PostgresDB struct {
    connection *sql.DB
}

func NewPostgresDB(connection *sql.DB) *PostgresDB {
    result := &PostgresDB{connection: connection}
    return result
}

func (database *PostgresDB) Insert(table string, primaryKey string, values map[string]any) (map[string]any, error) {
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
    for index, columnName := range columnNames {
        columnFragments = append(columnFragments, postgresQuoteIdentifier(columnName))
        placeholderFragments = append(placeholderFragments, postgresPlaceholder(index+1))
        argumentValues = append(argumentValues, values[columnName])
    }

    queryBuilder := strings.Builder{}
    queryBuilder.WriteString("INSERT INTO ")
    queryBuilder.WriteString(postgresQuoteIdentifier(table))
    queryBuilder.WriteString(" (")
    queryBuilder.WriteString(strings.Join(columnFragments, ", "))
    queryBuilder.WriteString(") VALUES (")
    queryBuilder.WriteString(strings.Join(placeholderFragments, ", "))
    queryBuilder.WriteString(")")

    if primaryKey != "" {
        if _, provided := values[primaryKey]; !provided {
            queryBuilder.WriteString(" RETURNING ")
            queryBuilder.WriteString(postgresQuoteIdentifier(primaryKey))
            row := database.connection.QueryRow(queryBuilder.String(), argumentValues...)
            var generated any
            if scanErr := row.Scan(&generated); scanErr != nil {
                return nil, scanErr
            }
            return map[string]any{primaryKey: generated}, nil
        }
    }

    _, executionError := database.connection.Exec(queryBuilder.String(), argumentValues...)
    if executionError != nil {
        return nil, executionError
    }
    return nil, nil
}

func (database *PostgresDB) Update(table string, primaryKey string, values map[string]any) error {
    primaryKeyValue, hasPrimaryKey := values[primaryKey]
    if !hasPrimaryKey {
        return fmt.Errorf("pdlgo: primary key %s missing", primaryKey)
    }
    columnNames := make([]string, 0, len(values))
    for columnName := range values {
        columnNames = append(columnNames, columnName)
    }
    sort.Strings(columnNames)
    setFragments := make([]string, 0, len(columnNames))
    argumentValues := make([]any, 0, len(columnNames))
    placeholderIndex := 1
    for _, columnName := range columnNames {
        if columnName == primaryKey {
            continue
        }
        setFragments = append(setFragments, fmt.Sprintf("%s = %s", postgresQuoteIdentifier(columnName), postgresPlaceholder(placeholderIndex)))
        argumentValues = append(argumentValues, values[columnName])
        placeholderIndex++
    }
    if len(setFragments) == 0 {
        return errors.New("pdlgo: update requires at least one column besides the primary key")
    }
    argumentValues = append(argumentValues, primaryKeyValue)
    queryBuilder := strings.Builder{}
    queryBuilder.WriteString("UPDATE ")
    queryBuilder.WriteString(postgresQuoteIdentifier(table))
    queryBuilder.WriteString(" SET ")
    queryBuilder.WriteString(strings.Join(setFragments, ", "))
    queryBuilder.WriteString(" WHERE ")
    queryBuilder.WriteString(postgresQuoteIdentifier(primaryKey))
    queryBuilder.WriteString(" = ")
    queryBuilder.WriteString(postgresPlaceholder(placeholderIndex))
    _, executionError := database.connection.Exec(queryBuilder.String(), argumentValues...)
    return executionError
}

func (database *PostgresDB) Delete(table string, primaryKey string, values map[string]any) error {
    primaryKeyValue, hasPrimaryKey := values[primaryKey]
    if !hasPrimaryKey {
        return fmt.Errorf("pdlgo: primary key %s missing", primaryKey)
    }
    queryText := fmt.Sprintf("DELETE FROM %s WHERE %s = %s", postgresQuoteIdentifier(table), postgresQuoteIdentifier(primaryKey), postgresPlaceholder(1))
    _, executionError := database.connection.Exec(queryText, primaryKeyValue)
    return executionError
}

func (database *PostgresDB) Select(table string, filters []Filter, projection []string, orderings []OrderClause, limit *int, offset *int) ([]map[string]any, error) {
    result := make([]map[string]any, 0)
    selectColumns := "*"
    if len(projection) > 0 {
        quotedColumns := make([]string, 0, len(projection))
        for _, column := range projection {
            quotedColumns = append(quotedColumns, postgresQuoteIdentifier(column))
        }
        selectColumns = strings.Join(quotedColumns, ", ")
    }

    queryBuilder := strings.Builder{}
    queryBuilder.WriteString("SELECT ")
    queryBuilder.WriteString(selectColumns)
    queryBuilder.WriteString(" FROM ")
    queryBuilder.WriteString(postgresQuoteIdentifier(table))

    argumentValues := make([]any, 0)
    placeholderIndex := 1
    if len(filters) > 0 {
        queryBuilder.WriteString(" WHERE ")
        clauses := make([]string, 0, len(filters))
        for _, filter := range filters {
            clause, clauseArgs, nextIndex, clauseErr := buildPostgresFilterClause(filter, placeholderIndex)
            if clauseErr != nil {
                return result, clauseErr
            }
            clauses = append(clauses, clause)
            argumentValues = append(argumentValues, clauseArgs...)
            placeholderIndex = nextIndex
        }
        queryBuilder.WriteString(strings.Join(clauses, " AND "))
    }

    if len(orderings) > 0 {
        fragments := make([]string, 0, len(orderings))
        for _, clause := range orderings {
            fragments = append(fragments, fmt.Sprintf("%s %s", postgresQuoteIdentifier(clause.Column), clause.Direction))
        }
        queryBuilder.WriteString(" ORDER BY ")
        queryBuilder.WriteString(strings.Join(fragments, ", "))
    }

    if limit != nil {
        queryBuilder.WriteString(" LIMIT ")
        queryBuilder.WriteString(postgresPlaceholder(placeholderIndex))
        argumentValues = append(argumentValues, *limit)
        placeholderIndex++
        if offset != nil {
            queryBuilder.WriteString(" OFFSET ")
            queryBuilder.WriteString(postgresPlaceholder(placeholderIndex))
            argumentValues = append(argumentValues, *offset)
            placeholderIndex++
        }
    } else if offset != nil {
        queryBuilder.WriteString(" OFFSET ")
        queryBuilder.WriteString(postgresPlaceholder(placeholderIndex))
        argumentValues = append(argumentValues, *offset)
        placeholderIndex++
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
        if scanError := rows.Scan(scanTargets...); scanError != nil {
            return result, scanError
        }
        rowMap := make(map[string]any, len(columnNames))
        for columnIndex, columnName := range columnNames {
            value := rawValues[columnIndex]
            if byteValue, ok := value.([]byte); ok {
                rowMap[columnName] = string(byteValue)
                continue
            }
            rowMap[columnName] = value
        }
        result = append(result, rowMap)
    }
    if iterationError := rows.Err(); iterationError != nil {
        return result, iterationError
    }
    return result, nil
}

func buildPostgresFilterClause(filter Filter, startIndex int) (string, []any, int, error) {
    switch filter.Op {
    case OpEq:
        clause := fmt.Sprintf("%s = %s", postgresQuoteIdentifier(filter.Field), postgresPlaceholder(startIndex))
        return clause, []any{filter.Value}, startIndex + 1, nil
    case OpNeq:
        clause := fmt.Sprintf("%s <> %s", postgresQuoteIdentifier(filter.Field), postgresPlaceholder(startIndex))
        return clause, []any{filter.Value}, startIndex + 1, nil
    case OpIn:
        slice, err := toInterfaceSlice(filter.Value)
        if err != nil {
            return "", nil, startIndex, err
        }
        if len(slice) == 0 {
            return "", nil, startIndex, errors.New("pdlgo: IN filter requires at least one value")
        }
        placeholders := make([]string, 0, len(slice))
        for i := range slice {
            placeholders = append(placeholders, postgresPlaceholder(startIndex+i))
        }
        clause := fmt.Sprintf("%s IN (%s)", postgresQuoteIdentifier(filter.Field), strings.Join(placeholders, ", "))
        return clause, slice, startIndex + len(slice), nil
    default:
        return "", nil, startIndex, fmt.Errorf("pdlgo: unsupported operator %s", filter.Op)
    }
}

func postgresQuoteIdentifier(identifier string) string {
    escaped := strings.ReplaceAll(identifier, "\"", "\"\"")
    return "\"" + escaped + "\""
}

func postgresPlaceholder(index int) string {
    return fmt.Sprintf("$%d", index)
}
