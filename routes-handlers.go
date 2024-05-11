package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is up and running!")
}
