#pragma once

#include <any>
#include <cstdint>
#include <optional>
#include <string>
#include <unordered_map>

namespace pdl::infrastructure::data {

using Value = std::any;
using RowMap = std::unordered_map<std::string, Value>;

enum class Operator {
    Eq,
    Neq,
    In,
};

enum class OrderDirection {
    Asc,
    Desc,
};

struct Filter {
    std::string field;
    Operator op;
    Value value;
};

struct OrderClause {
    std::string column;
    OrderDirection direction;
};

} // namespace pdl::infrastructure::data
