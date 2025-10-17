package io.pdl.infrastructure.data.store

import io.pdl.infrastructure.data.DBStore
import io.pdl.infrastructure.data.Filter
import io.pdl.infrastructure.data.Operator
import io.pdl.infrastructure.data.OrderClause
import io.pdl.infrastructure.data.OrderDirection
import java.util.concurrent.ConcurrentHashMap

class MemoryStore : DBStore {
    private val tables = ConcurrentHashMap<String, MutableList<MutableMap<String, Any?>>>()

    override fun insert(table: String, primaryKey: String, values: MutableMap<String, Any?>): Map<String, Any?> {
        val rows = tables.computeIfAbsent(table) { mutableListOf() }
        if (!values.containsKey(primaryKey)) {
            values[primaryKey] = rows.size + 1
        }
        rows += HashMap(values)
        return values
    }

    override fun update(table: String, primaryKey: String, values: Map<String, Any?>) {
        val rows = tables[table] ?: throw IllegalStateException("MemoryStore: table not found")
        val key = values[primaryKey]
        val row = rows.firstOrNull { it[primaryKey] == key }
            ?: throw IllegalStateException("MemoryStore: row not found")
        row.putAll(values)
    }

    override fun delete(table: String, primaryKey: String, values: Map<String, Any?>) {
        val rows = tables[table] ?: return
        rows.removeIf { it[primaryKey] == values[primaryKey] }
    }

    override fun select(
        table: String,
        filters: List<Filter>,
        projection: List<String>,
        orderings: List<OrderClause>,
        limit: Int?,
        offset: Int?
    ): List<Map<String, Any?>> {
        var rows: List<MutableMap<String, Any?>> = tables[table]
            ?.map { HashMap(it) as MutableMap<String, Any?> }
            ?: emptyList()

        filters.forEach { filter ->
            rows = rows.filter { matches(it, filter) }
        }

        if (orderings.isNotEmpty()) {
            rows = rows.sortedWith { left, right -> compare(left, right, orderings) }
        }

        val dropped = rows.drop((offset ?: 0).coerceAtLeast(0))
        rows = if (limit != null && limit >= 0) dropped.take(limit) else dropped

        if (projection.isNotEmpty()) {
            rows = rows.map { row ->
                projection.associateWith { column -> row[column] }
                    .toMutableMap()
            }
        }

        return rows
    }

    private fun matches(row: Map<String, Any?>, filter: Filter): Boolean {
        val value = row[filter.column]
        return when (filter.operator) {
            Operator.EQ -> value == filter.value
            Operator.NEQ -> value != filter.value
            Operator.IN -> (filter.value as? Iterable<*>)?.any { it == value } ?: false
        }
    }

    private fun compare(left: Map<String, Any?>, right: Map<String, Any?>, orderings: List<OrderClause>): Int {
        for (order in orderings) {
            val comparison = compareValues(left[order.column], right[order.column])
            if (comparison != 0) {
                return if (order.direction == OrderDirection.ASC) comparison else -comparison
            }
        }
        return 0
    }

    private fun compareValues(left: Any?, right: Any?): Int {
        if (left == right) return 0
        if (left == null) return -1
        if (right == null) return 1
        return when {
            left is Comparable<*> && right is Comparable<*> && left::class == right::class ->
                (left as Comparable<Any?>).compareTo(right)
            else -> left.toString().compareTo(right.toString())
        }
    }
}
