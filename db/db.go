package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// User strucy
type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt string
	UpdatedAt string
}

var db *sql.DB

func ConnectDB() {
	// Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres123", "fp_gdsc")

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
}

// Users thingy
func GetUsers() ([]byte, error) {
	// Query the database
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []User

	// Iterate over the rows
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData, nil
}

func AddUser(id int, username string, password string, createdAt time.Time, updatedAt time.Time) error {
	// Query the database
	_, err := db.Exec("INSERT INTO users (id,username,password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)", id, username, password, createdAt, updatedAt)
	if err != nil {
		return err
	}

	fmt.Println("User added successfully!")
	return err
}
