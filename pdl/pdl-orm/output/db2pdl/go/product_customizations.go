package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type ProductCustomizationRowRecord struct {
	Row             *pdlgo.Row `pdl:"-"`
	CreatedAt       time.Time  `pdl:"created_at"`
	DefaultValue    string     `pdl:"default_value"`
	Description     string     `pdl:"description"`
	Id              int64      `pdl:"id"`
	IsOption        string     `pdl:"is_option"`
	IsSoldOut       string     `pdl:"is_sold_out"`
	MaxQuantity     int        `pdl:"max_quantity"`
	MinQuantity     int        `pdl:"min_quantity"`
	Name            string     `pdl:"name"`
	NutritionalInfo string     `pdl:"nutritional_info"`
	PickupPrice     int        `pdl:"pickup_price"`
	Price           int        `pdl:"price"`
	ShowInOrder     string     `pdl:"show_in_order"`
	Status          string     `pdl:"status"`
	Type            string     `pdl:"type"`
	UpdatedAt       time.Time  `pdl:"updated_at"`
	UserId          int        `pdl:"user_id"`
	Uuid            string     `pdl:"uuid"`
}

type ProductCustomizationRowFactory struct{}

var ProductCustomizationRow = ProductCustomizationRowFactory{}

func (factory ProductCustomizationRowFactory) New() *ProductCustomizationRowRecord {
	result := &ProductCustomizationRowRecord{
		Row: pdlgo.NewRow("product_customizations", "id"),
	}
	return result
}

func (factory ProductCustomizationRowFactory) WithStore(store pdlgo.DBStore) *ProductCustomizationRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *ProductCustomizationRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *ProductCustomizationRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *ProductCustomizationRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type ProductCustomizationRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func ProductCustomizationRowWhere() ProductCustomizationRowWhereBuilder {
	result := ProductCustomizationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("product_customizations", nil)}
	return result
}

func ProductCustomizationRowWhereWithStore(store pdlgo.DBStore) ProductCustomizationRowWhereBuilder {
	result := ProductCustomizationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("product_customizations", store)}
	return result
}

func (builder ProductCustomizationRowWhereBuilder) CreatedAt(value time.Time) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) DefaultValue(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("default_value", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) Description(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("description", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) Id(value int64) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) IsOption(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_option", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) IsSoldOut(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_sold_out", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) MaxQuantity(value int) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("max_quantity", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) MinQuantity(value int) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("min_quantity", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) Name(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) NutritionalInfo(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("nutritional_info", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) PickupPrice(value int) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("pickup_price", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) Price(value int) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("price", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) ShowInOrder(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("show_in_order", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) Status(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) Type(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("type", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) UpdatedAt(value time.Time) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) UserId(value int) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}
func (builder ProductCustomizationRowWhereBuilder) Uuid(value string) ProductCustomizationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("uuid", pdlgo.OpEq, value)
	return builder
}

func (builder ProductCustomizationRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder ProductCustomizationRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type ProductCustomizationRowColumnsDefinition struct {
	CreatedAt       string
	DefaultValue    string
	Description     string
	Id              string
	IsOption        string
	IsSoldOut       string
	MaxQuantity     string
	MinQuantity     string
	Name            string
	NutritionalInfo string
	PickupPrice     string
	Price           string
	ShowInOrder     string
	Status          string
	Type            string
	UpdatedAt       string
	UserId          string
	Uuid            string
}

type ProductCustomizationRowOrderByDefinition struct {
	CreatedAt       string
	DefaultValue    string
	Description     string
	Id              string
	IsOption        string
	IsSoldOut       string
	MaxQuantity     string
	MinQuantity     string
	Name            string
	NutritionalInfo string
	PickupPrice     string
	Price           string
	ShowInOrder     string
	Status          string
	Type            string
	UpdatedAt       string
	UserId          string
	Uuid            string
}

var ProductCustomizationRowColumns = ProductCustomizationRowColumnsDefinition{
	CreatedAt:       "created_at",
	DefaultValue:    "default_value",
	Description:     "description",
	Id:              "id",
	IsOption:        "is_option",
	IsSoldOut:       "is_sold_out",
	MaxQuantity:     "max_quantity",
	MinQuantity:     "min_quantity",
	Name:            "name",
	NutritionalInfo: "nutritional_info",
	PickupPrice:     "pickup_price",
	Price:           "price",
	ShowInOrder:     "show_in_order",
	Status:          "status",
	Type:            "type",
	UpdatedAt:       "updated_at",
	UserId:          "user_id",
	Uuid:            "uuid",
}

var ProductCustomizationRowOrderBy = ProductCustomizationRowOrderByDefinition{
	CreatedAt:       "created_at",
	DefaultValue:    "default_value",
	Description:     "description",
	Id:              "id",
	IsOption:        "is_option",
	IsSoldOut:       "is_sold_out",
	MaxQuantity:     "max_quantity",
	MinQuantity:     "min_quantity",
	Name:            "name",
	NutritionalInfo: "nutritional_info",
	PickupPrice:     "pickup_price",
	Price:           "price",
	ShowInOrder:     "show_in_order",
	Status:          "status",
	Type:            "type",
	UpdatedAt:       "updated_at",
	UserId:          "user_id",
	Uuid:            "uuid",
}

func ProductCustomizationRowColumnList() []string {
	result := []string{
		"created_at",
		"default_value",
		"description",
		"id",
		"is_option",
		"is_sold_out",
		"max_quantity",
		"min_quantity",
		"name",
		"nutritional_info",
		"pickup_price",
		"price",
		"show_in_order",
		"status",
		"type",
		"updated_at",
		"user_id",
		"uuid",
	}
	return result
}
