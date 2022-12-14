package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Endpoints, handler functions and HTTP method
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	http.ListenAndServe(":8000", r)

}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "Ok!")
}
