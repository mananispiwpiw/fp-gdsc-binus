package api

import (
	"net/http"

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
