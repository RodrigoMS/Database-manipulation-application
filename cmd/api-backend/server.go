package server

import (
	"net/http"
)

func Start() {

	router := routes()

	http.ListenAndServe(":8080", router)
}
