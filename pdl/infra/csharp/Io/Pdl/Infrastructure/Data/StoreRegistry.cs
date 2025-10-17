using System;

namespace Io.Pdl.Infrastructure.Data;

public static class StoreRegistry
{
    private static IDbStore? _defaultStore;

    public static void SetDefaultStore(IDbStore store)
    {
        _defaultStore = store;
    }

    public static IDbStore GetDefaultStore()
    {
        return _defaultStore ?? throw new InvalidOperationException("Io.Pdl.Infrastructure.Data: default store not configured");
    }

    public static IDbStore Resolve(IDbStore? overrideStore)
    {
        return overrideStore ?? GetDefaultStore();
    }
}
