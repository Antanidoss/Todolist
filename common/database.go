package common

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"

)

var DB *sql.DB

func Init() *sql.DB {
	envErr := godotenv.Load()
	if envErr != nil {
		panic("Error loading credentials: %v" + envErr.Error())
	}

	// var (
	// 	user     = os.Getenv("MSSQL_DB_USER")
	// 	database = os.Getenv("MSSQL_DB_DATABASE")
	// )

	// connectionString := fmt.Sprintf("user id=%s;database=%s", user, database)

	db, err := sql.Open("mssql", "./CreateDB.sql")

	if err != nil {
		panic("db err: (Init) " + err.Error())
	}

	DB = db

	//createDb()

	return DB
}

func createDb() {
	path := filepath.Join("scripts", "CreateDB.sql")

	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	sql := string(buffer)
	DB.Exec(sql)
}
