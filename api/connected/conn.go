package connected

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var Connectionstring string

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("error .env file "))
	}

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	Connectionstring = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASS, HOST, PORT, DBNAME)
	DB, err = sql.Open("mysql", Connectionstring)

	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}
}
