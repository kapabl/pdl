using System;
using System.Collections.Generic;
using System.Linq;

namespace Io.Pdl.Infrastructure.Data.Store;

using Io.Pdl.Infrastructure.Data;

public sealed class MemoryStore : IDbStore
{
    private readonly Dictionary<string, List<Dictionary<string, object?>>> _tables = new();

    public IDictionary<string, object?> Insert(string table, string primaryKey, IDictionary<string, object?> values)
    {
        if (!_tables.TryGetValue(table, out var rows))
        {
            rows = new List<Dictionary<string, object?>>();
            _tables[table] = rows;
        }

        var row = new Dictionary<string, object?>(values);
        if (!row.TryGetValue(primaryKey, out var keyValue) || keyValue is null)
        {
            keyValue = rows.Count + 1;
            row[primaryKey] = keyValue;
        }

        rows.Add(row);
        return new Dictionary<string, object?> { [primaryKey] = keyValue };
    }

    public void Update(string table, string primaryKey, IDictionary<string, object?> values)
    {
        if (!_tables.TryGetValue(table, out var rows))
        {
            throw new InvalidOperationException("MemoryStore: table not found");
        }

        foreach (var row in rows)
        {
            if (Equals(row.GetValueOrDefault(primaryKey), values.GetValueOrDefault(primaryKey)))
            {
                foreach (var pair in values)
                {
                    row[pair.Key] = pair.Value;
                }
                return;
            }
        }

        throw new InvalidOperationException("MemoryStore: row not found");
    }

    public void Delete(string table, string primaryKey, IDictionary<string, object?> values)
    {
        if (!_tables.TryGetValue(table, out var rows))
        {
            return;
        }

        rows.RemoveAll(row => Equals(row.GetValueOrDefault(primaryKey), values.GetValueOrDefault(primaryKey)));
    }

    public IReadOnlyList<IDictionary<string, object?>> Select(
        string table,
        IReadOnlyList<Filter> filters,
        IReadOnlyList<string> projection,
        IReadOnlyList<OrderClause> orderings,
        int? limit,
        int? offset)
    {
        var rows = _tables.TryGetValue(table, out var data)
            ? data.Select(row => new Dictionary<string, object?>(row)).ToList()
            : new List<Dictionary<string, object?>>();

        foreach (var filter in filters)
        {
            rows = rows.Where(row => Matches(row, filter)).ToList();
        }

        if (orderings.Count > 0)
        {
            rows.Sort((left, right) => Compare(left, right, orderings));
        }

        if (offset.HasValue)
        {
            rows = rows.Skip(offset.Value).ToList();
        }

        if (limit.HasValue)
        {
            rows = rows.Take(limit.Value).ToList();
        }

        if (projection.Count > 0)
        {
            rows = rows
                .Select(row => projection.ToDictionary(column => column, column => row.GetValueOrDefault(column)))
                .ToList();
        }

        return rows;
    }

    private static bool Matches(IDictionary<string, object?> row, Filter filter)
    {
        var value = row.GetValueOrDefault(filter.Column);
        return filter.Operator switch
        {
            Operator.Eq => Equals(value, filter.Value),
            Operator.Neq => !Equals(value, filter.Value),
            Operator.In when filter.Value is IEnumerable<object?> collection => collection.Cast<object?>().Contains(value),
            _ => false,
        };
    }

    private static int Compare(IDictionary<string, object?> left, IDictionary<string, object?> right, IReadOnlyList<OrderClause> orderings)
    {
        foreach (var ordering in orderings)
        {
            var comparison = Comparer<object?>.Default.Compare(
                left.GetValueOrDefault(ordering.Column),
                right.GetValueOrDefault(ordering.Column));
            if (comparison == 0)
            {
                continue;
            }

            return ordering.Direction == OrderDirection.Desc ? -comparison : comparison;
        }

        return 0;
    }
}
