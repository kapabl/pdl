package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type TeamRowRecord struct {
	Row          *pdlgo.Row `pdl:"-"`
	CreatedAt    time.Time  `pdl:"created_at"`
	Id           int64      `pdl:"id"`
	Name         string     `pdl:"name"`
	PersonalTeam int        `pdl:"personal_team"`
	UpdatedAt    time.Time  `pdl:"updated_at"`
	UserId       int64      `pdl:"user_id"`
}

type TeamRowFactory struct{}

var TeamRow = TeamRowFactory{}

func (factory TeamRowFactory) New() *TeamRowRecord {
	result := &TeamRowRecord{
		Row: pdlgo.NewRow("teams", "id"),
	}
	return result
}

func (factory TeamRowFactory) WithStore(store pdlgo.DBStore) *TeamRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *TeamRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *TeamRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *TeamRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type TeamRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func TeamRowWhere() TeamRowWhereBuilder {
	result := TeamRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("teams", nil)}
	return result
}

func TeamRowWhereWithStore(store pdlgo.DBStore) TeamRowWhereBuilder {
	result := TeamRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("teams", store)}
	return result
}

func (builder TeamRowWhereBuilder) CreatedAt(value time.Time) TeamRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder TeamRowWhereBuilder) Id(value int64) TeamRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder TeamRowWhereBuilder) Name(value string) TeamRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder TeamRowWhereBuilder) PersonalTeam(value int) TeamRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("personal_team", pdlgo.OpEq, value)
	return builder
}
func (builder TeamRowWhereBuilder) UpdatedAt(value time.Time) TeamRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder TeamRowWhereBuilder) UserId(value int64) TeamRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}

func (builder TeamRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder TeamRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type TeamRowColumnsDefinition struct {
	CreatedAt    string
	Id           string
	Name         string
	PersonalTeam string
	UpdatedAt    string
	UserId       string
}

type TeamRowOrderByDefinition struct {
	CreatedAt    string
	Id           string
	Name         string
	PersonalTeam string
	UpdatedAt    string
	UserId       string
}

var TeamRowColumns = TeamRowColumnsDefinition{
	CreatedAt:    "created_at",
	Id:           "id",
	Name:         "name",
	PersonalTeam: "personal_team",
	UpdatedAt:    "updated_at",
	UserId:       "user_id",
}

var TeamRowOrderBy = TeamRowOrderByDefinition{
	CreatedAt:    "created_at",
	Id:           "id",
	Name:         "name",
	PersonalTeam: "personal_team",
	UpdatedAt:    "updated_at",
	UserId:       "user_id",
}

func TeamRowColumnList() []string {
	result := []string{
		"created_at",
		"id",
		"name",
		"personal_team",
		"updated_at",
		"user_id",
	}
	return result
}
