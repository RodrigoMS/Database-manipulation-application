package main

import (
	"net/http"

	"github.com/RodrigoMS/app/database"
)

func main() {
	database.Connection()

	Routes()

	http.ListenAndServe(":8080", nil)
}
