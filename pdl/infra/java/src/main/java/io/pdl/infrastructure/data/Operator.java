package io.pdl.infrastructure.data;

public enum Operator {
    EQ("="),
    NEQ("!="),
    IN("IN");

    private final String sql;

    Operator(String sql) {
        this.sql = sql;
    }

    public String sql() {
        return sql;
    }
}
