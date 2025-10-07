package db

import (
	"testing"

	"my-crud/config"

	"github.com/joho/godotenv"
)

func TestInitDatabase(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		t.Fatalf("Failed to load .env.test: %v", err)
	}
	cfg := config.LoadConfig()
	InitDatabase(cfg)

	if Database == nil {
		t.Fatal("Db is still nil after initDatabase()")
	}

	if err := Database.Ping(); err != nil {
		t.Fatalf("Ping failed after initDatabase(): %v", err)
	}
}
