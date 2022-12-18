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

// API incapsulates necessary dependencies to run server
type API struct {
	config      models.Config
	healthHandler *handlers.HealthHandler
	dataHandler  *handlers.DataHandler

	requestWriter io.Writer
}

// NewAPI creates the new API instance with dependencies
func NewAPI(
	config models.Config,
	health *handlers.HealthHandler,
	dataHandler *handlers.DataHandler,

) *API {

	return &API{
		config:      config,
		healthHandler: health,
		dataHandler:  dataHandler,
	}
}

// Run starts the Rest-API server
func (a *API) Run() {
	r := mux.NewRouter()

	a.registerRoutes(r)

	log.Println("Starting server on ", a.config.ListenAddr)
	muxWithMiddlewares := http.TimeoutHandler(r, 60 * time.Second, "Timeout!")
	log.Fatal(http.ListenAndServe(a.config.ListenAddr, muxWithMiddlewares))

}

// Middleware
func (a API) registerRoutes(r *mux.Router) {

	r.HandleFunc("/health", a.healthHandler.HealthCheck).Methods("GET")

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/data",  a.dataHandler.Get).Methods("GET")

}
