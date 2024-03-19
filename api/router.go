package api

import (
	"net/http"
)

func NewRouter() {
	//Handler thingy
	http.HandleFunc("/tasks", GetTasksHandler)
	http.HandleFunc("/task", AddTaskHandler)
}
