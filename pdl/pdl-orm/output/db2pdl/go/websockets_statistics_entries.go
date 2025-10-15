package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type WebsocketsStatisticsEntryRowRecord struct {
	Row                   *pdlgo.Row `pdl:"-"`
	ApiMessageCount       int        `pdl:"api_message_count"`
	AppId                 string     `pdl:"app_id"`
	CreatedAt             time.Time  `pdl:"created_at"`
	Id                    int        `pdl:"id"`
	PeakConnectionCount   int        `pdl:"peak_connection_count"`
	UpdatedAt             time.Time  `pdl:"updated_at"`
	WebsocketMessageCount int        `pdl:"websocket_message_count"`
}

type WebsocketsStatisticsEntryRowFactory struct{}

var WebsocketsStatisticsEntryRow = WebsocketsStatisticsEntryRowFactory{}

func (factory WebsocketsStatisticsEntryRowFactory) New() *WebsocketsStatisticsEntryRowRecord {
	result := &WebsocketsStatisticsEntryRowRecord{
		Row: pdlgo.NewRow("websockets_statistics_entries", "id"),
	}
	return result
}

func (factory WebsocketsStatisticsEntryRowFactory) WithStore(store pdlgo.DBStore) *WebsocketsStatisticsEntryRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *WebsocketsStatisticsEntryRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *WebsocketsStatisticsEntryRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *WebsocketsStatisticsEntryRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type WebsocketsStatisticsEntryRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func WebsocketsStatisticsEntryRowWhere() WebsocketsStatisticsEntryRowWhereBuilder {
	result := WebsocketsStatisticsEntryRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("websockets_statistics_entries", nil)}
	return result
}

func WebsocketsStatisticsEntryRowWhereWithStore(store pdlgo.DBStore) WebsocketsStatisticsEntryRowWhereBuilder {
	result := WebsocketsStatisticsEntryRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("websockets_statistics_entries", store)}
	return result
}

func (builder WebsocketsStatisticsEntryRowWhereBuilder) ApiMessageCount(value int) WebsocketsStatisticsEntryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("api_message_count", pdlgo.OpEq, value)
	return builder
}
func (builder WebsocketsStatisticsEntryRowWhereBuilder) AppId(value string) WebsocketsStatisticsEntryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("app_id", pdlgo.OpEq, value)
	return builder
}
func (builder WebsocketsStatisticsEntryRowWhereBuilder) CreatedAt(value time.Time) WebsocketsStatisticsEntryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder WebsocketsStatisticsEntryRowWhereBuilder) Id(value int) WebsocketsStatisticsEntryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder WebsocketsStatisticsEntryRowWhereBuilder) PeakConnectionCount(value int) WebsocketsStatisticsEntryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("peak_connection_count", pdlgo.OpEq, value)
	return builder
}
func (builder WebsocketsStatisticsEntryRowWhereBuilder) UpdatedAt(value time.Time) WebsocketsStatisticsEntryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder WebsocketsStatisticsEntryRowWhereBuilder) WebsocketMessageCount(value int) WebsocketsStatisticsEntryRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("websocket_message_count", pdlgo.OpEq, value)
	return builder
}

func (builder WebsocketsStatisticsEntryRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder WebsocketsStatisticsEntryRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type WebsocketsStatisticsEntryRowColumnsDefinition struct {
	ApiMessageCount       string
	AppId                 string
	CreatedAt             string
	Id                    string
	PeakConnectionCount   string
	UpdatedAt             string
	WebsocketMessageCount string
}

type WebsocketsStatisticsEntryRowOrderByDefinition struct {
	ApiMessageCount       string
	AppId                 string
	CreatedAt             string
	Id                    string
	PeakConnectionCount   string
	UpdatedAt             string
	WebsocketMessageCount string
}

var WebsocketsStatisticsEntryRowColumns = WebsocketsStatisticsEntryRowColumnsDefinition{
	ApiMessageCount:       "api_message_count",
	AppId:                 "app_id",
	CreatedAt:             "created_at",
	Id:                    "id",
	PeakConnectionCount:   "peak_connection_count",
	UpdatedAt:             "updated_at",
	WebsocketMessageCount: "websocket_message_count",
}

var WebsocketsStatisticsEntryRowOrderBy = WebsocketsStatisticsEntryRowOrderByDefinition{
	ApiMessageCount:       "api_message_count",
	AppId:                 "app_id",
	CreatedAt:             "created_at",
	Id:                    "id",
	PeakConnectionCount:   "peak_connection_count",
	UpdatedAt:             "updated_at",
	WebsocketMessageCount: "websocket_message_count",
}

func WebsocketsStatisticsEntryRowColumnList() []string {
	result := []string{
		"api_message_count",
		"app_id",
		"created_at",
		"id",
		"peak_connection_count",
		"updated_at",
		"websocket_message_count",
	}
	return result
}
