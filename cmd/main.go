package main

import (
	"log"

	"github.com/keshvan/user-service-forum-go/config"
	"github.com/keshvan/user-service-forum-go/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
