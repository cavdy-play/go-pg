package db

import (
	"log"
	"os"
	"github.com/go-pg/pg/v9"
)

// Connecting to db
func Connect() {
	opts := &pg.Options{
		User: "cavdy",
		Password: "postgres",
		Addr: "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	closeError := db.Close();

	if closeError != nil {
		log.Printf("Error while closing connection Error: %v\n", closeError)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
	return
}
