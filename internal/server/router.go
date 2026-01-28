package server

import (
	"net/http"

	"github.com/RodrigoMS/app/cmd/internal/handlers"
	"github.com/RodrigoMS/app/cmd/internal/views"
)

func routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /users", handlers.GetUsers)
	router.HandleFunc("POST /users", handlers.PostUser)
	router.HandleFunc("PATCH /users", handlers.PutUser)
	router.HandleFunc("DELETE /users", handlers.DeleteUser)

	router.HandleFunc("GET /database-info", handlers.GetDatabaseInfo)

	router.HandleFunc("/", views.HandleNotFound)

	return router
}