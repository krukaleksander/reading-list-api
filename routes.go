package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

func handleRoutes(dbConnection *pgx.Conn) *mux.Router {
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
	return router
}
