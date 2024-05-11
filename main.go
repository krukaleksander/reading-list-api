package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	router := handleRoutes()

	fmt.Printf("Server is running on port %s\n", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Println("Error starting server: %s\n", err)
	}
}
