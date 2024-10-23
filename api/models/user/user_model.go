package UserModel

import (
	"database/sql"
	"fmt"
	"log"
	"ultimate_backend/api/database"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type User struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateUserTable() error {
	db = database.GetDB()

	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			ID SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
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
	// hashedPassword, err := hashPassword(user.Password)
	// if err != nil {
	// 	return err
	// }

	if _, err := db.Exec("INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4)", user.Email, user.Password, user.FirstName, user.LastName); err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]User, error) {
	var users []User = []User{}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(email string) (User, error) {
	var user User

	log.Println("111")
	rows, err := db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName); err != nil {
			return User{}, err
		}
	}

	return user, nil
}

// func hashPassword(password string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(hash), nil
// }

func CheckPassword(hashedPassword []byte, password []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return err
	}

	return nil
}
