package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is up and running!")
}

func createRecordHandler(w http.ResponseWriter, r *http.Request, dbConnection *pgx.Conn) {
	var record Record

	jsonParsingError := json.NewDecoder(r.Body).Decode(&record)

	if jsonParsingError != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)

		return
	}

	err := insertNewRecord(dbConnection, record)

	if err != nil {
		http.Error(w, "Failed to insert record", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)

}
