// File: db/db_test.go

package db

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	InitDB()
	defer DB.Close()
	err := DB.Ping()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %s\n", err)
	}

	t.Log("Connected to the database successfully.")
}
