package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kskr24/workspacehub/internal/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg *config.Config) error {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	var err error
	DB, err = sql.Open("postgres", connStr)
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
