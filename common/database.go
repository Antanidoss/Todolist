package common

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var DB *sql.DB

const providerName = "sqlserver"

func Init() *sql.DB {
	envErr := godotenv.Load()
	if envErr != nil {
		panic("Error loading credentials: %v" + envErr.Error())
	}

	server := os.Getenv("MSSQL_DB_SERVER")
	sqlServerConnectionString := fmt.Sprintf("odbc:server=%s", server)

	db, err := sql.Open(providerName, sqlServerConnectionString)

	if err != nil {
		panic("db err: (Init) " + err.Error())
	}

	createDb(db)

	dbName := os.Getenv("MSSQL_DB_NAME")
	dbConnectionString := fmt.Sprintf("odbc:server=%s;database=%s", server, dbName)
	DB, err := sql.Open(providerName, dbConnectionString)

	if err != nil {
		panic("db err: (Init) " + err.Error())
	}

	return DB
}

func createDb(DB *sql.DB) {
	path := filepath.Join("common", "CreateDB.sql")

	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	sql := string(buffer)
	DB.Exec(sql)
}
