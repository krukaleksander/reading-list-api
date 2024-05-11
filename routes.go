package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", healthCheckHandler).Methods(http.MethodGet)
	return router
}
