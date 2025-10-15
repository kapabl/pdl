package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type CustomizationSearchRowRecord struct {
	Row                   *pdlgo.Row `pdl:"-"`
	CreatedAt             time.Time  `pdl:"created_at"`
	CustomizationFullName string     `pdl:"customization_full_name"`
	Id                    int64      `pdl:"id"`
	Name                  string     `pdl:"name"`
	RelationId            int        `pdl:"relation_id"`
	UpdatedAt             time.Time  `pdl:"updated_at"`
	UserId                int        `pdl:"user_id"`
}

type CustomizationSearchRowFactory struct{}

var CustomizationSearchRow = CustomizationSearchRowFactory{}

func (factory CustomizationSearchRowFactory) New() *CustomizationSearchRowRecord {
	result := &CustomizationSearchRowRecord{
		Row: pdlgo.NewRow("customization_search", "id"),
	}
	return result
}

func (factory CustomizationSearchRowFactory) WithStore(store pdlgo.DBStore) *CustomizationSearchRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *CustomizationSearchRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *CustomizationSearchRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *CustomizationSearchRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type CustomizationSearchRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func CustomizationSearchRowWhere() CustomizationSearchRowWhereBuilder {
	result := CustomizationSearchRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("customization_search", nil)}
	return result
}

func CustomizationSearchRowWhereWithStore(store pdlgo.DBStore) CustomizationSearchRowWhereBuilder {
	result := CustomizationSearchRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("customization_search", store)}
	return result
}

func (builder CustomizationSearchRowWhereBuilder) CreatedAt(value time.Time) CustomizationSearchRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder CustomizationSearchRowWhereBuilder) CustomizationFullName(value string) CustomizationSearchRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("customization_full_name", pdlgo.OpEq, value)
	return builder
}
func (builder CustomizationSearchRowWhereBuilder) Id(value int64) CustomizationSearchRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder CustomizationSearchRowWhereBuilder) Name(value string) CustomizationSearchRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder CustomizationSearchRowWhereBuilder) RelationId(value int) CustomizationSearchRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("relation_id", pdlgo.OpEq, value)
	return builder
}
func (builder CustomizationSearchRowWhereBuilder) UpdatedAt(value time.Time) CustomizationSearchRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder CustomizationSearchRowWhereBuilder) UserId(value int) CustomizationSearchRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}

func (builder CustomizationSearchRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder CustomizationSearchRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type CustomizationSearchRowColumnsDefinition struct {
	CreatedAt             string
	CustomizationFullName string
	Id                    string
	Name                  string
	RelationId            string
	UpdatedAt             string
	UserId                string
}

type CustomizationSearchRowOrderByDefinition struct {
	CreatedAt             string
	CustomizationFullName string
	Id                    string
	Name                  string
	RelationId            string
	UpdatedAt             string
	UserId                string
}

var CustomizationSearchRowColumns = CustomizationSearchRowColumnsDefinition{
	CreatedAt:             "created_at",
	CustomizationFullName: "customization_full_name",
	Id:                    "id",
	Name:                  "name",
	RelationId:            "relation_id",
	UpdatedAt:             "updated_at",
	UserId:                "user_id",
}

var CustomizationSearchRowOrderBy = CustomizationSearchRowOrderByDefinition{
	CreatedAt:             "created_at",
	CustomizationFullName: "customization_full_name",
	Id:                    "id",
	Name:                  "name",
	RelationId:            "relation_id",
	UpdatedAt:             "updated_at",
	UserId:                "user_id",
}

func CustomizationSearchRowColumnList() []string {
	result := []string{
		"created_at",
		"customization_full_name",
		"id",
		"name",
		"relation_id",
		"updated_at",
		"user_id",
	}
	return result
}
