package server

import (
	"net/http"

	"github.com/RodrigoMS/app/cmd/internal/database"
)

func Start() {
	database.ConnectionMonitor()
  	defer database.CloseConnection()

	router := routes()

	http.ListenAndServe(":8080", router)
}

/*import (
    "log"
    "net/http"
    "time"
)

type Server struct {
    addr string
}

func NewServer(addr string) *Server {
    return &Server{addr: addr}
}

func (s *Server) Start() error {
    router := NewRouter()
    
    server := &http.Server{
        Addr:         s.addr,
        Handler:      router,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }
    
    log.Printf("Server starting on %s", s.addr)
    return server.ListenAndServe()
}*/