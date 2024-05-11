package main

import (
	"fmt"
	"net/http"
)

func main() {
	handleRoutes()

	fmt.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server: %s\n", err)
	}
}
