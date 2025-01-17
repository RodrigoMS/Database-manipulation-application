package main

import (
	"net/http"
	"sync"

	"github.com/RodrigoMS/app/cmd/controllers"
)

var (
	userHandlers = map[string]func(http.ResponseWriter, *http.Request) {
		"GET":    controllers.GetUser,
		"POST":   controllers.PostUser,
		"PUT":    controllers.PutUser,
		"DELETE": controllers.DeleteUser,
		//"PATCH":  func() { models.GetUser() },*/
	}

	mutex sync.RWMutex
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
