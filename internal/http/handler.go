package http

import "github.com/sbaitmangalkar/tiny-ticket-and-message-service/internal/message"

func (s *Server) registerHandlers() {
	messageRouter := s.router.PathPrefix("/v1/message").Subrouter()
	message.RegisterHandlers(messageRouter)
}
