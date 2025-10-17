using System.Collections.Generic;

namespace Io.Pdl.Infrastructure.Data;

public static class RowExecutor
{
    public static void Create(Row record)
    {
        var store = record.ResolveStore();
        record.SetStore(store);
        var values = record.CollectValues();
        var inserted = store.Insert(record.Table, record.PrimaryKey, values);
        if (inserted.Count > 0)
        {
            record.ApplyValues(inserted);
        }
    }

    public static void Update(Row record)
    {
        var store = record.ResolveStore();
        record.SetStore(store);
        var values = record.CollectValues();
        store.Update(record.Table, record.PrimaryKey, values);
    }

    public static void Delete(Row record)
    {
        var store = record.ResolveStore();
        record.SetStore(store);
        var values = record.CollectValues();
        store.Delete(record.Table, record.PrimaryKey, values);
    }

    public static void Hydrate(Row record, IDictionary<string, object?> values)
    {
        record.ApplyValues(values);
    }
}
