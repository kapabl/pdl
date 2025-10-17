package io.pdl.infrastructure.data

object StoreRegistry {
    private var defaultStore: DBStore? = null

    fun setDefaultStore(store: DBStore) {
        defaultStore = store
    }

    fun defaultStore(): DBStore = defaultStore
        ?: error("io.pdl.infrastructure.data: default store not configured")

    fun resolve(override: DBStore?): DBStore = override ?: defaultStore()
}
