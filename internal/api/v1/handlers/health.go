package handlers

import (
	"github.com/Romanches/simple-rest-api-server-fan-in/pkg/render"
	"net/http"
)

// Handler provides health-check middleware
type HealthHandler struct {
	//
}

// NewHealthHandler creates a new Handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Handler for /health endpoint
func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	//fmt.Fprintf(w, "Ok!")

	render.OK(w, "Ok!")
}
