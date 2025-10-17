using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;

namespace Io.Pdl.Infrastructure.Data;

public abstract class Row
{
    private readonly IReadOnlyDictionary<string, string> _propertyToColumn;
    private readonly IReadOnlyDictionary<string, string> _columnToProperty;
    private IDbStore? _store;

    protected Row(string table, string primaryKey, IReadOnlyDictionary<string, string> columnMap, IDbStore? store = null)
    {
        Table = table;
        PrimaryKey = primaryKey;
        _propertyToColumn = columnMap;
        _columnToProperty = columnMap.ToDictionary(pair => pair.Value.ToLowerInvariant(), pair => pair.Key);
        _store = store;
    }

    public string Table { get; }

    public string PrimaryKey { get; }

    public IDbStore? Store => _store;

    public void SetStore(IDbStore? store)
    {
        _store = store;
    }

    public IDictionary<string, object?> CollectValues()
    {
        var values = new Dictionary<string, object?>();
        foreach (var (property, column) in _propertyToColumn)
        {
            var propertyInfo = GetType().GetProperty(property);
            if (propertyInfo == null)
            {
                continue;
            }

            values[column] = propertyInfo.GetValue(this);
        }

        return values;
    }

    public void ApplyValues(IDictionary<string, object?> values)
    {
        foreach (var (column, value) in values)
        {
            var key = column.ToLower(CultureInfo.InvariantCulture);
            if (!_columnToProperty.TryGetValue(key, out var property))
            {
                property = DerivePropertyName(column);
                if (property == null)
                {
                    continue;
                }
            }

            var propertyInfo = GetType().GetProperty(property);
            propertyInfo?.SetValue(this, value);
        }
    }

    public IDbStore ResolveStore()
    {
        return StoreRegistry.Resolve(_store);
    }

    private string? DerivePropertyName(string column)
    {
        var segments = column.Split('_');
        var property = string.Concat(segments.Select(static segment => char.ToUpperInvariant(segment[0]) + segment[1..].ToLowerInvariant()));
        if (property.Length == 0)
        {
            return null;
        }

        property = char.ToLowerInvariant(property[0]) + property[1..];
        return GetType().GetProperty(property) != null ? property : null;
    }
}
