package api

import (
	"net/http"
)

func NewRouter() {
	//Handler thingy
	http.HandleFunc("/users", GetUsersHandler)
	http.HandleFunc("/user", AddUserHandler)
}
