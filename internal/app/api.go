package app

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

// API incapsulates necessary dependencies for running server
type API struct {

	requestWriter io.Writer
}

// NewAPI creates API struct with dependencies
func NewAPI(

) *API {

	return &API{}
}

// Run starts an API server
func (a *API) Run() {
	r := mux.NewRouter()
	
	a.registerRoutes(r)

	log.Println("Starting server on ", "8000")
	log.Fatal(http.ListenAndServe("8000", r))

}

func (a API) registerRoutes(r *mux.Router) {

}
