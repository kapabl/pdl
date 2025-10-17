package io.pdl.infrastructure.data;

public final class Filter {
    private final String field;
    private final Operator operator;
    private final Object value;

    public Filter(String field, Operator operator, Object value) {
        this.field = field;
        this.operator = operator;
        this.value = value;
    }

    public String field() {
        return field;
    }

    public Operator operator() {
        return operator;
    }

    public Object value() {
        return value;
    }
}
