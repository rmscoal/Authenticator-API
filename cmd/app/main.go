package main

import (
	"log"

	"github.com/rmscoal/Authenticator-API/config"
	"github.com/rmscoal/Authenticator-API/internal/app"
)

func main() {
	// Configurations
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error %s", err)
	}

	// Run
	app.Run(cfg)
}
