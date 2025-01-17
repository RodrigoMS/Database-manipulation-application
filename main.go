package main

import (
	"net/http"

	"github.com/RodrigoMS/app/database"
)

func main() {
	database.Connection()

	router := routes()

	http.ListenAndServe(":8080", router)
}
