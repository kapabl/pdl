package io.pdl.infrastructure.data;

public final class OrderClause {
    private final String column;
    private final OrderDirection direction;

    public OrderClause(String column, OrderDirection direction) {
        this.column = column;
        this.direction = direction;
    }

    public String column() {
        return column;
    }

    public OrderDirection direction() {
        return direction;
    }
}
