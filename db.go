package main

import (
	"database/sql"
)

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}

func CreateUser(db *sql.DB, name string, age int) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", name, age).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func UpdateUserAge(db *sql.DB, id int, newAge int) error {
	_, err := db.Exec("UPDATE users SET age = $1 WHERE id = $2", newAge, id)
	return err
}

func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
