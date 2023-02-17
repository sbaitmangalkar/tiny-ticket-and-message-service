package http

import (
	"github.com/sbaitmangalkar/tiny-ticket-and-message-service/internal/message"
	"github.com/sbaitmangalkar/tiny-ticket-and-message-service/internal/ticket"
)

func (s *Server) registerHandlers() {
	messageRouter := s.router.PathPrefix("/v1/message").Subrouter()
	message.RegisterHandlers(messageRouter)
	ticketRouter := s.router.PathPrefix("/v1/ticket").Subrouter()
	ticket.RegisterHandlers(ticketRouter)
}
