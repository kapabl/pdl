package io.pdl.infrastructure.data;

import io.pdl.infrastructure.data.annotations.PdlColumn;

import java.lang.reflect.Field;
import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

final class RowIntrospector {
    private static final ConcurrentHashMap<Class<?>, RowMetadata> CACHE = new ConcurrentHashMap<>();

    private RowIntrospector() {
    }

    static RowMetadata introspect(Class<?> type) {
        return CACHE.computeIfAbsent(type, RowIntrospector::scan);
    }

    private static RowMetadata scan(Class<?> type) {
        List<ColumnField> columns = new ArrayList<>();
        Map<String, ColumnField> byColumn = new LinkedHashMap<>();
        ColumnField primary = null;
        Class<?> current = type;
        while (current != null && Row.class.isAssignableFrom(current)) {
            Field[] declared = current.getDeclaredFields();
            for (Field field : declared) {
                PdlColumn annotation = field.getAnnotation(PdlColumn.class);
                if (annotation == null) {
                    continue;
                }
                field.setAccessible(true);
                ColumnField columnField = new ColumnField(field, annotation.name(), annotation.primaryKey(), annotation.autoIncrement());
                if (byColumn.containsKey(columnField.column)) {
                    continue;
                }
                byColumn.put(columnField.column, columnField);
                columns.add(columnField);
                if (columnField.primaryKey && primary == null) {
                    primary = columnField;
                }
            }
            current = current.getSuperclass();
        }
        if (primary == null) {
            throw new IllegalStateException("Primary key not defined for " + type.getName());
        }
        return new RowMetadata(Collections.unmodifiableList(columns), byColumn, primary);
    }

    private static final class ColumnField {
        private final Field field;
        private final String column;
        private final boolean primaryKey;
        @SuppressWarnings("unused")
        private final boolean autoIncrement;

        private ColumnField(Field field, String column, boolean primaryKey, boolean autoIncrement) {
            this.field = field;
            this.column = column;
            this.primaryKey = primaryKey;
            this.autoIncrement = autoIncrement;
        }

        Object read(Row record) {
            try {
                return field.get(record);
            } catch (IllegalAccessException ex) {
                throw new RuntimeException("Unable to read field " + field.getName(), ex);
            }
        }

        void write(Row record, Object value) {
            Object converted = ValueConverters.convert(value, field.getType());
            try {
                field.set(record, converted);
            } catch (IllegalAccessException ex) {
                throw new RuntimeException("Unable to write field " + field.getName(), ex);
            }
        }

        String column() {
            return column;
        }
    }

    static final class RowMetadata {
        private final List<ColumnField> columns;
        private final Map<String, ColumnField> byColumn;
        private final ColumnField primary;

        private RowMetadata(List<ColumnField> columns, Map<String, ColumnField> byColumn, ColumnField primary) {
            this.columns = columns;
            this.byColumn = byColumn;
            this.primary = primary;
        }

        Map<String, Object> extractValues(Row record) {
            Map<String, Object> values = new LinkedHashMap<>();
            for (ColumnField column : columns) {
                values.put(column.column(), column.read(record));
            }
            return values;
        }

        void apply(Row record, Map<String, Object> values) {
            if (values == null) {
                return;
            }
            for (Map.Entry<String, Object> entry : values.entrySet()) {
                ColumnField column = byColumn.get(entry.getKey());
                if (column == null) {
                    continue;
                }
                column.write(record, entry.getValue());
            }
        }

        ColumnField primary() {
            return primary;
        }
    }
}
