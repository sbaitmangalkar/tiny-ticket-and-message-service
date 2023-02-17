package main

import (
	"github.com/sbaitmangalkar/tiny-ticket-and-message-service/config"
	"github.com/sbaitmangalkar/tiny-ticket-and-message-service/internal/http"
	"log"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	log.Println("starting API server")
	log.Printf("AppVersion: %s, Host: %s, Port: %s, Mode: %s\n", cfg.Server.AppVersion, cfg.Server.Host, cfg.Server.Port, cfg.Server.Mode)
	s := http.NewServer(&cfg)
	s.ListenAndServe(cfg.Server.Port)
}
