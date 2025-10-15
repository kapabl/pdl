package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type StoreRowRecord struct {
	Row                      *pdlgo.Row `pdl:"-"`
	BusinessHours            string     `pdl:"business_hours"`
	BusinessPhone            string     `pdl:"business_phone"`
	BusinessPhoto            string     `pdl:"business_photo"`
	Categories               string     `pdl:"categories"`
	CreatedAt                time.Time  `pdl:"created_at"`
	Currency                 string     `pdl:"currency"`
	DatetimeZone             string     `pdl:"datetime_zone"`
	DeliveryAddressId        int        `pdl:"delivery_address_id"`
	DeliveryCost             int        `pdl:"delivery_cost"`
	DeliveryOptions          string     `pdl:"delivery_options"`
	DeliveryTimeframeOptions string     `pdl:"delivery_timeframe_options"`
	Id                       int64      `pdl:"id"`
	IsAcceptingOrders        string     `pdl:"is_accepting_orders"`
	IsOpen                   string     `pdl:"is_open"`
	IsPublished              string     `pdl:"is_published"`
	Locale                   string     `pdl:"locale"`
	Name                     string     `pdl:"name"`
	OrderStateConfig         string     `pdl:"order_state_config"`
	PaymentOptions           string     `pdl:"payment_options"`
	PickupAddressId          int        `pdl:"pickup_address_id"`
	ServiceFeeOptions        string     `pdl:"service_fee_options"`
	Slogan                   string     `pdl:"slogan"`
	Slug                     string     `pdl:"slug"`
	Status                   string     `pdl:"status"`
	UpdatedAt                time.Time  `pdl:"updated_at"`
	UserId                   int        `pdl:"user_id"`
	Uuid                     string     `pdl:"uuid"`
}

type StoreRowFactory struct{}

var StoreRow = StoreRowFactory{}

func (factory StoreRowFactory) New() *StoreRowRecord {
	result := &StoreRowRecord{
		Row: pdlgo.NewRow("stores", "id"),
	}
	return result
}

func (factory StoreRowFactory) WithStore(store pdlgo.DBStore) *StoreRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *StoreRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *StoreRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *StoreRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type StoreRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func StoreRowWhere() StoreRowWhereBuilder {
	result := StoreRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("stores", nil)}
	return result
}

func StoreRowWhereWithStore(store pdlgo.DBStore) StoreRowWhereBuilder {
	result := StoreRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("stores", store)}
	return result
}

func (builder StoreRowWhereBuilder) BusinessHours(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("business_hours", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) BusinessPhone(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("business_phone", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) BusinessPhoto(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("business_photo", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Categories(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("categories", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) CreatedAt(value time.Time) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Currency(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("currency", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) DatetimeZone(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("datetime_zone", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) DeliveryAddressId(value int) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("delivery_address_id", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) DeliveryCost(value int) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("delivery_cost", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) DeliveryOptions(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("delivery_options", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) DeliveryTimeframeOptions(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("delivery_timeframe_options", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Id(value int64) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) IsAcceptingOrders(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_accepting_orders", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) IsOpen(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_open", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) IsPublished(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_published", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Locale(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("locale", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Name(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) OrderStateConfig(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("order_state_config", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) PaymentOptions(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("payment_options", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) PickupAddressId(value int) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("pickup_address_id", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) ServiceFeeOptions(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("service_fee_options", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Slogan(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("slogan", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Slug(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("slug", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Status(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) UpdatedAt(value time.Time) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) UserId(value int) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}
func (builder StoreRowWhereBuilder) Uuid(value string) StoreRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("uuid", pdlgo.OpEq, value)
	return builder
}

func (builder StoreRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder StoreRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type StoreRowColumnsDefinition struct {
	BusinessHours            string
	BusinessPhone            string
	BusinessPhoto            string
	Categories               string
	CreatedAt                string
	Currency                 string
	DatetimeZone             string
	DeliveryAddressId        string
	DeliveryCost             string
	DeliveryOptions          string
	DeliveryTimeframeOptions string
	Id                       string
	IsAcceptingOrders        string
	IsOpen                   string
	IsPublished              string
	Locale                   string
	Name                     string
	OrderStateConfig         string
	PaymentOptions           string
	PickupAddressId          string
	ServiceFeeOptions        string
	Slogan                   string
	Slug                     string
	Status                   string
	UpdatedAt                string
	UserId                   string
	Uuid                     string
}

type StoreRowOrderByDefinition struct {
	BusinessHours            string
	BusinessPhone            string
	BusinessPhoto            string
	Categories               string
	CreatedAt                string
	Currency                 string
	DatetimeZone             string
	DeliveryAddressId        string
	DeliveryCost             string
	DeliveryOptions          string
	DeliveryTimeframeOptions string
	Id                       string
	IsAcceptingOrders        string
	IsOpen                   string
	IsPublished              string
	Locale                   string
	Name                     string
	OrderStateConfig         string
	PaymentOptions           string
	PickupAddressId          string
	ServiceFeeOptions        string
	Slogan                   string
	Slug                     string
	Status                   string
	UpdatedAt                string
	UserId                   string
	Uuid                     string
}

var StoreRowColumns = StoreRowColumnsDefinition{
	BusinessHours:            "business_hours",
	BusinessPhone:            "business_phone",
	BusinessPhoto:            "business_photo",
	Categories:               "categories",
	CreatedAt:                "created_at",
	Currency:                 "currency",
	DatetimeZone:             "datetime_zone",
	DeliveryAddressId:        "delivery_address_id",
	DeliveryCost:             "delivery_cost",
	DeliveryOptions:          "delivery_options",
	DeliveryTimeframeOptions: "delivery_timeframe_options",
	Id:                       "id",
	IsAcceptingOrders:        "is_accepting_orders",
	IsOpen:                   "is_open",
	IsPublished:              "is_published",
	Locale:                   "locale",
	Name:                     "name",
	OrderStateConfig:         "order_state_config",
	PaymentOptions:           "payment_options",
	PickupAddressId:          "pickup_address_id",
	ServiceFeeOptions:        "service_fee_options",
	Slogan:                   "slogan",
	Slug:                     "slug",
	Status:                   "status",
	UpdatedAt:                "updated_at",
	UserId:                   "user_id",
	Uuid:                     "uuid",
}

var StoreRowOrderBy = StoreRowOrderByDefinition{
	BusinessHours:            "business_hours",
	BusinessPhone:            "business_phone",
	BusinessPhoto:            "business_photo",
	Categories:               "categories",
	CreatedAt:                "created_at",
	Currency:                 "currency",
	DatetimeZone:             "datetime_zone",
	DeliveryAddressId:        "delivery_address_id",
	DeliveryCost:             "delivery_cost",
	DeliveryOptions:          "delivery_options",
	DeliveryTimeframeOptions: "delivery_timeframe_options",
	Id:                       "id",
	IsAcceptingOrders:        "is_accepting_orders",
	IsOpen:                   "is_open",
	IsPublished:              "is_published",
	Locale:                   "locale",
	Name:                     "name",
	OrderStateConfig:         "order_state_config",
	PaymentOptions:           "payment_options",
	PickupAddressId:          "pickup_address_id",
	ServiceFeeOptions:        "service_fee_options",
	Slogan:                   "slogan",
	Slug:                     "slug",
	Status:                   "status",
	UpdatedAt:                "updated_at",
	UserId:                   "user_id",
	Uuid:                     "uuid",
}

func StoreRowColumnList() []string {
	result := []string{
		"business_hours",
		"business_phone",
		"business_photo",
		"categories",
		"created_at",
		"currency",
		"datetime_zone",
		"delivery_address_id",
		"delivery_cost",
		"delivery_options",
		"delivery_timeframe_options",
		"id",
		"is_accepting_orders",
		"is_open",
		"is_published",
		"locale",
		"name",
		"order_state_config",
		"payment_options",
		"pickup_address_id",
		"service_fee_options",
		"slogan",
		"slug",
		"status",
		"updated_at",
		"user_id",
		"uuid",
	}
	return result
}
