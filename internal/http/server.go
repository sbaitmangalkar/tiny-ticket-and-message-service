package http

import (
	"github.com/gorilla/mux"
	"github.com/sbaitmangalkar/tiny-ticket-and-message-service/config"
	"net/http"
)

// Server struct
type Server struct {
	config *config.Config
	router *mux.Router
}

// NewServer Server constructor
func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *Server) ListenAndServe(port string) {
	s.registerHandlers()
	http.Handle("/", s.router)
	http.ListenAndServe(":"+port, s.router)
}
