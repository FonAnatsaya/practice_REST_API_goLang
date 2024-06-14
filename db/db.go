package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:anatsaya@tcp(localhost:3306)/new_schema")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = CreateEventsTable()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func CreateEventsTable() error {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(255) NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INT
		)
		`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		return fmt.Errorf("could not create events table: %w", err)
	}

	return nil
}
