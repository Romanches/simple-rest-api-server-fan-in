package api

import (
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/handlers"
	"github.com/Romanches/simple-rest-api-server-fan-in/internal/api/v1/models"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

// API incapsulates necessary dependencies for running server
type API struct {
	config      models.Config
	healthHandler *handlers.HealthHandler

	requestWriter io.Writer
}

// NewAPI creates API struct with dependencies
func NewAPI(
	config models.Config,
	health *handlers.HealthHandler,

) *API {

	return &API{
		config:      config,
		healthHandler: health,
	}
}

// Run starts an API server
func (a *API) Run() {
	r := mux.NewRouter()

	a.registerRoutes(r)

	log.Println("Starting server on ", a.config.ListenAddr)
	muxWithMiddlewares := http.TimeoutHandler(r, 60 * time.Second, "Timeout!")
	log.Fatal(http.ListenAndServe(a.config.ListenAddr, muxWithMiddlewares))

}

func (a API) registerRoutes(r *mux.Router) {
	// Endpoints, handler functions and HTTP method
	r.HandleFunc("/health", a.healthHandler.HealthCheck).Methods("GET")

}
