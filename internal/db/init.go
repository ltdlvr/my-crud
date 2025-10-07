package db

import (
	"database/sql"
	"fmt"
	"log"
	"my-crud/config"

	_ "github.com/lib/pq"
)

func InitDatabase(cfg *config.Config) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	var err error
	Database, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	if err := Database.Ping(); err != nil {
		log.Fatal("Failed to ping database", err)
	} else {
		fmt.Println("Database connection established successfully!")
	}
}
