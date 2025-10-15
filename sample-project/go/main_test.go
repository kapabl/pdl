package sample

import (
	"bufio"
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	pdlgo "github.com/kapablanka/pdl/pdl/infra/go"
)

var (
	testDB    *sql.DB
	testStore *pdlgo.MySQLDB
)

func TestMain(m *testing.M) {
	env := loadTestingEnv()
	driver := strings.ToLower(env["DB_CONNECTION"])
	if driver == "" {
		driver = "mysql"
	}

	dsnKey := strings.ToUpper(driver) + "_TEST_DSN"
	dsn := env[dsnKey]

	switch driver {
	case "mysql":
		if dsn != "" {
			db, err := sql.Open("mysql", dsn)
			if err != nil {
				log.Fatalf("failed opening %s connection: %v", dsnKey, err)
			}
			if err := db.Ping(); err != nil {
				log.Fatalf("failed pinging %s: %v", dsnKey, err)
			}
			testDB = db
			testStore = pdlgo.NewMySQLDB(db)
			pdlgo.SetDefaultStore(testStore)
		}
	default:
		if dsn != "" {
			log.Printf("DB_CONNECTION=%s not yet supported in Go tests; skipping DB integration", driver)
		}
	}

	code := m.Run()

	if testDB != nil {
		testDB.Close()
		pdlgo.SetDefaultStore(nil)
	}

	os.Exit(code)
}

func loadTestingEnv() map[string]string {
	envPath := filepath.Join(".", ".env.test")
	file, err := os.Open(envPath)
	if err != nil {
		return map[string]string{}
	}
	defer file.Close()

	result := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		idx := strings.Index(line, "=")
		if idx <= 0 {
			continue
		}
		key := strings.TrimSpace(line[:idx])
		value := strings.TrimSpace(line[idx+1:])
		result[key] = value
	}
	return result
}

func tableLookupExists(table string) bool {
	if testDB == nil {
		return false
	}
	query := "SELECT 1 FROM `" + table + "` LIMIT 1"
	_, err := testDB.Exec(query)
	if err == nil {
		return true
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1146 {
		return false
	}
	log.Fatalf("unexpected error probing table %s: %v", table, err)
	return false
}

func mustExec(query string, args ...any) {
	if testDB == nil {
		return
	}
	if _, err := testDB.Exec(query, args...); err != nil {
		log.Fatalf("failed executing %s: %v", query, err)
	}
}
