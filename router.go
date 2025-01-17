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

func routes() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/users", usersHandler)
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	mutex.RLock()
	defer mutex.RUnlock()

	handler, ok := userHandlers[r.Method];

	if ok {
		handler(w, r)

	} else {
		// Status 405 - Método não suportado.
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	controllers.GetUsers(w, r)
}

