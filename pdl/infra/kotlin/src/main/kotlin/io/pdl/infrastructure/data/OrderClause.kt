package io.pdl.infrastructure.data

data class OrderClause(
    val column: String,
    val direction: OrderDirection,
)
