package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type MenuRowRecord struct {
	Row         *pdlgo.Row `pdl:"-"`
	CategoryIds string     `pdl:"category_ids"`
	CreatedAt   time.Time  `pdl:"created_at"`
	Description string     `pdl:"description"`
	Id          int64      `pdl:"id"`
	Name        string     `pdl:"name"`
	Schedule    string     `pdl:"schedule"`
	Status      string     `pdl:"status"`
	StoreId     int        `pdl:"store_id"`
	UpdatedAt   time.Time  `pdl:"updated_at"`
}

type MenuRowFactory struct{}

var MenuRow = MenuRowFactory{}

func (factory MenuRowFactory) New() *MenuRowRecord {
	result := &MenuRowRecord{
		Row: pdlgo.NewRow("menus", "id"),
	}
	return result
}

func (factory MenuRowFactory) WithStore(store pdlgo.DBStore) *MenuRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *MenuRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *MenuRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *MenuRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type MenuRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func MenuRowWhere() MenuRowWhereBuilder {
	result := MenuRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("menus", nil)}
	return result
}

func MenuRowWhereWithStore(store pdlgo.DBStore) MenuRowWhereBuilder {
	result := MenuRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("menus", store)}
	return result
}

func (builder MenuRowWhereBuilder) CategoryIds(value string) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("category_ids", pdlgo.OpEq, value)
	return builder
}
func (builder MenuRowWhereBuilder) CreatedAt(value time.Time) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder MenuRowWhereBuilder) Description(value string) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("description", pdlgo.OpEq, value)
	return builder
}
func (builder MenuRowWhereBuilder) Id(value int64) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder MenuRowWhereBuilder) Name(value string) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder MenuRowWhereBuilder) Schedule(value string) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("schedule", pdlgo.OpEq, value)
	return builder
}
func (builder MenuRowWhereBuilder) Status(value string) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder MenuRowWhereBuilder) StoreId(value int) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("store_id", pdlgo.OpEq, value)
	return builder
}
func (builder MenuRowWhereBuilder) UpdatedAt(value time.Time) MenuRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}

func (builder MenuRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder MenuRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type MenuRowColumnsDefinition struct {
	CategoryIds string
	CreatedAt   string
	Description string
	Id          string
	Name        string
	Schedule    string
	Status      string
	StoreId     string
	UpdatedAt   string
}

type MenuRowOrderByDefinition struct {
	CategoryIds string
	CreatedAt   string
	Description string
	Id          string
	Name        string
	Schedule    string
	Status      string
	StoreId     string
	UpdatedAt   string
}

var MenuRowColumns = MenuRowColumnsDefinition{
	CategoryIds: "category_ids",
	CreatedAt:   "created_at",
	Description: "description",
	Id:          "id",
	Name:        "name",
	Schedule:    "schedule",
	Status:      "status",
	StoreId:     "store_id",
	UpdatedAt:   "updated_at",
}

var MenuRowOrderBy = MenuRowOrderByDefinition{
	CategoryIds: "category_ids",
	CreatedAt:   "created_at",
	Description: "description",
	Id:          "id",
	Name:        "name",
	Schedule:    "schedule",
	Status:      "status",
	StoreId:     "store_id",
	UpdatedAt:   "updated_at",
}

func MenuRowColumnList() []string {
	result := []string{
		"category_ids",
		"created_at",
		"description",
		"id",
		"name",
		"schedule",
		"status",
		"store_id",
		"updated_at",
	}
	return result
}
