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

	var (
		server = os.Getenv("MSSQL_DB_SERVER")
		dbName = os.Getenv("MSSQL_DB_NAME")
	)

	dbConnectionString := fmt.Sprintf("odbc:server=%s;database=%s", server, dbName)
	db, isCreated := checkDbCreated(dbConnectionString)

	if isCreated {
		DB = db
		return DB
	}

	sqlServerConnectionString := fmt.Sprintf("odbc:server=%s", server)
	db, err := sql.Open(providerName, sqlServerConnectionString)

	if err != nil {
		panic("db err: (Init) %v" + err.Error())
	}

	createDb(db, dbName)

	DB, err := sql.Open(providerName, dbConnectionString)

	if err != nil {
		panic("db err: (Init) %v" + err.Error())
	}

	createDbTables(DB)

	return DB
}

func checkDbCreated(dbConnectionString string) (*sql.DB, bool) {
	db, err := sql.Open(providerName, dbConnectionString)

	if err != nil {
		panic("db err: (Init) %v" + err.Error())
	}

	err = db.Ping()

	if err == nil {
		return db, true
	}

	return nil, false
}

func createDb(DB *sql.DB, dbName string) {
	sql := fmt.Sprintf("CREATE DATABASE %s", dbName)
	DB.Exec(sql)
}

func createDbTables(DB *sql.DB) {
	path := filepath.Join("common", "CreateDB.sql")

	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	sql := string(buffer)
	DB.Exec(sql)
}
