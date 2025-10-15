package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type TeamUserRowRecord struct {
	Row       *pdlgo.Row `pdl:"-"`
	CreatedAt time.Time  `pdl:"created_at"`
	Id        int64      `pdl:"id"`
	Role      string     `pdl:"role"`
	TeamId    int64      `pdl:"team_id"`
	UpdatedAt time.Time  `pdl:"updated_at"`
	UserId    int64      `pdl:"user_id"`
}

type TeamUserRowFactory struct{}

var TeamUserRow = TeamUserRowFactory{}

func (factory TeamUserRowFactory) New() *TeamUserRowRecord {
	result := &TeamUserRowRecord{
		Row: pdlgo.NewRow("team_user", "id"),
	}
	return result
}

func (factory TeamUserRowFactory) WithStore(store pdlgo.DBStore) *TeamUserRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *TeamUserRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *TeamUserRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *TeamUserRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type TeamUserRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func TeamUserRowWhere() TeamUserRowWhereBuilder {
	result := TeamUserRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("team_user", nil)}
	return result
}

func TeamUserRowWhereWithStore(store pdlgo.DBStore) TeamUserRowWhereBuilder {
	result := TeamUserRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("team_user", store)}
	return result
}

func (builder TeamUserRowWhereBuilder) CreatedAt(value time.Time) TeamUserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder TeamUserRowWhereBuilder) Id(value int64) TeamUserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder TeamUserRowWhereBuilder) Role(value string) TeamUserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("role", pdlgo.OpEq, value)
	return builder
}
func (builder TeamUserRowWhereBuilder) TeamId(value int64) TeamUserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("team_id", pdlgo.OpEq, value)
	return builder
}
func (builder TeamUserRowWhereBuilder) UpdatedAt(value time.Time) TeamUserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder TeamUserRowWhereBuilder) UserId(value int64) TeamUserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}

func (builder TeamUserRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder TeamUserRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type TeamUserRowColumnsDefinition struct {
	CreatedAt string
	Id        string
	Role      string
	TeamId    string
	UpdatedAt string
	UserId    string
}

type TeamUserRowOrderByDefinition struct {
	CreatedAt string
	Id        string
	Role      string
	TeamId    string
	UpdatedAt string
	UserId    string
}

var TeamUserRowColumns = TeamUserRowColumnsDefinition{
	CreatedAt: "created_at",
	Id:        "id",
	Role:      "role",
	TeamId:    "team_id",
	UpdatedAt: "updated_at",
	UserId:    "user_id",
}

var TeamUserRowOrderBy = TeamUserRowOrderByDefinition{
	CreatedAt: "created_at",
	Id:        "id",
	Role:      "role",
	TeamId:    "team_id",
	UpdatedAt: "updated_at",
	UserId:    "user_id",
}

func TeamUserRowColumnList() []string {
	result := []string{
		"created_at",
		"id",
		"role",
		"team_id",
		"updated_at",
		"user_id",
	}
	return result
}
