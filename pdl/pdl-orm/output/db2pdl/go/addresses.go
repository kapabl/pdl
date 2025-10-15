package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type AddressRowRecord struct {
	Row             *pdlgo.Row `pdl:"-"`
	Address1        string     `pdl:"address1"`
	Address2        string     `pdl:"address2"`
	City            string     `pdl:"city"`
	Country         string     `pdl:"country"`
	CreatedAt       time.Time  `pdl:"created_at"`
	DefaultDelivery string     `pdl:"default_delivery"`
	DefaultPickup   string     `pdl:"default_pickup"`
	Id              int64      `pdl:"id"`
	IsTest          string     `pdl:"is_test"`
	Lat             float64    `pdl:"lat"`
	Lon             float64    `pdl:"lon"`
	Name            string     `pdl:"name"`
	Phone           string     `pdl:"phone"`
	State           string     `pdl:"state"`
	Status          string     `pdl:"status"`
	UpdatedAt       time.Time  `pdl:"updated_at"`
	UserId          int        `pdl:"user_id"`
	Zipcode         string     `pdl:"zipcode"`
}

type AddressRowFactory struct{}

var AddressRow = AddressRowFactory{}

func (factory AddressRowFactory) New() *AddressRowRecord {
	result := &AddressRowRecord{
		Row: pdlgo.NewRow("addresses", "id"),
	}
	return result
}

func (factory AddressRowFactory) WithStore(store pdlgo.DBStore) *AddressRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *AddressRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *AddressRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *AddressRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type AddressRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func AddressRowWhere() AddressRowWhereBuilder {
	result := AddressRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("addresses", nil)}
	return result
}

func AddressRowWhereWithStore(store pdlgo.DBStore) AddressRowWhereBuilder {
	result := AddressRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("addresses", store)}
	return result
}

func (builder AddressRowWhereBuilder) Address1(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("address1", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Address2(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("address2", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) City(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("city", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Country(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("country", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) CreatedAt(value time.Time) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) DefaultDelivery(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("default_delivery", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) DefaultPickup(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("default_pickup", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Id(value int64) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) IsTest(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_test", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Lat(value float64) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("lat", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Lon(value float64) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("lon", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Name(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Phone(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("phone", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) State(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("state", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Status(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) UpdatedAt(value time.Time) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) UserId(value int) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}
func (builder AddressRowWhereBuilder) Zipcode(value string) AddressRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("zipcode", pdlgo.OpEq, value)
	return builder
}

func (builder AddressRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder AddressRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type AddressRowColumnsDefinition struct {
	Address1        string
	Address2        string
	City            string
	Country         string
	CreatedAt       string
	DefaultDelivery string
	DefaultPickup   string
	Id              string
	IsTest          string
	Lat             string
	Lon             string
	Name            string
	Phone           string
	State           string
	Status          string
	UpdatedAt       string
	UserId          string
	Zipcode         string
}

type AddressRowOrderByDefinition struct {
	Address1        string
	Address2        string
	City            string
	Country         string
	CreatedAt       string
	DefaultDelivery string
	DefaultPickup   string
	Id              string
	IsTest          string
	Lat             string
	Lon             string
	Name            string
	Phone           string
	State           string
	Status          string
	UpdatedAt       string
	UserId          string
	Zipcode         string
}

var AddressRowColumns = AddressRowColumnsDefinition{
	Address1:        "address1",
	Address2:        "address2",
	City:            "city",
	Country:         "country",
	CreatedAt:       "created_at",
	DefaultDelivery: "default_delivery",
	DefaultPickup:   "default_pickup",
	Id:              "id",
	IsTest:          "is_test",
	Lat:             "lat",
	Lon:             "lon",
	Name:            "name",
	Phone:           "phone",
	State:           "state",
	Status:          "status",
	UpdatedAt:       "updated_at",
	UserId:          "user_id",
	Zipcode:         "zipcode",
}

var AddressRowOrderBy = AddressRowOrderByDefinition{
	Address1:        "address1",
	Address2:        "address2",
	City:            "city",
	Country:         "country",
	CreatedAt:       "created_at",
	DefaultDelivery: "default_delivery",
	DefaultPickup:   "default_pickup",
	Id:              "id",
	IsTest:          "is_test",
	Lat:             "lat",
	Lon:             "lon",
	Name:            "name",
	Phone:           "phone",
	State:           "state",
	Status:          "status",
	UpdatedAt:       "updated_at",
	UserId:          "user_id",
	Zipcode:         "zipcode",
}

func AddressRowColumnList() []string {
	result := []string{
		"address1",
		"address2",
		"city",
		"country",
		"created_at",
		"default_delivery",
		"default_pickup",
		"id",
		"is_test",
		"lat",
		"lon",
		"name",
		"phone",
		"state",
		"status",
		"updated_at",
		"user_id",
		"zipcode",
	}
	return result
}
