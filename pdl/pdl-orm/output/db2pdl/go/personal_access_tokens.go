package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type PersonalAccessTokenRowRecord struct {
	Row           *pdlgo.Row `pdl:"-"`
	Abilities     string     `pdl:"abilities"`
	CreatedAt     time.Time  `pdl:"created_at"`
	Id            int64      `pdl:"id"`
	LastUsedAt    time.Time  `pdl:"last_used_at"`
	Name          string     `pdl:"name"`
	Token         string     `pdl:"token"`
	TokenableId   int64      `pdl:"tokenable_id"`
	TokenableType string     `pdl:"tokenable_type"`
	UpdatedAt     time.Time  `pdl:"updated_at"`
}

type PersonalAccessTokenRowFactory struct{}

var PersonalAccessTokenRow = PersonalAccessTokenRowFactory{}

func (factory PersonalAccessTokenRowFactory) New() *PersonalAccessTokenRowRecord {
	result := &PersonalAccessTokenRowRecord{
		Row: pdlgo.NewRow("personal_access_tokens", "id"),
	}
	return result
}

func (factory PersonalAccessTokenRowFactory) WithStore(store pdlgo.DBStore) *PersonalAccessTokenRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *PersonalAccessTokenRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *PersonalAccessTokenRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *PersonalAccessTokenRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type PersonalAccessTokenRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func PersonalAccessTokenRowWhere() PersonalAccessTokenRowWhereBuilder {
	result := PersonalAccessTokenRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("personal_access_tokens", nil)}
	return result
}

func PersonalAccessTokenRowWhereWithStore(store pdlgo.DBStore) PersonalAccessTokenRowWhereBuilder {
	result := PersonalAccessTokenRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("personal_access_tokens", store)}
	return result
}

func (builder PersonalAccessTokenRowWhereBuilder) Abilities(value string) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("abilities", pdlgo.OpEq, value)
	return builder
}
func (builder PersonalAccessTokenRowWhereBuilder) CreatedAt(value time.Time) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder PersonalAccessTokenRowWhereBuilder) Id(value int64) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder PersonalAccessTokenRowWhereBuilder) LastUsedAt(value time.Time) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("last_used_at", pdlgo.OpEq, value)
	return builder
}
func (builder PersonalAccessTokenRowWhereBuilder) Name(value string) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder PersonalAccessTokenRowWhereBuilder) Token(value string) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("token", pdlgo.OpEq, value)
	return builder
}
func (builder PersonalAccessTokenRowWhereBuilder) TokenableId(value int64) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("tokenable_id", pdlgo.OpEq, value)
	return builder
}
func (builder PersonalAccessTokenRowWhereBuilder) TokenableType(value string) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("tokenable_type", pdlgo.OpEq, value)
	return builder
}
func (builder PersonalAccessTokenRowWhereBuilder) UpdatedAt(value time.Time) PersonalAccessTokenRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}

func (builder PersonalAccessTokenRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder PersonalAccessTokenRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type PersonalAccessTokenRowColumnsDefinition struct {
	Abilities     string
	CreatedAt     string
	Id            string
	LastUsedAt    string
	Name          string
	Token         string
	TokenableId   string
	TokenableType string
	UpdatedAt     string
}

type PersonalAccessTokenRowOrderByDefinition struct {
	Abilities     string
	CreatedAt     string
	Id            string
	LastUsedAt    string
	Name          string
	Token         string
	TokenableId   string
	TokenableType string
	UpdatedAt     string
}

var PersonalAccessTokenRowColumns = PersonalAccessTokenRowColumnsDefinition{
	Abilities:     "abilities",
	CreatedAt:     "created_at",
	Id:            "id",
	LastUsedAt:    "last_used_at",
	Name:          "name",
	Token:         "token",
	TokenableId:   "tokenable_id",
	TokenableType: "tokenable_type",
	UpdatedAt:     "updated_at",
}

var PersonalAccessTokenRowOrderBy = PersonalAccessTokenRowOrderByDefinition{
	Abilities:     "abilities",
	CreatedAt:     "created_at",
	Id:            "id",
	LastUsedAt:    "last_used_at",
	Name:          "name",
	Token:         "token",
	TokenableId:   "tokenable_id",
	TokenableType: "tokenable_type",
	UpdatedAt:     "updated_at",
}

func PersonalAccessTokenRowColumnList() []string {
	result := []string{
		"abilities",
		"created_at",
		"id",
		"last_used_at",
		"name",
		"token",
		"tokenable_id",
		"tokenable_type",
		"updated_at",
	}
	return result
}
