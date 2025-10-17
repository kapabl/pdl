package io.pdl.infrastructure.data

@Target(AnnotationTarget.FIELD)
@Retention(AnnotationRetention.RUNTIME)
annotation class PdlColumn(
    val name: String,
    val primaryKey: Boolean = false,
    val autoIncrement: Boolean = false,
)
