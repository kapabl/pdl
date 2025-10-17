package io.pdl.infrastructure.data

object RowExecutor {
    @JvmStatic
    fun create(record: Row) {
        val store = record.resolveStore()
        record.setStore(store)
        val values = record.collectValues()
        val inserted = store.insert(record.table(), record.primaryKey(), values)
        if (inserted.isNotEmpty()) {
            record.applyValues(inserted)
        }
    }

    @JvmStatic
    fun update(record: Row) {
        val store = record.resolveStore()
        record.setStore(store)
        val values = record.collectValues()
        store.update(record.table(), record.primaryKey(), values)
    }

    @JvmStatic
    fun delete(record: Row) {
        val store = record.resolveStore()
        record.setStore(store)
        val values = record.collectValues()
        store.delete(record.table(), record.primaryKey(), values)
    }

    @JvmStatic
    fun hydrate(record: Row, values: Map<String, Any?>) {
        record.applyValues(values)
    }
}
