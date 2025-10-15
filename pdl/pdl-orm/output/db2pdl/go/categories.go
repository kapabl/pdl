package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type CategoryRowRecord struct {
	Row       *pdlgo.Row `pdl:"-"`
	CreatedAt time.Time  `pdl:"created_at"`
	Id        int64      `pdl:"id"`
	Name      string     `pdl:"name"`
	Position  int64      `pdl:"position"`
	Slug      string     `pdl:"slug"`
	Status    string     `pdl:"status"`
	StoreId   int        `pdl:"store_id"`
	UpdatedAt time.Time  `pdl:"updated_at"`
}

type CategoryRowFactory struct{}

var CategoryRow = CategoryRowFactory{}

func (factory CategoryRowFactory) New() *CategoryRowRecord {
	result := &CategoryRowRecord{
		Row: pdlgo.NewRow("categories", "id"),
	}
	return result
}

func (factory CategoryRowFactory) WithStore(store pdlgo.DBStore) *CategoryRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *CategoryRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *CategoryRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *CategoryRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type CategoryRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func CategoryRowWhere() CategoryRowWhereBuilder {
	result := CategoryRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("categories", nil)}
	return result
}

func CategoryRowWhereWithStore(store pdlgo.DBStore) CategoryRowWhereBuilder {
	result := CategoryRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("categories", store)}
	return result
}

func (builder CategoryRowWhereBuilder) CreatedAt(value time.Time) CategoryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryRowWhereBuilder) Id(value int64) CategoryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryRowWhereBuilder) Name(value string) CategoryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryRowWhereBuilder) Position(value int64) CategoryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("position", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryRowWhereBuilder) Slug(value string) CategoryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("slug", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryRowWhereBuilder) Status(value string) CategoryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryRowWhereBuilder) StoreId(value int) CategoryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("store_id", pdlgo.OpEq, value)
	return builder
}
func (builder CategoryRowWhereBuilder) UpdatedAt(value time.Time) CategoryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}

func (builder CategoryRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder CategoryRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type CategoryRowColumnsDefinition struct {
	CreatedAt string
	Id        string
	Name      string
	Position  string
	Slug      string
	Status    string
	StoreId   string
	UpdatedAt string
}

type CategoryRowOrderByDefinition struct {
	CreatedAt string
	Id        string
	Name      string
	Position  string
	Slug      string
	Status    string
	StoreId   string
	UpdatedAt string
}

var CategoryRowColumns = CategoryRowColumnsDefinition{
	CreatedAt: "created_at",
	Id:        "id",
	Name:      "name",
	Position:  "position",
	Slug:      "slug",
	Status:    "status",
	StoreId:   "store_id",
	UpdatedAt: "updated_at",
}

var CategoryRowOrderBy = CategoryRowOrderByDefinition{
	CreatedAt: "created_at",
	Id:        "id",
	Name:      "name",
	Position:  "position",
	Slug:      "slug",
	Status:    "status",
	StoreId:   "store_id",
	UpdatedAt: "updated_at",
}

func CategoryRowColumnList() []string {
	result := []string{
		"created_at",
		"id",
		"name",
		"position",
		"slug",
		"status",
		"store_id",
		"updated_at",
	}
	return result
}
