package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

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
