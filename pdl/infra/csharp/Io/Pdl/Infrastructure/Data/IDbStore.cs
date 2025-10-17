using System.Collections.Generic;

namespace Io.Pdl.Infrastructure.Data;

public interface IDbStore
{
    IDictionary<string, object?> Insert(string table, string primaryKey, IDictionary<string, object?> values);

    void Update(string table, string primaryKey, IDictionary<string, object?> values);

    void Delete(string table, string primaryKey, IDictionary<string, object?> values);

    IReadOnlyList<IDictionary<string, object?>> Select(
        string table,
        IReadOnlyList<Filter> filters,
        IReadOnlyList<string> projection,
        IReadOnlyList<OrderClause> orderings,
        int? limit,
        int? offset);
}
