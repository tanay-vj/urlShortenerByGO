package models

import "time"

// URL represents a shortened URL entry with all its metadata
type URL struct {
	ID           string    `json:"id"`            // Unique identifier (same as short URL)
	OriginalURL  string    `json:"original_url"`  // The original long URL
	ShortURL     string    `json:"short_url"`     // The shortened URL identifier
	CreationDate time.Time `json:"creation_date"` // When this URL was created
}

// ShortenRequest represents the JSON request structure for shortening URLs
type ShortenRequest struct {
	URL string `json:"url"` // The URL to be shortened
}

// ShortenResponse represents the JSON response structure after shortening
type ShortenResponse struct {
	ShortURL string `json:"short_url"` // The generated short URL hash
}
