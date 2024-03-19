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
type Task struct {
	ID          int
	Title       string
	Description string
	CreatedAt   string
	UpdatedAt   string
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

// Get all tasks function
func GetTasks() ([]byte, error) {
	// Query the database
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var tasks []Task

	// Iterate over the rows
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData, nil
}

// AddTask function
func AddTask(id int, title string, description string, createdAt time.Time, updatedAt time.Time) error {
	// Query the database
	_, err := db.Exec("INSERT INTO tasks (id,title,description,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)", id, title, description, createdAt, updatedAt)
	if err != nil {
		return err
	}

	fmt.Println("New task added successfully!")
	return err
}

// DeleteTask function
