package server

import (
	"net/http"

	"github.com/RodrigoMS/app/cmd/internal/handlers"
	//"github.com/RodrigoMS/app/cmd/internal/views"
	//"github.com/RodrigoMS/app/cmd/internal/web"
)

func routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /users", handlers.GetUsers)
	router.HandleFunc("POST /users", handlers.PostUser)
	router.HandleFunc("PATCH /users", handlers.PutUser)
	router.HandleFunc("DELETE /users", handlers.DeleteUser)

	router.HandleFunc("GET /student-login/{class}", handlers.StudentLogin)
	router.HandleFunc("POST /student-authentication", handlers.StudentAuthentication)
	router.HandleFunc("POST /student-logout/{class}", handlers.LogoutHandler)

	router.HandleFunc("GET /student-dashboard", handlers.StudentDashboard)

	router.HandleFunc("GET /database-info", handlers.GetDatabaseInfo)

	router.HandleFunc("/", handlers.GetDatabaseInfo)

	//router.HandleFunc("/", views.HandleNotFound)

	//web.SetupEmbeddedStatic(router)

	return router
}

