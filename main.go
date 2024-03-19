package main

import (
	"fmt"
	"net/http"

	"github.com/mananispiwpiw/fp-gdsc-binus/api"
	"github.com/mananispiwpiw/fp-gdsc-binus/db"
)

func main() {
	db.ConnectDB()

	api.NewRouter()

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
