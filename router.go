package main

import (
	"net/http"

	"github.com/RodrigoMS/app/cmd/controllers"
)

func routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /user/", controllers.GetUser)
	router.HandleFunc("POST /user", controllers.PostUser)
	router.HandleFunc("PUT /user", controllers.PutUser)
	router.HandleFunc("DELETE /user", controllers.DeleteUser)

	router.HandleFunc("GET /users", controllers.GetUsers)
	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	return router
}
