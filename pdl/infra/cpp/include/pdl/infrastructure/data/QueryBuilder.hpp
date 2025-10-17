#pragma once

#include <optional>
#include <vector>

#include "pdl/infrastructure/data/Types.hpp"
#include "pdl/infrastructure/data/DBStore.hpp"

namespace pdl::infrastructure::data {

class QueryBuilder {
public:
    QueryBuilder(std::string table, DBStorePtr store);

    QueryBuilder& WithStore(DBStorePtr store);

    QueryBuilder& Filter(std::string field, Operator op, Value value);

    QueryBuilder& Project(std::initializer_list<std::string> columns);

    QueryBuilder& Project(const std::vector<std::string>& columns);

    QueryBuilder& Offset(std::int32_t value);

    QueryBuilder& Limit(std::int32_t value);

    QueryBuilder& Range(std::int32_t offset, std::int32_t limit);

    QueryBuilder& OrderBy(std::string column, OrderDirection direction);

    std::vector<RowMap> Load() const;

    void Delete(const std::string& primaryKey) const;

private:
    DBStorePtr ResolveStore() const;

    std::string table_;
    DBStorePtr store_;
    std::vector<Filter> filters_;
    std::vector<std::string> projection_;
    std::vector<OrderClause> orderings_;
    std::optional<std::int32_t> limit_;
    std::optional<std::int32_t> offset_;
};

} // namespace pdl::infrastructure::data
