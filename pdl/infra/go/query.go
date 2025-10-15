package pdlgo

// QueryBuilder is used by generated Where() helpers to compose filters.
type QueryBuilder struct {
	table      string
	store      DBStore
	filters    []Filter
	projection []string
	orderings  []OrderClause
	limit      *int
	offset     *int
}

func NewQueryBuilder(table string, store DBStore) QueryBuilder {
	return QueryBuilder{table: table, store: store, filters: make([]Filter, 0)}
}

func (builder QueryBuilder) WithStore(store DBStore) QueryBuilder {
	builder.store = store
	return builder
}

func (builder QueryBuilder) Eq(field string, value any) QueryBuilder {
	return builder.Filter(field, OpEq, value)
}

func (builder QueryBuilder) Neq(field string, value any) QueryBuilder {
	return builder.Filter(field, OpNeq, value)
}

func (builder QueryBuilder) In(field string, values any) QueryBuilder {
	return builder.Filter(field, OpIn, values)
}

func (builder QueryBuilder) Filter(field string, op Operator, value any) QueryBuilder {
	builder.filters = append(builder.filters, Filter{Field: field, Op: op, Value: value})
	return builder
}

func (builder QueryBuilder) Project(fields ...string) QueryBuilder {
	builder.projection = append(builder.projection, fields...)
	return builder
}

func (builder QueryBuilder) OrderBy(column string, direction OrderDirection) QueryBuilder {
	if column == "" {
		return builder
	}
	builder.orderings = append(builder.orderings, OrderClause{Column: column, Direction: direction})
	return builder
}

func (builder QueryBuilder) Asc(column string) QueryBuilder {
	return builder.OrderBy(column, OrderAsc)
}

func (builder QueryBuilder) Desc(column string) QueryBuilder {
	return builder.OrderBy(column, OrderDesc)
}

func (builder QueryBuilder) Limit(value int) QueryBuilder {
	builder.limit = &value
	return builder
}

func (builder QueryBuilder) Offset(value int) QueryBuilder {
	builder.offset = &value
	return builder
}

func (builder QueryBuilder) Range(offset int, limit int) QueryBuilder {
	builder.offset = &offset
	builder.limit = &limit
	return builder
}

func (builder QueryBuilder) Load() ([]map[string]any, error) {
	resolved, err := resolveStore(builder.store)
	if err != nil {
		return nil, err
	}
	return resolved.Select(builder.table, builder.filters, builder.projection, builder.orderings, builder.limit, builder.offset)
}

func (builder QueryBuilder) Delete(primaryKey string) error {
	resolved, err := resolveStore(builder.store)
	if err != nil {
		return err
	}
	rows, err := resolved.Select(builder.table, builder.filters, []string{primaryKey}, nil, nil, nil)
	if err != nil {
		return err
	}
	for _, row := range rows {
		err = resolved.Delete(builder.table, primaryKey, row)
		if err != nil {
			return err
		}
	}
	return nil
}
