package io.pdl.infrastructure.data

import java.lang.reflect.Field
import java.util.Locale

open class Row(
    private val table: String,
    private val primaryKey: String,
    store: DBStore? = null
) {
    private val columnToField: Map<String, Field>
    private val fieldToColumn: Map<Field, String>
    private var store: DBStore? = store

    init {
        val fields = javaClass.declaredFields
            .filterNot { it.isSynthetic }
            .onEach { it.isAccessible = true }

        fieldToColumn = fields.associateWith { field ->
            field.getAnnotation(PdlColumn::class.java)?.name ?: field.name
        }
        columnToField = fieldToColumn.entries.associate { (field, column) ->
            column.lowercase(Locale.getDefault()) to field
        }
    }

    fun table(): String = table

    fun primaryKey(): String = primaryKey

    fun store(): DBStore? = store

    fun setStore(store: DBStore?) {
        this.store = store
    }

    fun collectValues(): MutableMap<String, Any?> {
        val values = mutableMapOf<String, Any?>()
        for ((field, column) in fieldToColumn) {
            values[column] = field.get(this)
        }
        return values
    }

    fun applyValues(values: Map<String, Any?>) {
        for ((column, value) in values) {
            val field = columnToField[column.lowercase(Locale.getDefault())] ?: continue
            field.set(this, value)
        }
    }

    fun resolveStore(): DBStore = StoreRegistry.resolve(store)
}
