package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
)

type SessionRowRecord struct {
	Row          *pdlgo.Row `pdl:"-"`
	Id           string     `pdl:"id"`
	IpAddress    string     `pdl:"ip_address"`
	LastActivity int        `pdl:"last_activity"`
	Payload      string     `pdl:"payload"`
	UserAgent    string     `pdl:"user_agent"`
	UserId       int64      `pdl:"user_id"`
}

type SessionRowFactory struct{}

var SessionRow = SessionRowFactory{}

func (factory SessionRowFactory) New() *SessionRowRecord {
	result := &SessionRowRecord{
		Row: pdlgo.NewRow("sessions", "id"),
	}
	return result
}

func (factory SessionRowFactory) WithStore(store pdlgo.DBStore) *SessionRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *SessionRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *SessionRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *SessionRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type SessionRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func SessionRowWhere() SessionRowWhereBuilder {
	result := SessionRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("sessions", nil)}
	return result
}

func SessionRowWhereWithStore(store pdlgo.DBStore) SessionRowWhereBuilder {
	result := SessionRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("sessions", store)}
	return result
}

func (builder SessionRowWhereBuilder) Id(value string) SessionRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder SessionRowWhereBuilder) IpAddress(value string) SessionRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("ip_address", pdlgo.OpEq, value)
	return builder
}
func (builder SessionRowWhereBuilder) LastActivity(value int) SessionRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("last_activity", pdlgo.OpEq, value)
	return builder
}
func (builder SessionRowWhereBuilder) Payload(value string) SessionRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("payload", pdlgo.OpEq, value)
	return builder
}
func (builder SessionRowWhereBuilder) UserAgent(value string) SessionRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_agent", pdlgo.OpEq, value)
	return builder
}
func (builder SessionRowWhereBuilder) UserId(value int64) SessionRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}

func (builder SessionRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder SessionRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type SessionRowColumnsDefinition struct {
	Id           string
	IpAddress    string
	LastActivity string
	Payload      string
	UserAgent    string
	UserId       string
}

type SessionRowOrderByDefinition struct {
	Id           string
	IpAddress    string
	LastActivity string
	Payload      string
	UserAgent    string
	UserId       string
}

var SessionRowColumns = SessionRowColumnsDefinition{
	Id:           "id",
	IpAddress:    "ip_address",
	LastActivity: "last_activity",
	Payload:      "payload",
	UserAgent:    "user_agent",
	UserId:       "user_id",
}

var SessionRowOrderBy = SessionRowOrderByDefinition{
	Id:           "id",
	IpAddress:    "ip_address",
	LastActivity: "last_activity",
	Payload:      "payload",
	UserAgent:    "user_agent",
	UserId:       "user_id",
}

func SessionRowColumnList() []string {
	result := []string{
		"id",
		"ip_address",
		"last_activity",
		"payload",
		"user_agent",
		"user_id",
	}
	return result
}
