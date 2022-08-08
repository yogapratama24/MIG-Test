package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

var (
	db  *sql.DB
	err error
)

func Connect() *sql.DB {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err = sql.Open("mysql", uri)
	if err != nil {
		log.Fatalf("Can't connect to database with err: %s", err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	// if err := goose.Down(db, "migrations"); err != nil {
	// 	panic(err)
	// }

	// ---RESET TABLE---
	// if err := goose.Reset(db, "migrations"); err != nil {
	// 	panic(err)
	// }

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	return db
}
