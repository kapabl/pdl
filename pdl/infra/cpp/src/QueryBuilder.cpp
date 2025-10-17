#include "pdl/infrastructure/data/QueryBuilder.hpp"

#include <stdexcept>

#include "pdl/infrastructure/data/RowExecutor.hpp"

namespace pdl::infrastructure::data {

QueryBuilder::QueryBuilder(std::string table, DBStorePtr store)
    : table_(std::move(table)), store_(std::move(store)) {}

QueryBuilder& QueryBuilder::WithStore(DBStorePtr store) {
    store_ = std::move(store);
    return *this;
}

QueryBuilder& QueryBuilder::Filter(std::string field, Operator op, Value value) {
    filters_.push_back(Filter{std::move(field), op, std::move(value)});
    return *this;
}

QueryBuilder& QueryBuilder::Project(std::initializer_list<std::string> columns) {
    return Project(std::vector<std::string>(columns.begin(), columns.end()));
}

QueryBuilder& QueryBuilder::Project(const std::vector<std::string>& columns) {
    projection_.insert(projection_.end(), columns.begin(), columns.end());
    return *this;
}

QueryBuilder& QueryBuilder::Offset(std::int32_t value) {
    offset_ = value;
    return *this;
}

QueryBuilder& QueryBuilder::Limit(std::int32_t value) {
    limit_ = value;
    return *this;
}

QueryBuilder& QueryBuilder::Range(std::int32_t offset, std::int32_t limit) {
    offset_ = offset;
    limit_ = limit;
    return *this;
}

QueryBuilder& QueryBuilder::OrderBy(std::string column, OrderDirection direction) {
    orderings_.push_back(OrderClause{std::move(column), direction});
    return *this;
}

std::vector<RowMap> QueryBuilder::Load() const {
    auto effectiveStore = ResolveStore();
    return effectiveStore->Select(table_, filters_, projection_, orderings_, limit_, offset_);
}

void QueryBuilder::Delete(const std::string& primaryKey) const {
    auto effectiveStore = ResolveStore();
    auto rows = effectiveStore->Select(table_, filters_, {primaryKey}, {}, std::nullopt, std::nullopt);
    for (auto& row : rows) {
        effectiveStore->Delete(table_, primaryKey, row);
    }
}

DBStorePtr QueryBuilder::ResolveStore() const {
    return detail::ResolveStore(store_);
}

} // namespace pdl::infrastructure::data
