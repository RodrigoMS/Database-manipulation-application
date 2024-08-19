package main

import (
	"net/http"

	"github.com/RodrigoMS/app/cmd/controllers"
)

func routes() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/users", usersHandler)
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	//var routeMap = map[string]func(http.ResponseWriter, *http.Request){

	var userHandlers = map[string]func(http.ResponseWriter, *http.Request){
		"GET":    controllers.GetUser,
		"POST":   controllers.PostUser,
		"PUT":    controllers.PutUser,
		"DELETE": controllers.DeleteUser,
		//"PATCH":  func() { models.GetUser() },*/
	}

	if handler, ok := userHandlers[r.Method]; ok {
		handler(w, r)

	} else {
		// Status 405 - Método não suportado.
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	controllers.GetUsers(w, r)
}

