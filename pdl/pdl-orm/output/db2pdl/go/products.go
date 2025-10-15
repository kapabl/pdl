package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type ProductRowRecord struct {
	Row            *pdlgo.Row `pdl:"-"`
	CreatedAt      time.Time  `pdl:"created_at"`
	DeliveryPrice  int        `pdl:"delivery_price"`
	DepositOptions string     `pdl:"deposit_options"`
	Description    string     `pdl:"description"`
	Details        string     `pdl:"details"`
	Featured       int        `pdl:"featured"`
	Id             int        `pdl:"id"`
	Images         string     `pdl:"images"`
	Keywords       string     `pdl:"keywords"`
	Name           string     `pdl:"name"`
	Price          int        `pdl:"price"`
	Quantity       int        `pdl:"quantity"`
	Slug           string     `pdl:"slug"`
	Status         string     `pdl:"status"`
	StoreId        int        `pdl:"store_id"`
	UpdatedAt      time.Time  `pdl:"updated_at"`
	UserId         int64      `pdl:"user_id"`
	Uuid           string     `pdl:"uuid"`
}

type ProductRowFactory struct{}

var ProductRow = ProductRowFactory{}

func (factory ProductRowFactory) New() *ProductRowRecord {
	result := &ProductRowRecord{
		Row: pdlgo.NewRow("products", "id"),
	}
	return result
}

func (factory ProductRowFactory) WithStore(store pdlgo.DBStore) *ProductRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *ProductRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *ProductRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *ProductRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type ProductRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func ProductRowWhere() ProductRowWhereBuilder {
	result := ProductRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("products", nil)}
	return result
}

func ProductRowWhereWithStore(store pdlgo.DBStore) ProductRowWhereBuilder {
	result := ProductRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("products", store)}
	return result
}

func (builder ProductRowWhereBuilder) CreatedAt(value time.Time) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) DeliveryPrice(value int) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("delivery_price", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) DepositOptions(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("deposit_options", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Description(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("description", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Details(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("details", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Featured(value int) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("featured", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Id(value int) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Images(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("images", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Keywords(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("keywords", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Name(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Price(value int) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("price", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Quantity(value int) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("quantity", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Slug(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("slug", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Status(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) StoreId(value int) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("store_id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) UpdatedAt(value time.Time) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) UserId(value int64) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductRowWhereBuilder) Uuid(value string) ProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("uuid", pdlgo.OpEq, value)
	return builder
}

func (builder ProductRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder ProductRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type ProductRowColumnsDefinition struct {
	CreatedAt      string
	DeliveryPrice  string
	DepositOptions string
	Description    string
	Details        string
	Featured       string
	Id             string
	Images         string
	Keywords       string
	Name           string
	Price          string
	Quantity       string
	Slug           string
	Status         string
	StoreId        string
	UpdatedAt      string
	UserId         string
	Uuid           string
}

type ProductRowOrderByDefinition struct {
	CreatedAt      string
	DeliveryPrice  string
	DepositOptions string
	Description    string
	Details        string
	Featured       string
	Id             string
	Images         string
	Keywords       string
	Name           string
	Price          string
	Quantity       string
	Slug           string
	Status         string
	StoreId        string
	UpdatedAt      string
	UserId         string
	Uuid           string
}

var ProductRowColumns = ProductRowColumnsDefinition{
	CreatedAt:      "created_at",
	DeliveryPrice:  "delivery_price",
	DepositOptions: "deposit_options",
	Description:    "description",
	Details:        "details",
	Featured:       "featured",
	Id:             "id",
	Images:         "images",
	Keywords:       "keywords",
	Name:           "name",
	Price:          "price",
	Quantity:       "quantity",
	Slug:           "slug",
	Status:         "status",
	StoreId:        "store_id",
	UpdatedAt:      "updated_at",
	UserId:         "user_id",
	Uuid:           "uuid",
}

var ProductRowOrderBy = ProductRowOrderByDefinition{
	CreatedAt:      "created_at",
	DeliveryPrice:  "delivery_price",
	DepositOptions: "deposit_options",
	Description:    "description",
	Details:        "details",
	Featured:       "featured",
	Id:             "id",
	Images:         "images",
	Keywords:       "keywords",
	Name:           "name",
	Price:          "price",
	Quantity:       "quantity",
	Slug:           "slug",
	Status:         "status",
	StoreId:        "store_id",
	UpdatedAt:      "updated_at",
	UserId:         "user_id",
	Uuid:           "uuid",
}

func ProductRowColumnList() []string {
	result := []string{
		"created_at",
		"delivery_price",
		"deposit_options",
		"description",
		"details",
		"featured",
		"id",
		"images",
		"keywords",
		"name",
		"price",
		"quantity",
		"slug",
		"status",
		"store_id",
		"updated_at",
		"user_id",
		"uuid",
	}
	return result
}
