package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type DiscountRowRecord struct {
	Row          *pdlgo.Row `pdl:"-"`
	Code         string     `pdl:"code"`
	CreatedAt    time.Time  `pdl:"created_at"`
	Currency     string     `pdl:"currency"`
	DatetimeZone string     `pdl:"datetime_zone"`
	Description  string     `pdl:"description"`
	Duration     int        `pdl:"duration"`
	EndDate      time.Time  `pdl:"end_date"`
	Id           int64      `pdl:"id"`
	Locale       string     `pdl:"locale"`
	MaxAmount    float64    `pdl:"max_amount"`
	MinAmount    float64    `pdl:"min_amount"`
	Name         string     `pdl:"name"`
	Scope        string     `pdl:"scope"`
	StartDate    time.Time  `pdl:"start_date"`
	Status       string     `pdl:"status"`
	Target       string     `pdl:"target"`
	Type         string     `pdl:"type"`
	UpdatedAt    time.Time  `pdl:"updated_at"`
	Value        float64    `pdl:"value"`
}

type DiscountRowFactory struct{}

var DiscountRow = DiscountRowFactory{}

func (factory DiscountRowFactory) New() *DiscountRowRecord {
	result := &DiscountRowRecord{
		Row: pdlgo.NewRow("discounts", "id"),
	}
	return result
}

func (factory DiscountRowFactory) WithStore(store pdlgo.DBStore) *DiscountRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *DiscountRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *DiscountRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *DiscountRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type DiscountRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func DiscountRowWhere() DiscountRowWhereBuilder {
	result := DiscountRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("discounts", nil)}
	return result
}

func DiscountRowWhereWithStore(store pdlgo.DBStore) DiscountRowWhereBuilder {
	result := DiscountRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("discounts", store)}
	return result
}

func (builder DiscountRowWhereBuilder) Code(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("code", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) CreatedAt(value time.Time) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Currency(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("currency", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) DatetimeZone(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("datetime_zone", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Description(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("description", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Duration(value int) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("duration", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) EndDate(value time.Time) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("end_date", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Id(value int64) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Locale(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("locale", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) MaxAmount(value float64) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("max_amount", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) MinAmount(value float64) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("min_amount", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Name(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Scope(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("scope", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) StartDate(value time.Time) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("start_date", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Status(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Target(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("target", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Type(value string) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("type", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) UpdatedAt(value time.Time) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder DiscountRowWhereBuilder) Value(value float64) DiscountRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("value", pdlgo.OpEq, value)
	return builder
}

func (builder DiscountRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder DiscountRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type DiscountRowColumnsDefinition struct {
	Code         string
	CreatedAt    string
	Currency     string
	DatetimeZone string
	Description  string
	Duration     string
	EndDate      string
	Id           string
	Locale       string
	MaxAmount    string
	MinAmount    string
	Name         string
	Scope        string
	StartDate    string
	Status       string
	Target       string
	Type         string
	UpdatedAt    string
	Value        string
}

type DiscountRowOrderByDefinition struct {
	Code         string
	CreatedAt    string
	Currency     string
	DatetimeZone string
	Description  string
	Duration     string
	EndDate      string
	Id           string
	Locale       string
	MaxAmount    string
	MinAmount    string
	Name         string
	Scope        string
	StartDate    string
	Status       string
	Target       string
	Type         string
	UpdatedAt    string
	Value        string
}

var DiscountRowColumns = DiscountRowColumnsDefinition{
	Code:         "code",
	CreatedAt:    "created_at",
	Currency:     "currency",
	DatetimeZone: "datetime_zone",
	Description:  "description",
	Duration:     "duration",
	EndDate:      "end_date",
	Id:           "id",
	Locale:       "locale",
	MaxAmount:    "max_amount",
	MinAmount:    "min_amount",
	Name:         "name",
	Scope:        "scope",
	StartDate:    "start_date",
	Status:       "status",
	Target:       "target",
	Type:         "type",
	UpdatedAt:    "updated_at",
	Value:        "value",
}

var DiscountRowOrderBy = DiscountRowOrderByDefinition{
	Code:         "code",
	CreatedAt:    "created_at",
	Currency:     "currency",
	DatetimeZone: "datetime_zone",
	Description:  "description",
	Duration:     "duration",
	EndDate:      "end_date",
	Id:           "id",
	Locale:       "locale",
	MaxAmount:    "max_amount",
	MinAmount:    "min_amount",
	Name:         "name",
	Scope:        "scope",
	StartDate:    "start_date",
	Status:       "status",
	Target:       "target",
	Type:         "type",
	UpdatedAt:    "updated_at",
	Value:        "value",
}

func DiscountRowColumnList() []string {
	result := []string{
		"code",
		"created_at",
		"currency",
		"datetime_zone",
		"description",
		"duration",
		"end_date",
		"id",
		"locale",
		"max_amount",
		"min_amount",
		"name",
		"scope",
		"start_date",
		"status",
		"target",
		"type",
		"updated_at",
		"value",
	}
	return result
}
