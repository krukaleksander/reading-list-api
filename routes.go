package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/rs/cors"
)

func handleRoutes(dbConnection *pgx.Conn) http.Handler {
	router := mux.NewRouter()
	router.Use(AuthMidlleware)
	router.HandleFunc("/health-check", healthCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		createRecordHandler(w, r, dbConnection)
	}).Methods(http.MethodPost)
	router.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		getAllRecordsHandler(w, r, dbConnection)
	}).Methods(http.MethodGet)
	router.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		removeRecordHandler(w, r, dbConnection)
	}).Methods(http.MethodDelete)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	return corsHandler.Handler(router)
}
