package db2pdl

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/mysql"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/postgres"
	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl/shared"
)

type metadataDriver interface {
	Open(ctx context.Context) (*sql.DB, error)
	ListTables(ctx context.Context, connection *sql.DB) ([]string, error)
	BuildTableData(ctx context.Context, connection *sql.DB, tableName string, baseName string) (shared.TableData, error)
}

func createMetadataDriver(databaseType string, config shared.DB2PDLConfig, defaultNamespace string, phpNamespace string) (metadataDriver, error) {
	switch strings.ToLower(databaseType) {
	case "mysql":
		return mysql.NewDriver(config, defaultNamespace, phpNamespace), nil
	case "postgres", "postgresql":
		return postgres.NewDriver(config, defaultNamespace, phpNamespace), nil
	default:
		return nil, fmt.Errorf("unsupported database type %s", databaseType)
	}
}
