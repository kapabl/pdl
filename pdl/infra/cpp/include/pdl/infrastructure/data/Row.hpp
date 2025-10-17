#pragma once

#include <memory>
#include <string>

#include "pdl/infrastructure/data/DBStore.hpp"

namespace pdl::infrastructure::data {

class Row {
public:
    Row(std::string table, std::string primaryKey);

    const std::string& Table() const noexcept;

    const std::string& PrimaryKey() const noexcept;

    void SetStore(DBStorePtr store) noexcept;

    DBStorePtr Store() const noexcept;

private:
    std::string table_;
    std::string primaryKey_;
    DBStorePtr store_;
};

Row CreateRow(std::string table, std::string primaryKey);

} // namespace pdl::infrastructure::data
