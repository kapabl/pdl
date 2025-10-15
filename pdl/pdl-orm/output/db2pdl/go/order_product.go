package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type OrderProductRowRecord struct {
	Row       *pdlgo.Row `pdl:"-"`
	CreatedAt time.Time  `pdl:"created_at"`
	Id        int        `pdl:"id"`
	OrderId   int        `pdl:"order_id"`
	ProductId int        `pdl:"product_id"`
	Quantity  int        `pdl:"quantity"`
	UpdatedAt time.Time  `pdl:"updated_at"`
}

type OrderProductRowFactory struct{}

var OrderProductRow = OrderProductRowFactory{}

func (factory OrderProductRowFactory) New() *OrderProductRowRecord {
	result := &OrderProductRowRecord{
		Row: pdlgo.NewRow("order_product", "id"),
	}
	return result
}

func (factory OrderProductRowFactory) WithStore(store pdlgo.DBStore) *OrderProductRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *OrderProductRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *OrderProductRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *OrderProductRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type OrderProductRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func OrderProductRowWhere() OrderProductRowWhereBuilder {
	result := OrderProductRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("order_product", nil)}
	return result
}

func OrderProductRowWhereWithStore(store pdlgo.DBStore) OrderProductRowWhereBuilder {
	result := OrderProductRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("order_product", store)}
	return result
}

func (builder OrderProductRowWhereBuilder) CreatedAt(value time.Time) OrderProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder OrderProductRowWhereBuilder) Id(value int) OrderProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder OrderProductRowWhereBuilder) OrderId(value int) OrderProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("order_id", pdlgo.OpEq, value)
	return builder
}
func (builder OrderProductRowWhereBuilder) ProductId(value int) OrderProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("product_id", pdlgo.OpEq, value)
	return builder
}
func (builder OrderProductRowWhereBuilder) Quantity(value int) OrderProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("quantity", pdlgo.OpEq, value)
	return builder
}
func (builder OrderProductRowWhereBuilder) UpdatedAt(value time.Time) OrderProductRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}

func (builder OrderProductRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder OrderProductRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type OrderProductRowColumnsDefinition struct {
	CreatedAt string
	Id        string
	OrderId   string
	ProductId string
	Quantity  string
	UpdatedAt string
}

type OrderProductRowOrderByDefinition struct {
	CreatedAt string
	Id        string
	OrderId   string
	ProductId string
	Quantity  string
	UpdatedAt string
}

var OrderProductRowColumns = OrderProductRowColumnsDefinition{
	CreatedAt: "created_at",
	Id:        "id",
	OrderId:   "order_id",
	ProductId: "product_id",
	Quantity:  "quantity",
	UpdatedAt: "updated_at",
}

var OrderProductRowOrderBy = OrderProductRowOrderByDefinition{
	CreatedAt: "created_at",
	Id:        "id",
	OrderId:   "order_id",
	ProductId: "product_id",
	Quantity:  "quantity",
	UpdatedAt: "updated_at",
}

func OrderProductRowColumnList() []string {
	result := []string{
		"created_at",
		"id",
		"order_id",
		"product_id",
		"quantity",
		"updated_at",
	}
	return result
}
