#include "pdl/infrastructure/data/Row.hpp"

namespace pdl::infrastructure::data {

Row::Row(std::string table, std::string primaryKey)
    : table_(std::move(table)), primaryKey_(std::move(primaryKey)), store_(nullptr) {}

const std::string& Row::Table() const noexcept {
    return table_;
}

const std::string& Row::PrimaryKey() const noexcept {
    return primaryKey_;
}

void Row::SetStore(DBStorePtr store) noexcept {
    store_ = std::move(store);
}

DBStorePtr Row::Store() const noexcept {
    return store_;
}

Row CreateRow(std::string table, std::string primaryKey) {
    return Row(std::move(table), std::move(primaryKey));
}

} // namespace pdl::infrastructure::data
