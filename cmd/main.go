package main

import (
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/handlers"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/helpers/rest"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	rData "github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/repository/data"
	sData "github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/services/data"
	"log"
)

func main() {
	// Load config
	serverCfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	/* Our application consists of three layers:
		1. handlers - transport layer
		2. service - business logic layer
		3. repository - data source (db or external resources)

	These three layers help us to keep code clean and avoid circle imports/dependencies. */

	// Handler is all what we need for health-check route
	healthHandler := handlers.NewHealthHandler()

	// It should be DB connector, but in our application we need http-client to get data from third-party resources
	httpClient := rest.NewHttpClient()

	// Repository instance for /data endpoint
	dataRepository := rData.NewRepository(httpClient)

	// Service instance for /data endpoint
	dataService := sData.NewDataService(serverCfg.Resources, dataRepository)

	// Handler instance for /data endpoint
	dataHandler := handlers.NewDataHandler(dataService)

	server := api.NewAPI(
		serverCfg,
		healthHandler,
		dataHandler,
	)
	server.Run()

}

// Builds config
func loadConfig() (models.Config, error) {

	resources := []string{
		"https://raw.githubusercontent.com/assignment132/assignment/main/duckduckgo.json",
		"https://raw.githubusercontent.com/assignment132/assignment/main/google.json",
		"https://raw.githubusercontent.com/assignment132/assignment/main/wikipedia.json",
		//"http://localhost:3000/assignment132/assignment/main/duckduckgo",
		//"http://localhost:3000/assignment132/assignment/main/google",
		//"http://localhost:3000/assignment132/assignment/main/wikipedia",
	}

	serverCfg := models.Config{
		// Server port to run on it
		ListenAddr: ":8000",

		// The list of external resources to get data
		Resources: resources,
	}

	serverCfg.Validate()

	// Success
	return serverCfg, nil
}

