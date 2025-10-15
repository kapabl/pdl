package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
)

type CacheRowRecord struct {
	Row        *pdlgo.Row `pdl:"-"`
	Expiration int        `pdl:"expiration"`
	Key        string     `pdl:"key"`
	Value      string     `pdl:"value"`
}

type CacheRowFactory struct{}

var CacheRow = CacheRowFactory{}

func (factory CacheRowFactory) New() *CacheRowRecord {
	result := &CacheRowRecord{
		Row: pdlgo.NewRow("cache", "key"),
	}
	return result
}

func (factory CacheRowFactory) WithStore(store pdlgo.DBStore) *CacheRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *CacheRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *CacheRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *CacheRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type CacheRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func CacheRowWhere() CacheRowWhereBuilder {
	result := CacheRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("cache", nil)}
	return result
}

func CacheRowWhereWithStore(store pdlgo.DBStore) CacheRowWhereBuilder {
	result := CacheRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("cache", store)}
	return result
}

func (builder CacheRowWhereBuilder) Expiration(value int) CacheRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("expiration", pdlgo.OpEq, value)
	return builder
}
func (builder CacheRowWhereBuilder) Key(value string) CacheRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("key", pdlgo.OpEq, value)
	return builder
}
func (builder CacheRowWhereBuilder) Value(value string) CacheRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("value", pdlgo.OpEq, value)
	return builder
}

func (builder CacheRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder CacheRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("key")
	return result
}

type CacheRowColumnsDefinition struct {
	Expiration string
	Key        string
	Value      string
}

type CacheRowOrderByDefinition struct {
	Expiration string
	Key        string
	Value      string
}

var CacheRowColumns = CacheRowColumnsDefinition{
	Expiration: "expiration",
	Key:        "key",
	Value:      "value",
}

var CacheRowOrderBy = CacheRowOrderByDefinition{
	Expiration: "expiration",
	Key:        "key",
	Value:      "value",
}

func CacheRowColumnList() []string {
	result := []string{
		"expiration",
		"key",
		"value",
	}
	return result
}
