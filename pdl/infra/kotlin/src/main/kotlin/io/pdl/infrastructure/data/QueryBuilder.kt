package io.pdl.infrastructure.data

class QueryBuilder private constructor(
    private val table: String,
    private var store: DBStore?
) {
    private val filters = mutableListOf<Filter>()
    private val projection = mutableListOf<String>()
    private val orderings = mutableListOf<OrderClause>()
    private var limit: Int? = null
    private var offset: Int? = null

    fun withStore(store: DBStore?): QueryBuilder {
        this.store = store
        return this
    }

    fun filter(column: String, operator: Operator, value: Any?): QueryBuilder {
        filters += Filter(column, operator, value)
        return this
    }

    fun project(vararg columns: String): QueryBuilder {
        projection += columns
        return this
    }

    fun offset(value: Int): QueryBuilder {
        offset = value
        return this
    }

    fun limit(value: Int): QueryBuilder {
        limit = value
        return this
    }

    fun range(offset: Int, limit: Int): QueryBuilder {
        this.offset = offset
        this.limit = limit
        return this
    }

    fun asc(column: String): QueryBuilder = orderBy(column, OrderDirection.ASC)

    fun desc(column: String): QueryBuilder = orderBy(column, OrderDirection.DESC)

    fun orderBy(column: String, direction: OrderDirection): QueryBuilder {
        orderings += OrderClause(column, direction)
        return this
    }

    fun load(): List<Map<String, Any?>> {
        val resolved = StoreRegistry.resolve(store)
        return resolved.select(table, filters, projection, orderings, limit, offset)
    }

    fun delete(primaryKey: String) {
        val resolved = StoreRegistry.resolve(store)
        val rows = resolved.select(table, filters, listOf(primaryKey), emptyList(), null, null)
        rows.forEach { resolved.delete(table, primaryKey, it) }
    }

    companion object {
        @JvmStatic
        fun table(name: String, store: DBStore? = null): QueryBuilder = QueryBuilder(name, store)
    }
}
