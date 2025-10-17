package io.pdl.infrastructure.data;

public abstract class Row {
    private static volatile DBStore defaultStore;

    private final String table;
    private final String primaryKey;
    private DBStore store;

    protected Row(String table, String primaryKey) {
        this.table = table;
        this.primaryKey = primaryKey;
    }

    public static void setDefaultStore(DBStore store) {
        defaultStore = store;
    }

    static DBStore resolveStore(DBStore override) {
        if (override != null) {
            return override;
        }
        DBStore resolved = defaultStore;
        if (resolved == null) {
            throw new IllegalStateException("No DBStore configured for db2pdl rows. Call Row.setDefaultStore first.");
        }
        return resolved;
    }

    DBStore internalStore() {
        return store;
    }

    void setInternalStore(DBStore store) {
        this.store = store;
    }

    public void setStore(DBStore store) {
        this.store = store;
    }

    public DBStore getStore() {
        return store;
    }

    public String getTable() {
        return table;
    }

    public String getPrimaryKey() {
        return primaryKey;
    }
}
