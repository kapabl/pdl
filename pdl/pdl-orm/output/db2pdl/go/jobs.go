package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
)

type JobRowRecord struct {
	Row         *pdlgo.Row `pdl:"-"`
	Attempts    int        `pdl:"attempts"`
	AvailableAt int        `pdl:"available_at"`
	CreatedAt   int        `pdl:"created_at"`
	Id          int64      `pdl:"id"`
	Payload     string     `pdl:"payload"`
	Queue       string     `pdl:"queue"`
	ReservedAt  int        `pdl:"reserved_at"`
}

type JobRowFactory struct{}

var JobRow = JobRowFactory{}

func (factory JobRowFactory) New() *JobRowRecord {
	result := &JobRowRecord{
		Row: pdlgo.NewRow("jobs", "id"),
	}
	return result
}

func (factory JobRowFactory) WithStore(store pdlgo.DBStore) *JobRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *JobRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *JobRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *JobRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type JobRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func JobRowWhere() JobRowWhereBuilder {
	result := JobRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("jobs", nil)}
	return result
}

func JobRowWhereWithStore(store pdlgo.DBStore) JobRowWhereBuilder {
	result := JobRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("jobs", store)}
	return result
}

func (builder JobRowWhereBuilder) Attempts(value int) JobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("attempts", pdlgo.OpEq, value)
	return builder
}
func (builder JobRowWhereBuilder) AvailableAt(value int) JobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("available_at", pdlgo.OpEq, value)
	return builder
}
func (builder JobRowWhereBuilder) CreatedAt(value int) JobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder JobRowWhereBuilder) Id(value int64) JobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder JobRowWhereBuilder) Payload(value string) JobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("payload", pdlgo.OpEq, value)
	return builder
}
func (builder JobRowWhereBuilder) Queue(value string) JobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("queue", pdlgo.OpEq, value)
	return builder
}
func (builder JobRowWhereBuilder) ReservedAt(value int) JobRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("reserved_at", pdlgo.OpEq, value)
	return builder
}

func (builder JobRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder JobRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type JobRowColumnsDefinition struct {
	Attempts    string
	AvailableAt string
	CreatedAt   string
	Id          string
	Payload     string
	Queue       string
	ReservedAt  string
}

type JobRowOrderByDefinition struct {
	Attempts    string
	AvailableAt string
	CreatedAt   string
	Id          string
	Payload     string
	Queue       string
	ReservedAt  string
}

var JobRowColumns = JobRowColumnsDefinition{
	Attempts:    "attempts",
	AvailableAt: "available_at",
	CreatedAt:   "created_at",
	Id:          "id",
	Payload:     "payload",
	Queue:       "queue",
	ReservedAt:  "reserved_at",
}

var JobRowOrderBy = JobRowOrderByDefinition{
	Attempts:    "attempts",
	AvailableAt: "available_at",
	CreatedAt:   "created_at",
	Id:          "id",
	Payload:     "payload",
	Queue:       "queue",
	ReservedAt:  "reserved_at",
}

func JobRowColumnList() []string {
	result := []string{
		"attempts",
		"available_at",
		"created_at",
		"id",
		"payload",
		"queue",
		"reserved_at",
	}
	return result
}
