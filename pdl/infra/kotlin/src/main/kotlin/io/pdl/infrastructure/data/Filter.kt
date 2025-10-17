package io.pdl.infrastructure.data

data class Filter(
    val column: String,
    val operator: Operator,
    val value: Any?,
)
