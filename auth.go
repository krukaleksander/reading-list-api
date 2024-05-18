package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !validateCredentials(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func validateCredentials(username, password string) bool {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	usernameEnv := os.Getenv("AUTH_USERNAME")
	passwordEnv := os.Getenv("AUTH_PASSWORD")
	return username == usernameEnv && password == passwordEnv
}

func BypassAuthForPublicEndpoints(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/health-check" {
			next.ServeHTTP(w, r)
			return
		}
		BasicAuthMiddleware(next).ServeHTTP(w, r)
	})
}
