package main

import (
	"net/http"

	"github.com/RodrigoMS/app/internal/database"
)

func main() {

	database.ConnectionMonitor()
    defer database.CloseConnection()

	// Aguarda conexão com o banco e executa lógica dependente
    /*go func() {
        <-database.ConnectedChan
        database.GetDB().GetDBInfo()
    }()*/

	routes()

	http.ListenAndServe(":8080", nil)
}
