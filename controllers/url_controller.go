package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/models"
	"url-shortener/services"
)

// URLController handles HTTP requests for URL operations
type URLController struct {
	urlService *services.URLService
}

// NewURLController creates a new instance of URLController
func NewURLController(urlService *services.URLService) *URLController {
	return &URLController{
		urlService: urlService,
	}
}

// WelcomeHandler handles GET requests to the root path "/"
// Provides a simple welcome page with basic API usage information
func (controller *URLController) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome endpoint accessed")

	// Simple text-based homepage with essential information
	welcomeMessage := `Welcome to the URL Shortener Service!

📋 Available Endpoints:

1. GET  /              - This welcome page
2. POST /shorten       - Create a short URL
3. GET  /redirect/{id} - Redirect to original URL

🚀 Quick Usage:

1. Shorten a URL:
   POST /shorten
   Content-Type: application/json
   Body: {"url": "https://your-long-url.com"}
   
   Response: {"short_url": "a1b2c3d4"}

2. Use the short URL:
   Visit: /redirect/a1b2c3d4
   → Redirects to your original URL

� Example with cURL:
curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"url":"https://www.linkedin.com/in/tanayvijay/"}'

Built with Go | Port 8080
`

	fmt.Fprint(w, welcomeMessage)
}

// ShortenURLHandler handles POST requests to "/shorten"
// Accepts a JSON payload with a URL and returns a shortened version
func (controller *URLController) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON from request body
	var requestData models.ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Body - Unable to parse JSON", http.StatusBadRequest)
		return
	}

	// Validate that URL is provided
	if requestData.URL == "" {
		http.Error(w, "URL field is required", http.StatusBadRequest)
		return
	}

	// Generate shortened URL using service
	shortURLHash, err := controller.urlService.CreateShortenedURL(requestData.URL)
	if err != nil {
		http.Error(w, "Failed to create shortened URL: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := models.ShortenResponse{
		ShortURL: shortURLHash,
	}

	// Set response headers and send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Printf("Shortened URL created: %s -> %s\n", requestData.URL, shortURLHash)
}

// RedirectURLHandler handles GET requests to "/redirect/{shortURLHash}"
// Redirects the user to the original URL associated with the short URL hash
func (controller *URLController) RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL hash from the URL path
	shortURLHash := r.URL.Path[len("/redirect/"):]

	if shortURLHash == "" {
		http.Error(w, "Short URL hash is required", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL using service
	urlEntry, err := controller.urlService.GetOriginalURL(shortURLHash)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	fmt.Printf("Redirecting %s to %s\n", shortURLHash, urlEntry.OriginalURL)
	http.Redirect(w, r, urlEntry.OriginalURL, http.StatusFound)
}
