package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "GoTicket"
)

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %s\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s\n", err)
	}

	fmt.Println("Successfully connected to the database")
}

func ExpireSeats() {
	query := "UPDATE seats SET status = 1 WHERE status = 2 AND updated_at < NOW() - INTERVAL '20 minutes'"
	_, err := DB.Exec(query)
	if err != nil {
		log.Printf("Error updating expired seats: %v\n", err)
	} else {
		log.Println("Expired seats updated successfully.")
	}
}
