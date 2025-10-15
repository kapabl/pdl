package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type FailedJobRowRecord struct {
	Row        *pdlgo.Row `pdl:"-"`
	Connection string     `pdl:"connection"`
	Exception  string     `pdl:"exception"`
	FailedAt   time.Time  `pdl:"failed_at"`
	Id         int64      `pdl:"id"`
	Payload    string     `pdl:"payload"`
	Queue      string     `pdl:"queue"`
	Uuid       string     `pdl:"uuid"`
}

type FailedJobRowFactory struct{}

var FailedJobRow = FailedJobRowFactory{}

func (factory FailedJobRowFactory) New() *FailedJobRowRecord {
	result := &FailedJobRowRecord{
		Row: pdlgo.NewRow("failed_jobs", "id"),
	}
	return result
}

func (factory FailedJobRowFactory) WithStore(store pdlgo.DBStore) *FailedJobRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *FailedJobRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *FailedJobRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *FailedJobRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type FailedJobRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func FailedJobRowWhere() FailedJobRowWhereBuilder {
	result := FailedJobRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("failed_jobs", nil)}
	return result
}

func FailedJobRowWhereWithStore(store pdlgo.DBStore) FailedJobRowWhereBuilder {
	result := FailedJobRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("failed_jobs", store)}
	return result
}

func (builder FailedJobRowWhereBuilder) Connection(value string) FailedJobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("connection", pdlgo.OpEq, value)
	return builder
}
func (builder FailedJobRowWhereBuilder) Exception(value string) FailedJobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("exception", pdlgo.OpEq, value)
	return builder
}
func (builder FailedJobRowWhereBuilder) FailedAt(value time.Time) FailedJobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("failed_at", pdlgo.OpEq, value)
	return builder
}
func (builder FailedJobRowWhereBuilder) Id(value int64) FailedJobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder FailedJobRowWhereBuilder) Payload(value string) FailedJobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("payload", pdlgo.OpEq, value)
	return builder
}
func (builder FailedJobRowWhereBuilder) Queue(value string) FailedJobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("queue", pdlgo.OpEq, value)
	return builder
}
func (builder FailedJobRowWhereBuilder) Uuid(value string) FailedJobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("uuid", pdlgo.OpEq, value)
	return builder
}

func (builder FailedJobRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder FailedJobRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type FailedJobRowColumnsDefinition struct {
	Connection string
	Exception  string
	FailedAt   string
	Id         string
	Payload    string
	Queue      string
	Uuid       string
}

type FailedJobRowOrderByDefinition struct {
	Connection string
	Exception  string
	FailedAt   string
	Id         string
	Payload    string
	Queue      string
	Uuid       string
}

var FailedJobRowColumns = FailedJobRowColumnsDefinition{
	Connection: "connection",
	Exception:  "exception",
	FailedAt:   "failed_at",
	Id:         "id",
	Payload:    "payload",
	Queue:      "queue",
	Uuid:       "uuid",
}

var FailedJobRowOrderBy = FailedJobRowOrderByDefinition{
	Connection: "connection",
	Exception:  "exception",
	FailedAt:   "failed_at",
	Id:         "id",
	Payload:    "payload",
	Queue:      "queue",
	Uuid:       "uuid",
}

func FailedJobRowColumnList() []string {
	result := []string{
		"connection",
		"exception",
		"failed_at",
		"id",
		"payload",
		"queue",
		"uuid",
	}
	return result
}
