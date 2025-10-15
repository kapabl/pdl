package pdlgo

import (
	"errors"
	"fmt"
	"reflect"
)

// Store abstracts persistence operations for generated rows.
type DBStore interface {
	Insert(table string, primaryKey string, values map[string]any) (map[string]any, error)
	Update(table string, primaryKey string, values map[string]any) error
	Delete(table string, primaryKey string, values map[string]any) error
	Select(table string, filters []Filter, projection []string, orderings []OrderClause, limit *int, offset *int) ([]map[string]any, error)
}

// Filter represents a simple comparison used in query builders.
type Filter struct {
	Field string
	Op    Operator
	Value any
}

// Operator enumerates supported comparison operators.
type Operator string

const (
	OpEq  Operator = "="
	OpNeq Operator = "!="
	OpIn  Operator = "IN"
)

type OrderDirection string

const (
	OrderAsc  OrderDirection = "ASC"
	OrderDesc OrderDirection = "DESC"
)

type OrderClause struct {
	Column    string
	Direction OrderDirection
}

var defaultStore DBStore

func SetDefaultStore(store DBStore) {
	defaultStore = store
}

func resolveStore(storeOverride DBStore) (DBStore, error) {
	if storeOverride != nil {
		return storeOverride, nil
	}
	return defaultStore, nil
}

// BaseRow is embedded into generated structs to provide CRUD helpers.
type Row struct {
	table      string
	primaryKey string
	store      DBStore
}

func NewRow(table string, primaryKey string) *Row {
	resolved, _ := resolveStore(nil)
	return &Row{
		table:      table,
		primaryKey: primaryKey,
		store:      resolved,
	}
}

func (row *Row) SetStore(store DBStore) {
	row.store = store
}

func (row *Row) Store() DBStore {
	return row.store
}

func (row *Row) Table() string {
	return row.table
}

func (row *Row) PrimaryKey() string {
	return row.primaryKey
}

func Create(record interface{}) error {
	rowMeta, values, err := buildValues(record)
	if err != nil {
		return err
	}
	resolved, err := resolveStore(rowMeta.store)
	if err != nil {
		return err
	}
	rowMeta.store = resolved
	insertedValues, err := resolved.Insert(rowMeta.table, rowMeta.primaryKey, values)
	if err != nil {
		return err
	}
	if len(insertedValues) > 0 {
		for key, value := range insertedValues {
			values[key] = value
		}
		_ = Hydrate(record, values)
	}
	return nil
}

func Update(record interface{}) error {
	rowMeta, values, err := buildValues(record)
	if err != nil {
		return err
	}
	resolved, err := resolveStore(rowMeta.store)
	if err != nil {
		return err
	}
	rowMeta.store = resolved
	return resolved.Update(rowMeta.table, rowMeta.primaryKey, values)
}

func Delete(record interface{}) error {
	rowMeta, values, err := buildValues(record)
	if err != nil {
		return err
	}
	resolved, err := resolveStore(rowMeta.store)
	if err != nil {
		return err
	}
	rowMeta.store = resolved
	return resolved.Delete(rowMeta.table, rowMeta.primaryKey, values)
}

type RowValues map[string]any

func MultiInsertRows(records ...interface{}) error {
	if len(records) == 0 {
		return nil
	}
	firstMeta, _, err := buildValues(records[0])
	if err != nil {
		return err
	}
	resolved, err := resolveStore(firstMeta.store)
	if err != nil {
		return err
	}
	firstMeta.store = resolved
	for _, record := range records {
		meta, values, err := buildValues(record)
		if err != nil {
			return err
		}
		if meta.table != firstMeta.table {
			return errors.New("pdlgo: mixed tables not supported in MultiInsertRows")
		}
		insertedValues, err := resolved.Insert(meta.table, meta.primaryKey, copyMap(map[string]any(values)))
		if err != nil {
			return err
		}
		if len(insertedValues) > 0 {
			for key, value := range insertedValues {
				values[key] = value
			}
			_ = Hydrate(record, values)
		}
	}
	return nil
}

func buildValues(record interface{}) (*Row, RowValues, error) {
	if record == nil {
		return nil, nil, errors.New("pdlgo: nil record")
	}
	value := reflect.ValueOf(record)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return nil, nil, errors.New("pdlgo: record must be a non-nil pointer")
	}
	value = value.Elem()
	if value.Kind() != reflect.Struct {
		return nil, nil, errors.New("pdlgo: record must point to a struct")
	}

	var rowMeta *Row
	values := make(RowValues)
	valueType := value.Type()

	for index := 0; index < value.NumField(); index++ {
		field := value.Field(index)
		fieldType := valueType.Field(index)

		if field.Type() == reflect.TypeOf(&Row{}) {
			if field.IsNil() {
				field.Set(reflect.ValueOf(NewRow("", "")))
			}
			rowMeta = field.Interface().(*Row)
			continue
		}

		column := fieldType.Tag.Get("pdl")
		if column == "" {
			continue
		}

		values[column] = field.Interface()
	}

	if rowMeta == nil {
		return nil, nil, errors.New("pdlgo: record missing embedded *Row field")
	}

	return rowMeta, values, nil
}

func copyMap(source map[string]any) map[string]any {
	clone := make(map[string]any, len(source))
	for key, value := range source {
		clone[key] = value
	}
	return clone
}

func Hydrate(record interface{}, values map[string]any) error {
	if record == nil {
		return errors.New("pdlgo: record must not be nil")
	}
	value := reflect.ValueOf(record)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return errors.New("pdlgo: record must be a non-nil pointer")
	}
	value = value.Elem()
	if value.Kind() != reflect.Struct {
		return errors.New("pdlgo: record must point to a struct")
	}
	typeOfRowPtr := reflect.TypeOf(&Row{})
	structType := value.Type()
	for index := 0; index < value.NumField(); index++ {
		field := value.Field(index)
		fieldType := structType.Field(index)
		if field.Type() == typeOfRowPtr {
			if field.IsNil() {
				field.Set(reflect.ValueOf(NewRow("", "")))
			}
			continue
		}
		column := fieldType.Tag.Get("pdl")
		if column == "" {
			continue
		}
		incoming, found := values[column]
		if !found {
			continue
		}
		if !field.CanSet() {
			continue
		}
		incomingValue := reflect.ValueOf(incoming)
		if !incomingValue.IsValid() {
			field.Set(reflect.Zero(field.Type()))
			continue
		}
		if incomingValue.Type().AssignableTo(field.Type()) {
			field.Set(incomingValue)
			continue
		}
		if incomingValue.Type().ConvertibleTo(field.Type()) {
			field.Set(incomingValue.Convert(field.Type()))
			continue
		}
		if field.Kind() == reflect.String {
			field.SetString(fmt.Sprint(incoming))
			continue
		}
	}
	return nil
}
