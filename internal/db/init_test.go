package db

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestInitDatabase(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		t.Fatalf("не удалось загрузить .env.test: %v", err)
	}
	InitDatabase()

	if Database == nil {
		t.Fatal("db всё еще nil после initDatabase()")
	}

	if err := Database.Ping(); err != nil {
		t.Fatalf("Пинг не прошел после initDatabase(): %v", err)
	}
}
