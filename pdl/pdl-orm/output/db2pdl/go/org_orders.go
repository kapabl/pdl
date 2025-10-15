package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type OrgOrderRowRecord struct {
	Row                 *pdlgo.Row `pdl:"-"`
	BillingAddress      string     `pdl:"billing_address"`
	BillingCity         string     `pdl:"billing_city"`
	BillingDiscount     int        `pdl:"billing_discount"`
	BillingDiscountCode string     `pdl:"billing_discount_code"`
	BillingEmail        string     `pdl:"billing_email"`
	BillingName         string     `pdl:"billing_name"`
	BillingNameOnCard   string     `pdl:"billing_name_on_card"`
	BillingPhone        string     `pdl:"billing_phone"`
	BillingPostalcode   string     `pdl:"billing_postalcode"`
	BillingProvince     string     `pdl:"billing_province"`
	BillingSubtotal     int        `pdl:"billing_subtotal"`
	BillingTax          int        `pdl:"billing_tax"`
	BillingTotal        int        `pdl:"billing_total"`
	CreatedAt           time.Time  `pdl:"created_at"`
	Error               string     `pdl:"error"`
	Id                  int        `pdl:"id"`
	PaymentGateway      string     `pdl:"payment_gateway"`
	Shipped             int        `pdl:"shipped"`
	UpdatedAt           time.Time  `pdl:"updated_at"`
	UserId              int        `pdl:"user_id"`
}

type OrgOrderRowFactory struct{}

var OrgOrderRow = OrgOrderRowFactory{}

func (factory OrgOrderRowFactory) New() *OrgOrderRowRecord {
	result := &OrgOrderRowRecord{
		Row: pdlgo.NewRow("org_orders", "id"),
	}
	return result
}

func (factory OrgOrderRowFactory) WithStore(store pdlgo.DBStore) *OrgOrderRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *OrgOrderRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *OrgOrderRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *OrgOrderRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type OrgOrderRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func OrgOrderRowWhere() OrgOrderRowWhereBuilder {
	result := OrgOrderRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("org_orders", nil)}
	return result
}

func OrgOrderRowWhereWithStore(store pdlgo.DBStore) OrgOrderRowWhereBuilder {
	result := OrgOrderRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("org_orders", store)}
	return result
}

func (builder OrgOrderRowWhereBuilder) BillingAddress(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_address", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingCity(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_city", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingDiscount(value int) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_discount", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingDiscountCode(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_discount_code", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingEmail(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_email", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingName(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_name", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingNameOnCard(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_name_on_card", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingPhone(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_phone", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingPostalcode(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_postalcode", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingProvince(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_province", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingSubtotal(value int) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_subtotal", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingTax(value int) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_tax", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) BillingTotal(value int) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("billing_total", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) CreatedAt(value time.Time) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) Error(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("error", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) Id(value int) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) PaymentGateway(value string) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("payment_gateway", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) Shipped(value int) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("shipped", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) UpdatedAt(value time.Time) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder OrgOrderRowWhereBuilder) UserId(value int) OrgOrderRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("user_id", pdlgo.OpEq, value)
	return builder
}

func (builder OrgOrderRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder OrgOrderRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type OrgOrderRowColumnsDefinition struct {
	BillingAddress      string
	BillingCity         string
	BillingDiscount     string
	BillingDiscountCode string
	BillingEmail        string
	BillingName         string
	BillingNameOnCard   string
	BillingPhone        string
	BillingPostalcode   string
	BillingProvince     string
	BillingSubtotal     string
	BillingTax          string
	BillingTotal        string
	CreatedAt           string
	Error               string
	Id                  string
	PaymentGateway      string
	Shipped             string
	UpdatedAt           string
	UserId              string
}

type OrgOrderRowOrderByDefinition struct {
	BillingAddress      string
	BillingCity         string
	BillingDiscount     string
	BillingDiscountCode string
	BillingEmail        string
	BillingName         string
	BillingNameOnCard   string
	BillingPhone        string
	BillingPostalcode   string
	BillingProvince     string
	BillingSubtotal     string
	BillingTax          string
	BillingTotal        string
	CreatedAt           string
	Error               string
	Id                  string
	PaymentGateway      string
	Shipped             string
	UpdatedAt           string
	UserId              string
}

var OrgOrderRowColumns = OrgOrderRowColumnsDefinition{
	BillingAddress:      "billing_address",
	BillingCity:         "billing_city",
	BillingDiscount:     "billing_discount",
	BillingDiscountCode: "billing_discount_code",
	BillingEmail:        "billing_email",
	BillingName:         "billing_name",
	BillingNameOnCard:   "billing_name_on_card",
	BillingPhone:        "billing_phone",
	BillingPostalcode:   "billing_postalcode",
	BillingProvince:     "billing_province",
	BillingSubtotal:     "billing_subtotal",
	BillingTax:          "billing_tax",
	BillingTotal:        "billing_total",
	CreatedAt:           "created_at",
	Error:               "error",
	Id:                  "id",
	PaymentGateway:      "payment_gateway",
	Shipped:             "shipped",
	UpdatedAt:           "updated_at",
	UserId:              "user_id",
}

var OrgOrderRowOrderBy = OrgOrderRowOrderByDefinition{
	BillingAddress:      "billing_address",
	BillingCity:         "billing_city",
	BillingDiscount:     "billing_discount",
	BillingDiscountCode: "billing_discount_code",
	BillingEmail:        "billing_email",
	BillingName:         "billing_name",
	BillingNameOnCard:   "billing_name_on_card",
	BillingPhone:        "billing_phone",
	BillingPostalcode:   "billing_postalcode",
	BillingProvince:     "billing_province",
	BillingSubtotal:     "billing_subtotal",
	BillingTax:          "billing_tax",
	BillingTotal:        "billing_total",
	CreatedAt:           "created_at",
	Error:               "error",
	Id:                  "id",
	PaymentGateway:      "payment_gateway",
	Shipped:             "shipped",
	UpdatedAt:           "updated_at",
	UserId:              "user_id",
}

func OrgOrderRowColumnList() []string {
	result := []string{
		"billing_address",
		"billing_city",
		"billing_discount",
		"billing_discount_code",
		"billing_email",
		"billing_name",
		"billing_name_on_card",
		"billing_phone",
		"billing_postalcode",
		"billing_province",
		"billing_subtotal",
		"billing_tax",
		"billing_total",
		"created_at",
		"error",
		"id",
		"payment_gateway",
		"shipped",
		"updated_at",
		"user_id",
	}
	return result
}
