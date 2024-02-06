package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitConnection(connection string) {
	// Wait for railway private network to initialize...
	time.Sleep(3 * time.Second)

	// Open a connection to the database
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
}
