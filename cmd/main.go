package main

import "github.com/RodrigoMS/app/cmd/internal/server"

func main() {

	server.Start()

	//database.ConnectionMonitor()
	//  defer database.CloseConnection()

	// Aguarda conexão com o banco e executa lógica dependente
	/*go func() {
	    <-database.ConnectedChan
	    database.GetDB().GetDBInfo()
	}()*/

	//routes()

	//http.ListenAndServe(":8080", nil)
}
