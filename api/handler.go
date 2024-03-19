package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mananispiwpiw/fp-gdsc-binus/db"
)

// GetUsers function
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//Check if the request method is GET
	if r.Method == "GET" {
		//Get the users from the database
		users, err := db.GetUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Write the users to the response
		w.Write(users)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// AddUser funcction
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	//Chcek if the request method is not POST
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse and decode the request body
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the user to the database
	err = db.AddUser(user.ID, user.Username, user.Password, time.Now(), time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
