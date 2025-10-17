package io.pdl.infrastructure.data;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Map;

public final class QueryBuilder {
    private final String table;
    private DBStore store;
    private final List<Filter> filters;
    private final List<String> projection;
    private final List<OrderClause> orderings;
    private Integer limit;
    private Integer offset;

    public QueryBuilder(String table) {
        this(table, null);
    }

    public QueryBuilder(String table, DBStore store) {
        this.table = table;
        this.store = store;
        this.filters = new ArrayList<>();
        this.projection = new ArrayList<>();
        this.orderings = new ArrayList<>();
    }

    public static QueryBuilder table(String table) {
        return new QueryBuilder(table);
    }

    public QueryBuilder withStore(DBStore store) {
        this.store = store;
        return this;
    }

    public QueryBuilder eq(String field, Object value) {
        return filter(field, Operator.EQ, value);
    }

    public QueryBuilder neq(String field, Object value) {
        return filter(field, Operator.NEQ, value);
    }

    public QueryBuilder in(String field, Object value) {
        return filter(field, Operator.IN, value);
    }

    public QueryBuilder filter(String field, Operator operator, Object value) {
        filters.add(new Filter(field, operator, value));
        return this;
    }

    public QueryBuilder project(String... columns) {
        Collections.addAll(projection, columns);
        return this;
    }

    public QueryBuilder orderBy(String column, OrderDirection direction) {
        if (column == null || column.isEmpty()) {
            return this;
        }
        orderings.add(new OrderClause(column, direction));
        return this;
    }

    public QueryBuilder asc(String column) {
        return orderBy(column, OrderDirection.ASC);
    }

    public QueryBuilder desc(String column) {
        return orderBy(column, OrderDirection.DESC);
    }

    public QueryBuilder limit(int value) {
        this.limit = value;
        return this;
    }

    public QueryBuilder offset(int value) {
        this.offset = value;
        return this;
    }

    public QueryBuilder range(int offsetValue, int limitValue) {
        this.offset = offsetValue;
        this.limit = limitValue;
        return this;
    }

    public List<Map<String, Object>> load() throws Exception {
        DBStore resolved = Row.resolveStore(store);
        return resolved.select(table, filters, projection, orderings, limit, offset);
    }

    public void delete(String primaryKey) throws Exception {
        DBStore resolved = Row.resolveStore(store);
        List<String> primaryProjection = new ArrayList<>(1);
        primaryProjection.add(primaryKey);
        List<Map<String, Object>> rows = resolved.select(table, filters, primaryProjection, orderings, limit, offset);
        for (Map<String, Object> entry : rows) {
            resolved.delete(table, primaryKey, entry);
        }
    }
}
