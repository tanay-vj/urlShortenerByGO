package services

import (
	"errors"
	"time"
	"url-shortener/models"
	"url-shortener/repositories"
	"url-shortener/utils"
)

// URLService handles business logic for URL operations
type URLService struct {
	urlRepo *repositories.URLRepository
}

// NewURLService creates a new instance of URLService
func NewURLService(urlRepo *repositories.URLRepository) *URLService {
	return &URLService{
		urlRepo: urlRepo,
	}
}

// CreateShortenedURL generates a short URL and stores it in the database
// Returns the generated short URL hash and any error
func (service *URLService) CreateShortenedURL(originalURL string) (string, error) {
	// Validate input
	if originalURL == "" {
		return "", errors.New("original URL cannot be empty")
	}

	// Generate unique hash for the original URL
	shortURLHash := utils.GenerateShortURLHash(originalURL)
	urlID := shortURLHash

	// Create URL entry
	urlEntry := models.URL{
		ID:           urlID,
		OriginalURL:  originalURL,
		ShortURL:     shortURLHash,
		CreationDate: time.Now(),
	}

	// Store in database
	err := service.urlRepo.Save(urlEntry)
	if err != nil {
		return "", err
	}

	return shortURLHash, nil
}

// GetOriginalURL retrieves the original URL using the short URL hash
// Returns the URL entry and any error
func (service *URLService) GetOriginalURL(shortURLHash string) (models.URL, error) {
	// Validate input
	if shortURLHash == "" {
		return models.URL{}, errors.New("short URL hash cannot be empty")
	}

	// Retrieve from database
	urlEntry, err := service.urlRepo.FindByShortHash(shortURLHash)
	if err != nil {
		return models.URL{}, err
	}

	return urlEntry, nil
}
