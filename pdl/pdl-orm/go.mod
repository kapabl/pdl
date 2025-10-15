module github.com/kapablanka/pdl/pdl-orm

go 1.21

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/lib/pq v1.10.9
)

require github.com/kapablanka/pdl/pdl/infra/go v0.0.0

replace github.com/kapablanka/pdl/pdl/infra/go => ../infra/go
