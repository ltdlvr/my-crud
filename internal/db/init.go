package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)

	var err error
	Database, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка соединения с БД", err)
	}

	if err := Database.Ping(); err != nil {
		log.Fatal("Ошибка пинга БД", err)
	} else {
		fmt.Println("УРААААААА С ПОДКЛЮЧЕНИЕМ!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	}
}
