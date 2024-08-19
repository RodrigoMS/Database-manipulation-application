package main

import (
	"net/http"

	"github.com/RodrigoMS/app/database"
)

func main() {
	database.Connection()

	routes()

	http.ListenAndServe(":8080", nil)
}
