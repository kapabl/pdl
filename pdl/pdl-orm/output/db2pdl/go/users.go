package sample

import (
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
	"time"
)

type UserRowRecord struct {
	Row                    *pdlgo.Row `pdl:"-"`
	CanSell                string     `pdl:"can_sell"`
	CashAppId              string     `pdl:"cash_app_id"`
	CreatedAt              time.Time  `pdl:"created_at"`
	CurrentTeamId          int64      `pdl:"current_team_id"`
	DatetimeZone           string     `pdl:"datetime_zone"`
	DeliveryAddressId      int        `pdl:"delivery_address_id"`
	Email                  string     `pdl:"email"`
	EmailVerifiedAt        time.Time  `pdl:"email_verified_at"`
	FacebookId             string     `pdl:"facebook_id"`
	GoogleId               string     `pdl:"google_id"`
	Id                     int64      `pdl:"id"`
	IsAvailable            string     `pdl:"is_available"`
	IsSystemAdmin          string     `pdl:"is_system_admin"`
	IsTestUser             string     `pdl:"is_test_user"`
	LastLocation           string     `pdl:"last_location"`
	LastLocationTime       time.Time  `pdl:"last_location_time"`
	Locale                 string     `pdl:"locale"`
	Name                   string     `pdl:"name"`
	Password               string     `pdl:"password"`
	Phone                  string     `pdl:"phone"`
	ProfilePhotoPath       string     `pdl:"profile_photo_path"`
	RememberToken          string     `pdl:"remember_token"`
	Role                   string     `pdl:"role"`
	Status                 string     `pdl:"status"`
	StoreId                int        `pdl:"store_id"`
	TestSeller             string     `pdl:"test_seller"`
	TwoFactorRecoveryCodes string     `pdl:"two_factor_recovery_codes"`
	TwoFactorSecret        string     `pdl:"two_factor_secret"`
	UpdatedAt              time.Time  `pdl:"updated_at"`
	Uuid                   string     `pdl:"uuid"`
	ZelleEmail             string     `pdl:"zelle_email"`
	ZellePhone             string     `pdl:"zelle_phone"`
}

type UserRowFactory struct{}

var UserRow = UserRowFactory{}

func (factory UserRowFactory) New() *UserRowRecord {
	result := &UserRowRecord{
		Row: pdlgo.NewRow("users", "id"),
	}
	return result
}

func (factory UserRowFactory) WithStore(store pdlgo.DBStore) *UserRowRecord {
	result := factory.New()
	result.Row.SetStore(store)
	return result
}

func (record *UserRowRecord) Create() error {
	result := pdlgo.Create(record)
	return result
}

func (record *UserRowRecord) Update() error {
	result := pdlgo.Update(record)
	return result
}

func (record *UserRowRecord) Delete() error {
	result := pdlgo.Delete(record)
	return result
}

type UserRowWhereBuilder struct {
	pdlgo.QueryBuilder
}

func UserRowWhere() UserRowWhereBuilder {
	result := UserRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("users", nil)}
	return result
}

func UserRowWhereWithStore(store pdlgo.DBStore) UserRowWhereBuilder {
	result := UserRowWhereBuilder{QueryBuilder: pdlgo.NewQueryBuilder("users", store)}
	return result
}

func (builder UserRowWhereBuilder) CanSell(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("can_sell", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) CashAppId(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("cash_app_id", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) CreatedAt(value time.Time) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("created_at", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) CurrentTeamId(value int64) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("current_team_id", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) DatetimeZone(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("datetime_zone", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) DeliveryAddressId(value int) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("delivery_address_id", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Email(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("email", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) EmailVerifiedAt(value time.Time) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("email_verified_at", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) FacebookId(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("facebook_id", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) GoogleId(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("google_id", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Id(value int64) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("id", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) IsAvailable(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_available", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) IsSystemAdmin(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_system_admin", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) IsTestUser(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("is_test_user", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) LastLocation(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("last_location", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) LastLocationTime(value time.Time) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("last_location_time", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Locale(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("locale", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Name(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("name", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Password(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("password", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Phone(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("phone", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) ProfilePhotoPath(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("profile_photo_path", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) RememberToken(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("remember_token", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Role(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("role", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Status(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("status", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) StoreId(value int) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("store_id", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) TestSeller(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("test_seller", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) TwoFactorRecoveryCodes(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("two_factor_recovery_codes", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) TwoFactorSecret(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("two_factor_secret", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) UpdatedAt(value time.Time) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("updated_at", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) Uuid(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("uuid", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) ZelleEmail(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("zelle_email", pdlgo.OpEq, value)
	return builder
}
func (builder UserRowWhereBuilder) ZellePhone(value string) UserRowWhereBuilder {
	builder.QueryBuilder = builder.QueryBuilder.Filter("zelle_phone", pdlgo.OpEq, value)
	return builder
}

func (builder UserRowWhereBuilder) Load() ([]map[string]any, error) {
	result, loadError := builder.QueryBuilder.Load()
	return result, loadError
}

func (builder UserRowWhereBuilder) Delete() error {
	result := builder.QueryBuilder.Delete("id")
	return result
}

type UserRowColumnsDefinition struct {
	CanSell                string
	CashAppId              string
	CreatedAt              string
	CurrentTeamId          string
	DatetimeZone           string
	DeliveryAddressId      string
	Email                  string
	EmailVerifiedAt        string
	FacebookId             string
	GoogleId               string
	Id                     string
	IsAvailable            string
	IsSystemAdmin          string
	IsTestUser             string
	LastLocation           string
	LastLocationTime       string
	Locale                 string
	Name                   string
	Password               string
	Phone                  string
	ProfilePhotoPath       string
	RememberToken          string
	Role                   string
	Status                 string
	StoreId                string
	TestSeller             string
	TwoFactorRecoveryCodes string
	TwoFactorSecret        string
	UpdatedAt              string
	Uuid                   string
	ZelleEmail             string
	ZellePhone             string
}

type UserRowOrderByDefinition struct {
	CanSell                string
	CashAppId              string
	CreatedAt              string
	CurrentTeamId          string
	DatetimeZone           string
	DeliveryAddressId      string
	Email                  string
	EmailVerifiedAt        string
	FacebookId             string
	GoogleId               string
	Id                     string
	IsAvailable            string
	IsSystemAdmin          string
	IsTestUser             string
	LastLocation           string
	LastLocationTime       string
	Locale                 string
	Name                   string
	Password               string
	Phone                  string
	ProfilePhotoPath       string
	RememberToken          string
	Role                   string
	Status                 string
	StoreId                string
	TestSeller             string
	TwoFactorRecoveryCodes string
	TwoFactorSecret        string
	UpdatedAt              string
	Uuid                   string
	ZelleEmail             string
	ZellePhone             string
}

var UserRowColumns = UserRowColumnsDefinition{
	CanSell:                "can_sell",
	CashAppId:              "cash_app_id",
	CreatedAt:              "created_at",
	CurrentTeamId:          "current_team_id",
	DatetimeZone:           "datetime_zone",
	DeliveryAddressId:      "delivery_address_id",
	Email:                  "email",
	EmailVerifiedAt:        "email_verified_at",
	FacebookId:             "facebook_id",
	GoogleId:               "google_id",
	Id:                     "id",
	IsAvailable:            "is_available",
	IsSystemAdmin:          "is_system_admin",
	IsTestUser:             "is_test_user",
	LastLocation:           "last_location",
	LastLocationTime:       "last_location_time",
	Locale:                 "locale",
	Name:                   "name",
	Password:               "password",
	Phone:                  "phone",
	ProfilePhotoPath:       "profile_photo_path",
	RememberToken:          "remember_token",
	Role:                   "role",
	Status:                 "status",
	StoreId:                "store_id",
	TestSeller:             "test_seller",
	TwoFactorRecoveryCodes: "two_factor_recovery_codes",
	TwoFactorSecret:        "two_factor_secret",
	UpdatedAt:              "updated_at",
	Uuid:                   "uuid",
	ZelleEmail:             "zelle_email",
	ZellePhone:             "zelle_phone",
}

var UserRowOrderBy = UserRowOrderByDefinition{
	CanSell:                "can_sell",
	CashAppId:              "cash_app_id",
	CreatedAt:              "created_at",
	CurrentTeamId:          "current_team_id",
	DatetimeZone:           "datetime_zone",
	DeliveryAddressId:      "delivery_address_id",
	Email:                  "email",
	EmailVerifiedAt:        "email_verified_at",
	FacebookId:             "facebook_id",
	GoogleId:               "google_id",
	Id:                     "id",
	IsAvailable:            "is_available",
	IsSystemAdmin:          "is_system_admin",
	IsTestUser:             "is_test_user",
	LastLocation:           "last_location",
	LastLocationTime:       "last_location_time",
	Locale:                 "locale",
	Name:                   "name",
	Password:               "password",
	Phone:                  "phone",
	ProfilePhotoPath:       "profile_photo_path",
	RememberToken:          "remember_token",
	Role:                   "role",
	Status:                 "status",
	StoreId:                "store_id",
	TestSeller:             "test_seller",
	TwoFactorRecoveryCodes: "two_factor_recovery_codes",
	TwoFactorSecret:        "two_factor_secret",
	UpdatedAt:              "updated_at",
	Uuid:                   "uuid",
	ZelleEmail:             "zelle_email",
	ZellePhone:             "zelle_phone",
}

func UserRowColumnList() []string {
	result := []string{
		"can_sell",
		"cash_app_id",
		"created_at",
		"current_team_id",
		"datetime_zone",
		"delivery_address_id",
		"email",
		"email_verified_at",
		"facebook_id",
		"google_id",
		"id",
		"is_available",
		"is_system_admin",
		"is_test_user",
		"last_location",
		"last_location_time",
		"locale",
		"name",
		"password",
		"phone",
		"profile_photo_path",
		"remember_token",
		"role",
		"status",
		"store_id",
		"test_seller",
		"two_factor_recovery_codes",
		"two_factor_secret",
		"updated_at",
		"uuid",
		"zelle_email",
		"zelle_phone",
	}
	return result
}
