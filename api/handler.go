package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/mananispiwpiw/fp-gdsc-binus/db"
)

// Handler for GetTasks
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	//Check if the request method is GET
	if r.Method == "GET" {
		//Get the tasks from the database
		tasks, err := db.GetTasks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Write the tasks to the response
		w.Write(tasks)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Handler for AddTask
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	//Chcek if the request method is not POST
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse and decode the request body
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the task to the database
	err = db.AddTask(task.ID, task.Title, task.Description, time.Now(), time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// Handler for delete a task
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	//Check if the request method is not DELETE
	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	//Change to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
	}

	// Delete the task from the database
	err = db.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Task deleted successfully!")
}
