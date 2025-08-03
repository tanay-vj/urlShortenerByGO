package repositories

import (
	"errors"
	"url-shortener/models"
)

// URLRepository handles all database operations for URLs
type URLRepository struct {
	// In-memory database to store URL mappings
	// Key: shortURL (8-character hash), Value: URL struct
	urlDatabase map[string]models.URL
}

// NewURLRepository creates a new instance of URLRepository
func NewURLRepository() *URLRepository {
	return &URLRepository{
		urlDatabase: make(map[string]models.URL),
	}
}

// Save stores a URL entry in the database
func (repo *URLRepository) Save(url models.URL) error {
	repo.urlDatabase[url.ID] = url
	return nil
}

// FindByShortHash retrieves a URL entry by its short hash
// Returns the URL struct and an error if not found
func (repo *URLRepository) FindByShortHash(shortURLHash string) (models.URL, error) {
	urlEntry, exists := repo.urlDatabase[shortURLHash]
	if !exists {
		return models.URL{}, errors.New("URL not found for hash: " + shortURLHash)
	}
	return urlEntry, nil
}

// GetAllURLs returns all stored URLs (useful for debugging/admin)
func (repo *URLRepository) GetAllURLs() map[string]models.URL {
	return repo.urlDatabase
}
