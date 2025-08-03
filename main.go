package main

import (
	"fmt"
	"net/http"
	"url-shortener/controllers"
	"url-shortener/repositories"
	"url-shortener/services"
)

// main is the entry point of the URL shortener application
// Sets up dependency injection, HTTP routes and starts the web server
func main() {
	fmt.Println("URL Shortener Service Starting...")

	// Dependency Injection - Create instances in proper order
	urlRepo := repositories.NewURLRepository()
	urlService := services.NewURLService(urlRepo)
	urlController := controllers.NewURLController(urlService)

	// Register HTTP route handlers
	http.HandleFunc("/", urlController.WelcomeHandler)
	http.HandleFunc("/shorten", urlController.ShortenURLHandler)
	http.HandleFunc("/redirect/", urlController.RedirectURLHandler)

	// Start the HTTP server
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
