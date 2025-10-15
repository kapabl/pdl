package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type DiscountInstanceRowRecord struct {
	Row          *pdlgo.Row `pdl:"-"`
	CreatedAt    time.Time  `pdl:"created_at"`
	DiscountId   int        `pdl:"discount_id"`
	Id           int64      `pdl:"id"`
	InstanceInfo string     `pdl:"instance_info"`
	ProductId    int        `pdl:"product_id"`
	StoreId      int        `pdl:"store_id"`
	UpdatedAt    time.Time  `pdl:"updated_at"`
}

type DiscountInstanceRowFactory struct{}

var DiscountInstanceRow = DiscountInstanceRowFactory{}

func (factory DiscountInstanceRowFactory) New() *DiscountInstanceRowRecord {
	result := &DiscountInstanceRowRecord{
		Row: pdlgo.NewRow("discount_instances", "id"),
	}
	return result
}

func (factory DiscountInstanceRowFactory) WithStore(store pdlgo.DBStore) *DiscountInstanceRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *DiscountInstanceRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *DiscountInstanceRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *DiscountInstanceRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type DiscountInstanceRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func DiscountInstanceRowWhere() DiscountInstanceRowWhereBuilder {
	result := DiscountInstanceRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("discount_instances", nil)}
	return result
}

func DiscountInstanceRowWhereWithStore(store pdlgo.DBStore) DiscountInstanceRowWhereBuilder {
	result := DiscountInstanceRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("discount_instances", store)}
	return result
}

func (builder DiscountInstanceRowWhereBuilder) CreatedAt(value time.Time) DiscountInstanceRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountInstanceRowWhereBuilder) DiscountId(value int) DiscountInstanceRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("discount_id", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountInstanceRowWhereBuilder) Id(value int64) DiscountInstanceRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountInstanceRowWhereBuilder) InstanceInfo(value string) DiscountInstanceRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("instance_info", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountInstanceRowWhereBuilder) ProductId(value int) DiscountInstanceRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("product_id", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountInstanceRowWhereBuilder) StoreId(value int) DiscountInstanceRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("store_id", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountInstanceRowWhereBuilder) UpdatedAt(value time.Time) DiscountInstanceRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}

func (builder DiscountInstanceRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder DiscountInstanceRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type DiscountInstanceRowColumnsDefinition struct {
	CreatedAt    string
	DiscountId   string
	Id           string
	InstanceInfo string
	ProductId    string
	StoreId      string
	UpdatedAt    string
}

type DiscountInstanceRowOrderByDefinition struct {
	CreatedAt    string
	DiscountId   string
	Id           string
	InstanceInfo string
	ProductId    string
	StoreId      string
	UpdatedAt    string
}

var DiscountInstanceRowColumns = DiscountInstanceRowColumnsDefinition{
	CreatedAt:    "created_at",
	DiscountId:   "discount_id",
	Id:           "id",
	InstanceInfo: "instance_info",
	ProductId:    "product_id",
	StoreId:      "store_id",
	UpdatedAt:    "updated_at",
}

var DiscountInstanceRowOrderBy = DiscountInstanceRowOrderByDefinition{
	CreatedAt:    "created_at",
	DiscountId:   "discount_id",
	Id:           "id",
	InstanceInfo: "instance_info",
	ProductId:    "product_id",
	StoreId:      "store_id",
	UpdatedAt:    "updated_at",
}

func DiscountInstanceRowColumnList() []string {
	result := []string{
		"created_at",
		"discount_id",
		"id",
		"instance_info",
		"product_id",
		"store_id",
		"updated_at",
	}
	return result
}
