package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Conn is database connection
var Conn *sql.DB

func init() {
	conn, err := sql.Open("sqlite3", "/Users/ryo/go/src/vue-golang-payment-app/backend-api/shop.db")
	if err != nil {
		panic(err.Error)
	}

	Conn = conn
}
