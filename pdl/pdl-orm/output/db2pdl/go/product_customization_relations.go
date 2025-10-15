package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type ProductCustomizationRelationRowRecord struct {
	Row                    *pdlgo.Row `pdl:"-"`
	CreatedAt              time.Time  `pdl:"created_at"`
	Id                     int64      `pdl:"id"`
	ParentId               int        `pdl:"parent_id"`
	Position               int        `pdl:"position"`
	ProductCustomizationId int        `pdl:"product_customization_id"`
	ProductId              int        `pdl:"product_id"`
	UpdatedAt              time.Time  `pdl:"updated_at"`
	UserId                 int        `pdl:"user_id"`
}

type ProductCustomizationRelationRowFactory struct{}

var ProductCustomizationRelationRow = ProductCustomizationRelationRowFactory{}

func (factory ProductCustomizationRelationRowFactory) New() *ProductCustomizationRelationRowRecord {
	result := &ProductCustomizationRelationRowRecord{
		Row: pdlgo.NewRow("product_customization_relations", "id"),
	}
	return result
}

func (factory ProductCustomizationRelationRowFactory) WithStore(store pdlgo.DBStore) *ProductCustomizationRelationRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *ProductCustomizationRelationRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *ProductCustomizationRelationRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *ProductCustomizationRelationRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type ProductCustomizationRelationRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func ProductCustomizationRelationRowWhere() ProductCustomizationRelationRowWhereBuilder {
	result := ProductCustomizationRelationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("product_customization_relations", nil)}
	return result
}

func ProductCustomizationRelationRowWhereWithStore(store pdlgo.DBStore) ProductCustomizationRelationRowWhereBuilder {
	result := ProductCustomizationRelationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("product_customization_relations", store)}
	return result
}

func (builder ProductCustomizationRelationRowWhereBuilder) CreatedAt(value time.Time) ProductCustomizationRelationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRelationRowWhereBuilder) Id(value int64) ProductCustomizationRelationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRelationRowWhereBuilder) ParentId(value int) ProductCustomizationRelationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("parent_id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRelationRowWhereBuilder) Position(value int) ProductCustomizationRelationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("position", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRelationRowWhereBuilder) ProductCustomizationId(value int) ProductCustomizationRelationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("product_customization_id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRelationRowWhereBuilder) ProductId(value int) ProductCustomizationRelationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("product_id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRelationRowWhereBuilder) UpdatedAt(value time.Time) ProductCustomizationRelationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRelationRowWhereBuilder) UserId(value int) ProductCustomizationRelationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}

func (builder ProductCustomizationRelationRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder ProductCustomizationRelationRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type ProductCustomizationRelationRowColumnsDefinition struct {
	CreatedAt              string
	Id                     string
	ParentId               string
	Position               string
	ProductCustomizationId string
	ProductId              string
	UpdatedAt              string
	UserId                 string
}

type ProductCustomizationRelationRowOrderByDefinition struct {
	CreatedAt              string
	Id                     string
	ParentId               string
	Position               string
	ProductCustomizationId string
	ProductId              string
	UpdatedAt              string
	UserId                 string
}

var ProductCustomizationRelationRowColumns = ProductCustomizationRelationRowColumnsDefinition{
	CreatedAt:              "created_at",
	Id:                     "id",
	ParentId:               "parent_id",
	Position:               "position",
	ProductCustomizationId: "product_customization_id",
	ProductId:              "product_id",
	UpdatedAt:              "updated_at",
	UserId:                 "user_id",
}

var ProductCustomizationRelationRowOrderBy = ProductCustomizationRelationRowOrderByDefinition{
	CreatedAt:              "created_at",
	Id:                     "id",
	ParentId:               "parent_id",
	Position:               "position",
	ProductCustomizationId: "product_customization_id",
	ProductId:              "product_id",
	UpdatedAt:              "updated_at",
	UserId:                 "user_id",
}

func ProductCustomizationRelationRowColumnList() []string {
	result := []string{
		"created_at",
		"id",
		"parent_id",
		"position",
		"product_customization_id",
		"product_id",
		"updated_at",
		"user_id",
	}
	return result
}
