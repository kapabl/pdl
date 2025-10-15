package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type NotificationRowRecord struct {
	Row                *pdlgo.Row `pdl:"-"`
	Body               string     `pdl:"body"`
	CreatedAt          time.Time  `pdl:"created_at"`
	Data               string     `pdl:"data"`
	FromStoreId        int        `pdl:"from_store_Id"`
	FromUserId         int        `pdl:"from_user_id"`
	Id                 int64      `pdl:"id"`
	Name               string     `pdl:"name"`
	NotificationTime   time.Time  `pdl:"notification_time"`
	OrderId            string     `pdl:"order_id"`
	ReadTime           time.Time  `pdl:"read_time"`
	RetrievedFirstTime time.Time  `pdl:"retrieved_first_time"`
	Source             string     `pdl:"source"`
	Title              string     `pdl:"title"`
	ToStoreId          int        `pdl:"to_store_id"`
	ToUserId           int        `pdl:"to_user_id"`
	Type               string     `pdl:"type"`
	Unread             string     `pdl:"unread"`
	UpdatedAt          time.Time  `pdl:"updated_at"`
	Url                string     `pdl:"url"`
}

type NotificationRowFactory struct{}

var NotificationRow = NotificationRowFactory{}

func (factory NotificationRowFactory) New() *NotificationRowRecord {
	result := &NotificationRowRecord{
		Row: pdlgo.NewRow("notifications", "id"),
	}
	return result
}

func (factory NotificationRowFactory) WithStore(store pdlgo.DBStore) *NotificationRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *NotificationRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *NotificationRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *NotificationRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type NotificationRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func NotificationRowWhere() NotificationRowWhereBuilder {
	result := NotificationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("notifications", nil)}
	return result
}

func NotificationRowWhereWithStore(store pdlgo.DBStore) NotificationRowWhereBuilder {
	result := NotificationRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("notifications", store)}
	return result
}

func (builder NotificationRowWhereBuilder) Body(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("body", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) CreatedAt(value time.Time) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) Data(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("data", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) FromStoreId(value int) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("from_store_Id", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) FromUserId(value int) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("from_user_id", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) Id(value int64) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) Name(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) NotificationTime(value time.Time) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("notification_time", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) OrderId(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("order_id", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) ReadTime(value time.Time) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("read_time", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) RetrievedFirstTime(value time.Time) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("retrieved_first_time", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) Source(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("source", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) Title(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("title", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) ToStoreId(value int) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("to_store_id", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) ToUserId(value int) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("to_user_id", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) Type(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("type", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) Unread(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("unread", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) UpdatedAt(value time.Time) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder NotificationRowWhereBuilder) Url(value string) NotificationRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("url", pdlgo.OpEq, value)
	return builder
}

func (builder NotificationRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder NotificationRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type NotificationRowColumnsDefinition struct {
	Body               string
	CreatedAt          string
	Data               string
	FromStoreId        string
	FromUserId         string
	Id                 string
	Name               string
	NotificationTime   string
	OrderId            string
	ReadTime           string
	RetrievedFirstTime string
	Source             string
	Title              string
	ToStoreId          string
	ToUserId           string
	Type               string
	Unread             string
	UpdatedAt          string
	Url                string
}

type NotificationRowOrderByDefinition struct {
	Body               string
	CreatedAt          string
	Data               string
	FromStoreId        string
	FromUserId         string
	Id                 string
	Name               string
	NotificationTime   string
	OrderId            string
	ReadTime           string
	RetrievedFirstTime string
	Source             string
	Title              string
	ToStoreId          string
	ToUserId           string
	Type               string
	Unread             string
	UpdatedAt          string
	Url                string
}

var NotificationRowColumns = NotificationRowColumnsDefinition{
	Body:               "body",
	CreatedAt:          "created_at",
	Data:               "data",
	FromStoreId:        "from_store_Id",
	FromUserId:         "from_user_id",
	Id:                 "id",
	Name:               "name",
	NotificationTime:   "notification_time",
	OrderId:            "order_id",
	ReadTime:           "read_time",
	RetrievedFirstTime: "retrieved_first_time",
	Source:             "source",
	Title:              "title",
	ToStoreId:          "to_store_id",
	ToUserId:           "to_user_id",
	Type:               "type",
	Unread:             "unread",
	UpdatedAt:          "updated_at",
	Url:                "url",
}

var NotificationRowOrderBy = NotificationRowOrderByDefinition{
	Body:               "body",
	CreatedAt:          "created_at",
	Data:               "data",
	FromStoreId:        "from_store_Id",
	FromUserId:         "from_user_id",
	Id:                 "id",
	Name:               "name",
	NotificationTime:   "notification_time",
	OrderId:            "order_id",
	ReadTime:           "read_time",
	RetrievedFirstTime: "retrieved_first_time",
	Source:             "source",
	Title:              "title",
	ToStoreId:          "to_store_id",
	ToUserId:           "to_user_id",
	Type:               "type",
	Unread:             "unread",
	UpdatedAt:          "updated_at",
	Url:                "url",
}

func NotificationRowColumnList() []string {
	result := []string{
		"body",
		"created_at",
		"data",
		"from_store_Id",
		"from_user_id",
		"id",
		"name",
		"notification_time",
		"order_id",
		"read_time",
		"retrieved_first_time",
		"source",
		"title",
		"to_store_id",
		"to_user_id",
		"type",
		"unread",
		"updated_at",
		"url",
	}
	return result
}
