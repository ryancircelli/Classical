package db

import (
	"database/sql"
	f "fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

// GetPostgresDB - get postgres db config connection
func Connect() {
	var err error

	DB, err = sql.Open("mysql", "root:password123@tcp(localhost:3306)/classical")
	if err != nil {
		panic(err)
	}

	err = PingDB()
	if err != nil {
		panic(err)
	}

	f.Println("CONNECTED")

	DB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(2)
}

func PingDB() error {
	err := DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
