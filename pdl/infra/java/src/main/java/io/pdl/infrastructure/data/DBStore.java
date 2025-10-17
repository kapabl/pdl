package io.pdl.infrastructure.data;

import java.util.List;
import java.util.Map;

public interface DBStore {
    Map<String, Object> insert(String table, String primaryKey, Map<String, Object> values) throws Exception;

    void update(String table, String primaryKey, Map<String, Object> values) throws Exception;

    void delete(String table, String primaryKey, Map<String, Object> values) throws Exception;

    List<Map<String, Object>> select(
            String table,
            List<Filter> filters,
            List<String> projection,
            List<OrderClause> orderings,
            Integer limit,
            Integer offset) throws Exception;
}
