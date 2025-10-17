using System.Collections.Generic;
using System.Linq;

namespace Io.Pdl.Infrastructure.Data;

public sealed class QueryBuilder
{
    private readonly List<Filter> _filters = new();
    private readonly List<string> _projection = new();
    private readonly List<OrderClause> _orderings = new();
    private int? _limit;
    private int? _offset;
    private IDbStore? _store;

    public QueryBuilder(string table, IDbStore? store)
    {
        Table = table;
        _store = store;
    }

    public string Table { get; }

    public QueryBuilder WithStore(IDbStore? store)
    {
        _store = store;
        return this;
    }

    public QueryBuilder Filter(string column, Operator op, object? value)
    {
        _filters.Add(new Filter(column, op, value));
        return this;
    }

    public QueryBuilder Project(params string[] columns)
    {
        _projection.AddRange(columns);
        return this;
    }

    public QueryBuilder Offset(int value)
    {
        _offset = value;
        return this;
    }

    public QueryBuilder Limit(int value)
    {
        _limit = value;
        return this;
    }

    public QueryBuilder Range(int offset, int limit)
    {
        _offset = offset;
        _limit = limit;
        return this;
    }

    public QueryBuilder OrderBy(string column, OrderDirection direction)
    {
        _orderings.Add(new OrderClause(column, direction));
        return this;
    }

    public IReadOnlyList<IDictionary<string, object?>> Load()
    {
        var store = StoreRegistry.Resolve(_store);
        return store.Select(Table, _filters, _projection, _orderings, _limit, _offset);
    }

    public void Delete(string primaryKey)
    {
        var store = StoreRegistry.Resolve(_store);
        var rows = store.Select(Table, _filters, new[] { primaryKey }, new List<OrderClause>(), null, null);
        foreach (var row in rows)
        {
            store.Delete(Table, primaryKey, row);
        }
    }
}
