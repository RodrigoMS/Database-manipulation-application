package server

import (
	"net/http"
	"sync"

	"github.com/RodrigoMS/app/cmd/internal/handlers"
	"github.com/RodrigoMS/app/cmd/internal/web"
)

var (
	userHandlers = map[string]func(http.ResponseWriter, *http.Request) {
		"GET":    handler.GetUser,
		"POST":   handler.PostUser,
		"PUT":    handler.PutUser,
		"DELETE": handler.DeleteUser,
		//"PATCH":  func() { models.GetUser() },*/
	}

	mutex sync.RWMutex
)

func routes() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user/{id}", userHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/server", web.InformationDatabase)
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
	handler.GetUsers(w, r)
}

