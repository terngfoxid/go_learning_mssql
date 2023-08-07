package Config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

var server = "127.0.0.1"
var port = 1433
var user = "admin"
var password = "12345678"
var database = "test_go_ms_sql"

func GetConStr() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
}

func StartConnection() {
	// Build connection string
	connString := GetConStr()

	var err error

	// Create connection pool
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = DB.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
}
