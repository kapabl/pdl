package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type CourierLocationRowRecord struct {
	Row              *pdlgo.Row `pdl:"-"`
	Accuracy         float64    `pdl:"accuracy"`
	Altitude         float64    `pdl:"altitude"`
	AltitudeAccuracy float64    `pdl:"altitude_accuracy"`
	CourierUuid      string     `pdl:"courier_uuid"`
	Heading          float64    `pdl:"heading"`
	Id               int64      `pdl:"id"`
	Lat              float64    `pdl:"lat"`
	LocationTime     time.Time  `pdl:"location_time"`
	Lon              float64    `pdl:"lon"`
	Speed            float64    `pdl:"speed"`
}

type CourierLocationRowFactory struct{}

var CourierLocationRow = CourierLocationRowFactory{}

func (factory CourierLocationRowFactory) New() *CourierLocationRowRecord {
	result := &CourierLocationRowRecord{
		Row: pdlgo.NewRow("courier_locations", "id"),
	}
	return result
}

func (factory CourierLocationRowFactory) WithStore(store pdlgo.DBStore) *CourierLocationRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *CourierLocationRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *CourierLocationRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *CourierLocationRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type CourierLocationRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func CourierLocationRowWhere() CourierLocationRowWhereBuilder {
	result := CourierLocationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("courier_locations", nil)}
	return result
}

func CourierLocationRowWhereWithStore(store pdlgo.DBStore) CourierLocationRowWhereBuilder {
	result := CourierLocationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("courier_locations", store)}
	return result
}

func (builder CourierLocationRowWhereBuilder) Accuracy(value float64) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("accuracy", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) Altitude(value float64) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("altitude", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) AltitudeAccuracy(value float64) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("altitude_accuracy", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) CourierUuid(value string) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("courier_uuid", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) Heading(value float64) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("heading", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) Id(value int64) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) Lat(value float64) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("lat", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) LocationTime(value time.Time) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("location_time", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) Lon(value float64) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("lon", pdlgo.OpEq, value)
	return builder
}
func (builder CourierLocationRowWhereBuilder) Speed(value float64) CourierLocationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("speed", pdlgo.OpEq, value)
	return builder
}

func (builder CourierLocationRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder CourierLocationRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type CourierLocationRowColumnsDefinition struct {
	Accuracy         string
	Altitude         string
	AltitudeAccuracy string
	CourierUuid      string
	Heading          string
	Id               string
	Lat              string
	LocationTime     string
	Lon              string
	Speed            string
}

type CourierLocationRowOrderByDefinition struct {
	Accuracy         string
	Altitude         string
	AltitudeAccuracy string
	CourierUuid      string
	Heading          string
	Id               string
	Lat              string
	LocationTime     string
	Lon              string
	Speed            string
}

var CourierLocationRowColumns = CourierLocationRowColumnsDefinition{
	Accuracy:         "accuracy",
	Altitude:         "altitude",
	AltitudeAccuracy: "altitude_accuracy",
	CourierUuid:      "courier_uuid",
	Heading:          "heading",
	Id:               "id",
	Lat:              "lat",
	LocationTime:     "location_time",
	Lon:              "lon",
	Speed:            "speed",
}

var CourierLocationRowOrderBy = CourierLocationRowOrderByDefinition{
	Accuracy:         "accuracy",
	Altitude:         "altitude",
	AltitudeAccuracy: "altitude_accuracy",
	CourierUuid:      "courier_uuid",
	Heading:          "heading",
	Id:               "id",
	Lat:              "lat",
	LocationTime:     "location_time",
	Lon:              "lon",
	Speed:            "speed",
}

func CourierLocationRowColumnList() []string {
	result := []string{
		"accuracy",
		"altitude",
		"altitude_accuracy",
		"courier_uuid",
		"heading",
		"id",
		"lat",
		"location_time",
		"lon",
		"speed",
	}
	return result
}
