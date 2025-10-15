module github.com/kapablanka/pdl/sample/tests

go 1.21

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/kapablanka/pdl/pdl/infra/go v0.0.0
	github.com/kapablanka/pdl/sample v0.0.0
)

replace github.com/kapablanka/pdl/pdl/infra/go => ../pdl/infra/go

replace github.com/kapablanka/pdl/sample => ../pdl-project/output/db2pdl/go
