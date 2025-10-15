package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
)

type MigrationRowRecord struct {
	Row       *pdlgo.Row `pdl:"-"`
	Batch     int        `pdl:"batch"`
	Id        int        `pdl:"id"`
	Migration string     `pdl:"migration"`
}

type MigrationRowFactory struct{}

var MigrationRow = MigrationRowFactory{}

func (factory MigrationRowFactory) New() *MigrationRowRecord {
	result := &MigrationRowRecord{
		Row: pdlgo.NewRow("migrations", "id"),
	}
	return result
}

func (factory MigrationRowFactory) WithStore(store pdlgo.DBStore) *MigrationRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *MigrationRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *MigrationRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *MigrationRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type MigrationRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func MigrationRowWhere() MigrationRowWhereBuilder {
	result := MigrationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("migrations", nil)}
	return result
}

func MigrationRowWhereWithStore(store pdlgo.DBStore) MigrationRowWhereBuilder {
	result := MigrationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("migrations", store)}
	return result
}

func (builder MigrationRowWhereBuilder) Batch(value int) MigrationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("batch", pdlgo.OpEq, value)
	return builder
}
func (builder MigrationRowWhereBuilder) Id(value int) MigrationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder MigrationRowWhereBuilder) Migration(value string) MigrationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("migration", pdlgo.OpEq, value)
	return builder
}

func (builder MigrationRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder MigrationRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type MigrationRowColumnsDefinition struct {
	Batch     string
	Id        string
	Migration string
}

type MigrationRowOrderByDefinition struct {
	Batch     string
	Id        string
	Migration string
}

var MigrationRowColumns = MigrationRowColumnsDefinition{
	Batch:     "batch",
	Id:        "id",
	Migration: "migration",
}

var MigrationRowOrderBy = MigrationRowOrderByDefinition{
	Batch:     "batch",
	Id:        "id",
	Migration: "migration",
}

func MigrationRowColumnList() []string {
	result := []string{
		"batch",
		"id",
		"migration",
	}
	return result
}
