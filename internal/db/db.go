package db

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() error {

	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open a DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`)

	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	return nil
}
