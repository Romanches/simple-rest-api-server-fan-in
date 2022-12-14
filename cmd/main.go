package main

import (
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/handlers"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"log"
)

func main() {
	// Load config
	serverCfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	healthHandler := handlers.NewHealthHandler()

	apiV1 := api.NewAPI(
		serverCfg,
		healthHandler,
	)
	apiV1.Run()
}

// Reads config from file
func loadConfig() (models.Config, error) {

	serverCfg := models.Config{
		ListenAddr: ":8000",
	}

	serverCfg.Validate()

	// Success
	return serverCfg, nil
}

