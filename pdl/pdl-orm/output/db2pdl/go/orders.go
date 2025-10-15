package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type OrderRowRecord struct {
	Row               *pdlgo.Row `pdl:"-"`
	CourierUuid       string     `pdl:"courier_uuid"`
	CreatedAt         time.Time  `pdl:"created_at"`
	DeliveryTaxAmount int        `pdl:"delivery_tax_amount"`
	DiscountAmount    int        `pdl:"discount_amount"`
	Id                int        `pdl:"id"`
	Model             string     `pdl:"model"`
	Notes             string     `pdl:"notes"`
	OrderId           string     `pdl:"order_id"`
	SellerTaxAmount   int        `pdl:"seller_tax_amount"`
	SellerTotal       int        `pdl:"seller_total"`
	Status            string     `pdl:"status"`
	StoreId           int        `pdl:"store_id"`
	SubTotal          int        `pdl:"sub_total"`
	TaxAmount         int        `pdl:"tax_amount"`
	Total             int        `pdl:"total"`
	UpdatedAt         time.Time  `pdl:"updated_at"`
	UserId            int        `pdl:"user_id"`
}

type OrderRowFactory struct{}

var OrderRow = OrderRowFactory{}

func (factory OrderRowFactory) New() *OrderRowRecord {
	result := &OrderRowRecord{
		Row: pdlgo.NewRow("orders", "id"),
	}
	return result
}

func (factory OrderRowFactory) WithStore(store pdlgo.DBStore) *OrderRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *OrderRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *OrderRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *OrderRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type OrderRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func OrderRowWhere() OrderRowWhereBuilder {
	result := OrderRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("orders", nil)}
	return result
}

func OrderRowWhereWithStore(store pdlgo.DBStore) OrderRowWhereBuilder {
	result := OrderRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("orders", store)}
	return result
}

func (builder OrderRowWhereBuilder) CourierUuid(value string) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("courier_uuid", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) CreatedAt(value time.Time) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) DeliveryTaxAmount(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("delivery_tax_amount", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) DiscountAmount(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("discount_amount", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) Id(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) Model(value string) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("model", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) Notes(value string) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("notes", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) OrderId(value string) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("order_id", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) SellerTaxAmount(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("seller_tax_amount", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) SellerTotal(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("seller_total", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) Status(value string) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) StoreId(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("store_id", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) SubTotal(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("sub_total", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) TaxAmount(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("tax_amount", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) Total(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("total", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) UpdatedAt(value time.Time) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder OrderRowWhereBuilder) UserId(value int) OrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}

func (builder OrderRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder OrderRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type OrderRowColumnsDefinition struct {
	CourierUuid       string
	CreatedAt         string
	DeliveryTaxAmount string
	DiscountAmount    string
	Id                string
	Model             string
	Notes             string
	OrderId           string
	SellerTaxAmount   string
	SellerTotal       string
	Status            string
	StoreId           string
	SubTotal          string
	TaxAmount         string
	Total             string
	UpdatedAt         string
	UserId            string
}

type OrderRowOrderByDefinition struct {
	CourierUuid       string
	CreatedAt         string
	DeliveryTaxAmount string
	DiscountAmount    string
	Id                string
	Model             string
	Notes             string
	OrderId           string
	SellerTaxAmount   string
	SellerTotal       string
	Status            string
	StoreId           string
	SubTotal          string
	TaxAmount         string
	Total             string
	UpdatedAt         string
	UserId            string
}

var OrderRowColumns = OrderRowColumnsDefinition{
	CourierUuid:       "courier_uuid",
	CreatedAt:         "created_at",
	DeliveryTaxAmount: "delivery_tax_amount",
	DiscountAmount:    "discount_amount",
	Id:                "id",
	Model:             "model",
	Notes:             "notes",
	OrderId:           "order_id",
	SellerTaxAmount:   "seller_tax_amount",
	SellerTotal:       "seller_total",
	Status:            "status",
	StoreId:           "store_id",
	SubTotal:          "sub_total",
	TaxAmount:         "tax_amount",
	Total:             "total",
	UpdatedAt:         "updated_at",
	UserId:            "user_id",
}

var OrderRowOrderBy = OrderRowOrderByDefinition{
	CourierUuid:       "courier_uuid",
	CreatedAt:         "created_at",
	DeliveryTaxAmount: "delivery_tax_amount",
	DiscountAmount:    "discount_amount",
	Id:                "id",
	Model:             "model",
	Notes:             "notes",
	OrderId:           "order_id",
	SellerTaxAmount:   "seller_tax_amount",
	SellerTotal:       "seller_total",
	Status:            "status",
	StoreId:           "store_id",
	SubTotal:          "sub_total",
	TaxAmount:         "tax_amount",
	Total:             "total",
	UpdatedAt:         "updated_at",
	UserId:            "user_id",
}

func OrderRowColumnList() []string {
	result := []string{
		"courier_uuid",
		"created_at",
		"delivery_tax_amount",
		"discount_amount",
		"id",
		"model",
		"notes",
		"order_id",
		"seller_tax_amount",
		"seller_total",
		"status",
		"store_id",
		"sub_total",
		"tax_amount",
		"total",
		"updated_at",
		"user_id",
	}
	return result
}
