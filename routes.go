package main

import (
	"net/http"
)

func handleRoutes() {
	http.HandleFunc("/health-check", healthCheckHandler)
}
