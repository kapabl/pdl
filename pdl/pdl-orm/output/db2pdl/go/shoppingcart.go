package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type ShoppingcartRowRecord struct {
	Row        *pdlgo.Row `pdl:"-"`
	Content    string     `pdl:"content"`
	CreatedAt  time.Time  `pdl:"created_at"`
	Identifier string     `pdl:"identifier"`
	Instance   string     `pdl:"instance"`
	UpdatedAt  time.Time  `pdl:"updated_at"`
}

type ShoppingcartRowFactory struct{}

var ShoppingcartRow = ShoppingcartRowFactory{}

func (factory ShoppingcartRowFactory) New() *ShoppingcartRowRecord {
	result := &ShoppingcartRowRecord{
		Row: pdlgo.NewRow("shoppingcart", "instance"),
	}
	return result
}

func (factory ShoppingcartRowFactory) WithStore(store pdlgo.DBStore) *ShoppingcartRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *ShoppingcartRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *ShoppingcartRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *ShoppingcartRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type ShoppingcartRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func ShoppingcartRowWhere() ShoppingcartRowWhereBuilder {
	result := ShoppingcartRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("shoppingcart", nil)}
	return result
}

func ShoppingcartRowWhereWithStore(store pdlgo.DBStore) ShoppingcartRowWhereBuilder {
	result := ShoppingcartRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("shoppingcart", store)}
	return result
}

func (builder ShoppingcartRowWhereBuilder) Content(value string) ShoppingcartRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("content", pdlgo.OpEq, value)
	return builder
}
func (builder ShoppingcartRowWhereBuilder) CreatedAt(value time.Time) ShoppingcartRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder ShoppingcartRowWhereBuilder) Identifier(value string) ShoppingcartRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("identifier", pdlgo.OpEq, value)
	return builder
}
func (builder ShoppingcartRowWhereBuilder) Instance(value string) ShoppingcartRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("instance", pdlgo.OpEq, value)
	return builder
}
func (builder ShoppingcartRowWhereBuilder) UpdatedAt(value time.Time) ShoppingcartRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}

func (builder ShoppingcartRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder ShoppingcartRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("instance")
	return result
}

type ShoppingcartRowColumnsDefinition struct {
	Content    string
	CreatedAt  string
	Identifier string
	Instance   string
	UpdatedAt  string
}

type ShoppingcartRowOrderByDefinition struct {
	Content    string
	CreatedAt  string
	Identifier string
	Instance   string
	UpdatedAt  string
}

var ShoppingcartRowColumns = ShoppingcartRowColumnsDefinition{
	Content:    "content",
	CreatedAt:  "created_at",
	Identifier: "identifier",
	Instance:   "instance",
	UpdatedAt:  "updated_at",
}

var ShoppingcartRowOrderBy = ShoppingcartRowOrderByDefinition{
	Content:    "content",
	CreatedAt:  "created_at",
	Identifier: "identifier",
	Instance:   "instance",
	UpdatedAt:  "updated_at",
}

func ShoppingcartRowColumnList() []string {
	result := []string{
		"content",
		"created_at",
		"identifier",
		"instance",
		"updated_at",
	}
	return result
}
