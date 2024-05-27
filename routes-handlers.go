package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v4"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)

}

func getAllRecordsHandler(w http.ResponseWriter, r *http.Request, dbConnection *pgx.Conn) {
	records, err := getAllRecords(dbConnection)

	if err != nil {
		http.Error(w, "Failed to fetch records", http.StatusInternalServerError)
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func removeRecordHandler(w http.ResponseWriter, r *http.Request, dbConnection *pgx.Conn) {
	idFromParam := r.URL.Query().Get("id")
	if idFromParam == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idFromParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = removeRecord(dbConnection, id)
	if err != nil {
		http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		return
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Record deleted successfully"}`))
}
