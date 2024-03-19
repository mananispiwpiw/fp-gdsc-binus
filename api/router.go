package api

import (
	"net/http"
)

func NewRouter() {
	//Handler thingy
	http.HandleFunc("/tasks", GetTasksHandler)
	http.HandleFunc("/task", AddTaskHandler)
	http.HandleFunc("DELETE /task/{id}", DeleteTaskHandler)
	http.HandleFunc("PUT /task/{id}", UpdateTaskHandler)
}
