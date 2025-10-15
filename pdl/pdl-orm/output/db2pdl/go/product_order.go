package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type ProductOrderRowRecord struct {
	Row       *pdlgo.Row `pdl:"-"`
	CreatedAt time.Time  `pdl:"created_at"`
	Id        int64      `pdl:"id"`
	Position  int        `pdl:"position"`
	ProductId int64      `pdl:"product_id"`
	UpdatedAt time.Time  `pdl:"updated_at"`
	UserId    int64      `pdl:"user_id"`
}

type ProductOrderRowFactory struct{}

var ProductOrderRow = ProductOrderRowFactory{}

func (factory ProductOrderRowFactory) New() *ProductOrderRowRecord {
	result := &ProductOrderRowRecord{
		Row: pdlgo.NewRow("product_order", "id"),
	}
	return result
}

func (factory ProductOrderRowFactory) WithStore(store pdlgo.DBStore) *ProductOrderRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *ProductOrderRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *ProductOrderRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *ProductOrderRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type ProductOrderRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func ProductOrderRowWhere() ProductOrderRowWhereBuilder {
	result := ProductOrderRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("product_order", nil)}
	return result
}

func ProductOrderRowWhereWithStore(store pdlgo.DBStore) ProductOrderRowWhereBuilder {
	result := ProductOrderRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("product_order", store)}
	return result
}

func (builder ProductOrderRowWhereBuilder) CreatedAt(value time.Time) ProductOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder ProductOrderRowWhereBuilder) Id(value int64) ProductOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductOrderRowWhereBuilder) Position(value int) ProductOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("position", pdlgo.OpEq, value)
	return builder
}
func (builder ProductOrderRowWhereBuilder) ProductId(value int64) ProductOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("product_id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductOrderRowWhereBuilder) UpdatedAt(value time.Time) ProductOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder ProductOrderRowWhereBuilder) UserId(value int64) ProductOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}

func (builder ProductOrderRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder ProductOrderRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type ProductOrderRowColumnsDefinition struct {
	CreatedAt string
	Id        string
	Position  string
	ProductId string
	UpdatedAt string
	UserId    string
}

type ProductOrderRowOrderByDefinition struct {
	CreatedAt string
	Id        string
	Position  string
	ProductId string
	UpdatedAt string
	UserId    string
}

var ProductOrderRowColumns = ProductOrderRowColumnsDefinition{
	CreatedAt: "created_at",
	Id:        "id",
	Position:  "position",
	ProductId: "product_id",
	UpdatedAt: "updated_at",
	UserId:    "user_id",
}

var ProductOrderRowOrderBy = ProductOrderRowOrderByDefinition{
	CreatedAt: "created_at",
	Id:        "id",
	Position:  "position",
	ProductId: "product_id",
	UpdatedAt: "updated_at",
	UserId:    "user_id",
}

func ProductOrderRowColumnList() []string {
	result := []string{
		"created_at",
		"id",
		"position",
		"product_id",
		"updated_at",
		"user_id",
	}
	return result
}
