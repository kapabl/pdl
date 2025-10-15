package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type PasswordResetRowRecord struct {
	Row       *pdlgo.Row `pdl:"-"`
	CreatedAt time.Time  `pdl:"created_at"`
	Email     string     `pdl:"email"`
	Token     string     `pdl:"token"`
}

type PasswordResetRowFactory struct{}

var PasswordResetRow = PasswordResetRowFactory{}

func (factory PasswordResetRowFactory) New() *PasswordResetRowRecord {
	result := &PasswordResetRowRecord{
		Row: pdlgo.NewRow("password_resets", ""),
	}
	return result
}

func (factory PasswordResetRowFactory) WithStore(store pdlgo.DBStore) *PasswordResetRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *PasswordResetRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *PasswordResetRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *PasswordResetRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type PasswordResetRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func PasswordResetRowWhere() PasswordResetRowWhereBuilder {
	result := PasswordResetRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("password_resets", nil)}
	return result
}

func PasswordResetRowWhereWithStore(store pdlgo.DBStore) PasswordResetRowWhereBuilder {
	result := PasswordResetRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("password_resets", store)}
	return result
}

func (builder PasswordResetRowWhereBuilder) CreatedAt(value time.Time) PasswordResetRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder PasswordResetRowWhereBuilder) Email(value string) PasswordResetRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("email", pdlgo.OpEq, value)
	return builder
}
func (builder PasswordResetRowWhereBuilder) Token(value string) PasswordResetRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("token", pdlgo.OpEq, value)
	return builder
}

func (builder PasswordResetRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder PasswordResetRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("")
	return result
}

type PasswordResetRowColumnsDefinition struct {
	CreatedAt string
	Email     string
	Token     string
}

type PasswordResetRowOrderByDefinition struct {
	CreatedAt string
	Email     string
	Token     string
}

var PasswordResetRowColumns = PasswordResetRowColumnsDefinition{
	CreatedAt: "created_at",
	Email:     "email",
	Token:     "token",
}

var PasswordResetRowOrderBy = PasswordResetRowOrderByDefinition{
	CreatedAt: "created_at",
	Email:     "email",
	Token:     "token",
}

func PasswordResetRowColumnList() []string {
	result := []string{
		"created_at",
		"email",
		"token",
	}
	return result
}
