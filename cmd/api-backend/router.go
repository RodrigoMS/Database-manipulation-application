package server

import (
	"net/http"

	"github.com/RodrigoMS/App_Go/cmd/api-backend/controllers"
	"github.com/RodrigoMS/App_Go/cmd/api-backend/views"
)

func routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /users", controllers.GetUsers)
	router.HandleFunc("POST /users", controllers.PostUser)
	router.HandleFunc("PATCH /users", controllers.PutUser)
	router.HandleFunc("DELETE /users", controllers.DeleteUser)

	router.HandleFunc("GET /products", controllers.GetProducts)
	router.HandleFunc("POST /products", controllers.PostProduct)
	router.HandleFunc("PATCH /products", controllers.PutProduct)
	router.HandleFunc("DELETE /products", controllers.DeleteProduct)

	router.HandleFunc("GET /database-info", controllers.GetDatabaseInfo)

	router.HandleFunc("/", views.HandleNotFound)

	return router
}
