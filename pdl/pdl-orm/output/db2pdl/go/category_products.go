package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type CategoryProductRowRecord struct {
	Row        *pdlgo.Row `pdl:"-"`
	CategoryId int        `pdl:"category_id"`
	CreatedAt  time.Time  `pdl:"created_at"`
	Id         int64      `pdl:"id"`
	Position   int        `pdl:"position"`
	ProductId  int        `pdl:"product_id"`
	StoreId    int        `pdl:"store_id"`
	UpdatedAt  time.Time  `pdl:"updated_at"`
}

type CategoryProductRowFactory struct{}

var CategoryProductRow = CategoryProductRowFactory{}

func (factory CategoryProductRowFactory) New() *CategoryProductRowRecord {
	result := &CategoryProductRowRecord{
		Row: pdlgo.NewRow("category_products", "id"),
	}
	return result
}

func (factory CategoryProductRowFactory) WithStore(store pdlgo.DBStore) *CategoryProductRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *CategoryProductRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *CategoryProductRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *CategoryProductRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type CategoryProductRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func CategoryProductRowWhere() CategoryProductRowWhereBuilder {
	result := CategoryProductRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("category_products", nil)}
	return result
}

func CategoryProductRowWhereWithStore(store pdlgo.DBStore) CategoryProductRowWhereBuilder {
	result := CategoryProductRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("category_products", store)}
	return result
}

func (builder CategoryProductRowWhereBuilder) CategoryId(value int) CategoryProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("category_id", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryProductRowWhereBuilder) CreatedAt(value time.Time) CategoryProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryProductRowWhereBuilder) Id(value int64) CategoryProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryProductRowWhereBuilder) Position(value int) CategoryProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("position", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryProductRowWhereBuilder) ProductId(value int) CategoryProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("product_id", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryProductRowWhereBuilder) StoreId(value int) CategoryProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("store_id", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryProductRowWhereBuilder) UpdatedAt(value time.Time) CategoryProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}

func (builder CategoryProductRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder CategoryProductRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type CategoryProductRowColumnsDefinition struct {
	CategoryId string
	CreatedAt  string
	Id         string
	Position   string
	ProductId  string
	StoreId    string
	UpdatedAt  string
}

type CategoryProductRowOrderByDefinition struct {
	CategoryId string
	CreatedAt  string
	Id         string
	Position   string
	ProductId  string
	StoreId    string
	UpdatedAt  string
}

var CategoryProductRowColumns = CategoryProductRowColumnsDefinition{
	CategoryId: "category_id",
	CreatedAt:  "created_at",
	Id:         "id",
	Position:   "position",
	ProductId:  "product_id",
	StoreId:    "store_id",
	UpdatedAt:  "updated_at",
}

var CategoryProductRowOrderBy = CategoryProductRowOrderByDefinition{
	CategoryId: "category_id",
	CreatedAt:  "created_at",
	Id:         "id",
	Position:   "position",
	ProductId:  "product_id",
	StoreId:    "store_id",
	UpdatedAt:  "updated_at",
}

func CategoryProductRowColumnList() []string {
	result := []string{
		"category_id",
		"created_at",
		"id",
		"position",
		"product_id",
		"store_id",
		"updated_at",
	}
	return result
}
