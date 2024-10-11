package UserModel

import (
	"database/sql"
	"fmt"
	"ultimate_backend/api/database"
)

var db *sql.DB

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateUserTable() error {
	db = database.GetDB()

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			ID SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL
		)
	`); err != nil {
		return err
	}

	fmt.Println("Table user created successfully")
	return nil
}

func InsertUser(user User) error {
	_, err := db.Exec("INSERT INTO users (username, password, first_name, last_name) VALUES ($1, $2, $3, $4)", user.Username, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
