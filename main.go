package main

import (
	server "github.com/RodrigoMS/App_Go/cmd/api-backend"
	"github.com/RodrigoMS/App_Go/internal/database"
)

func main() {
	database.ConnectionMonitor()
	defer database.CloseConnection()

	server.Start()
}
