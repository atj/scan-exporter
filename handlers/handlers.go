package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// HandleFunc handles /metrics and not found path.
func HandleFunc() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	r.Handle("/health", http.HandlerFunc(healthCheckPage))
	r.NotFoundHandler = http.HandlerFunc(notFoundPage)

	return r
}

// notFoundPage set the response header to 404 status and prints an error message.
func notFoundPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 page not found</h1>")
}

func healthCheckPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I'm up !")
}
