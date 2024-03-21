package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/mananispiwpiw/fp-gdsc-binus/model"
)

var Task []model.Task

var db *sql.DB

func ConnectDB() {
	// Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"ceu9lmqblp8t3q.cluster-czrs8kj4isg7.us-east-1.rds.amazonaws.com", 5432, "u8o4gqqc9mml7c", "pcde730fa9d14a91b62c6c1db6424d7cace30556ce67e70e614cd078de2134fdd", "d380bv5hepsdcd")

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

	var tasks []model.Task

	// Iterate over the rows
	for rows.Next() {
		var task model.Task
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
func DeleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	fmt.Println("Task deleted sucessfully!")
	return nil
}

// UpdateTask function
func UpdateTask(id int, title string, description string, updatedAt time.Time) error {
	_, err := db.Exec("UPDATE tasks SET title = $1, description = $2, updated_at = $3 WHERE id = $4", title, description, updatedAt, id)
	if err != nil {
		return err
	}

	fmt.Println("Task updated successfully!")
	return nil
}
