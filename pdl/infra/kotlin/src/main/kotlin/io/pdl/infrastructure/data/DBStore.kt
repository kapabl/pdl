package io.pdl.infrastructure.data

interface DBStore {
    fun insert(table: String, primaryKey: String, values: MutableMap<String, Any?>): Map<String, Any?>

    fun update(table: String, primaryKey: String, values: Map<String, Any?>)

    fun delete(table: String, primaryKey: String, values: Map<String, Any?>)

    fun select(
        table: String,
        filters: List<Filter>,
        projection: List<String>,
        orderings: List<OrderClause>,
        limit: Int?,
        offset: Int?,
    ): List<Map<String, Any?>>
}
