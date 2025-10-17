#pragma once

#include <memory>
#include <vector>

#include "pdl/infrastructure/data/Types.hpp"

namespace pdl::infrastructure::data {

class DBStore {
public:
    virtual ~DBStore() = default;

    virtual RowMap Insert(const std::string& table, const std::string& primaryKey, RowMap values) = 0;

    virtual void Update(const std::string& table, const std::string& primaryKey, const RowMap& values) = 0;

    virtual void Delete(const std::string& table, const std::string& primaryKey, const RowMap& values) = 0;

    virtual std::vector<RowMap> Select(const std::string& table,
                                       const std::vector<Filter>& filters,
                                       const std::vector<std::string>& projection,
                                       const std::vector<OrderClause>& orderings,
                                       const std::optional<std::int32_t>& limit,
                                       const std::optional<std::int32_t>& offset) = 0;
};

using DBStorePtr = std::shared_ptr<DBStore>;

} // namespace pdl::infrastructure::data
