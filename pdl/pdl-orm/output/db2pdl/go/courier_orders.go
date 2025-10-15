package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type CourierOrderRowRecord struct {
	Row         *pdlgo.Row `pdl:"-"`
	CourierUuid string     `pdl:"courier_uuid"`
	CreatedAt   time.Time  `pdl:"created_at"`
	Id          int64      `pdl:"id"`
	OrderId     int        `pdl:"order_id"`
	Status      string     `pdl:"status"`
	UpdatedAt   time.Time  `pdl:"updated_at"`
}

type CourierOrderRowFactory struct{}

var CourierOrderRow = CourierOrderRowFactory{}

func (factory CourierOrderRowFactory) New() *CourierOrderRowRecord {
	result := &CourierOrderRowRecord{
		Row: pdlgo.NewRow("courier_orders", "id"),
	}
	return result
}

func (factory CourierOrderRowFactory) WithStore(store pdlgo.DBStore) *CourierOrderRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *CourierOrderRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *CourierOrderRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *CourierOrderRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type CourierOrderRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func CourierOrderRowWhere() CourierOrderRowWhereBuilder {
	result := CourierOrderRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("courier_orders", nil)}
	return result
}

func CourierOrderRowWhereWithStore(store pdlgo.DBStore) CourierOrderRowWhereBuilder {
	result := CourierOrderRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("courier_orders", store)}
	return result
}

func (builder CourierOrderRowWhereBuilder) CourierUuid(value string) CourierOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("courier_uuid", pdlgo.OpEq, value)
	return builder
}
func (builder CourierOrderRowWhereBuilder) CreatedAt(value time.Time) CourierOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder CourierOrderRowWhereBuilder) Id(value int64) CourierOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder CourierOrderRowWhereBuilder) OrderId(value int) CourierOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("order_id", pdlgo.OpEq, value)
	return builder
}
func (builder CourierOrderRowWhereBuilder) Status(value string) CourierOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder CourierOrderRowWhereBuilder) UpdatedAt(value time.Time) CourierOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}

func (builder CourierOrderRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder CourierOrderRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type CourierOrderRowColumnsDefinition struct {
	CourierUuid string
	CreatedAt   string
	Id          string
	OrderId     string
	Status      string
	UpdatedAt   string
}

type CourierOrderRowOrderByDefinition struct {
	CourierUuid string
	CreatedAt   string
	Id          string
	OrderId     string
	Status      string
	UpdatedAt   string
}

var CourierOrderRowColumns = CourierOrderRowColumnsDefinition{
	CourierUuid: "courier_uuid",
	CreatedAt:   "created_at",
	Id:          "id",
	OrderId:     "order_id",
	Status:      "status",
	UpdatedAt:   "updated_at",
}

var CourierOrderRowOrderBy = CourierOrderRowOrderByDefinition{
	CourierUuid: "courier_uuid",
	CreatedAt:   "created_at",
	Id:          "id",
	OrderId:     "order_id",
	Status:      "status",
	UpdatedAt:   "updated_at",
}

func CourierOrderRowColumnList() []string {
	result := []string{
		"courier_uuid",
		"created_at",
		"id",
		"order_id",
		"status",
		"updated_at",
	}
	return result
}
