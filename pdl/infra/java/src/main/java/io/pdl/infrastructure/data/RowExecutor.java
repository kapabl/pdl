package io.pdl.infrastructure.data;

import io.pdl.infrastructure.data.RowIntrospector.RowMetadata;

import java.util.Map;

public final class RowExecutor {
    private RowExecutor() {
    }

    public static void create(Row record) throws Exception {
        RowMetadata metadata = RowIntrospector.introspect(record.getClass());
        Map<String, Object> values = metadata.extractValues(record);
        DBStore resolved = Row.resolveStore(record.internalStore());
        record.setInternalStore(resolved);
        Map<String, Object> generated = resolved.insert(record.getTable(), record.getPrimaryKey(), values);
        if (generated != null && !generated.isEmpty()) {
            metadata.apply(record, generated);
        }
    }

    public static void update(Row record) throws Exception {
        RowMetadata metadata = RowIntrospector.introspect(record.getClass());
        Map<String, Object> values = metadata.extractValues(record);
        DBStore resolved = Row.resolveStore(record.internalStore());
        record.setInternalStore(resolved);
        resolved.update(record.getTable(), record.getPrimaryKey(), values);
    }

    public static void delete(Row record) throws Exception {
        RowMetadata metadata = RowIntrospector.introspect(record.getClass());
        Map<String, Object> values = metadata.extractValues(record);
        DBStore resolved = Row.resolveStore(record.internalStore());
        record.setInternalStore(resolved);
        resolved.delete(record.getTable(), record.getPrimaryKey(), values);
    }

    public static void hydrate(Row record, Map<String, Object> values) {
        RowMetadata metadata = RowIntrospector.introspect(record.getClass());
        metadata.apply(record, values);
    }
}
