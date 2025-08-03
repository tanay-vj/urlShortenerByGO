package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// GenerateShortURLHash creates a unique 8-character hash from the original URL
// Uses MD5 hashing algorithm to ensure consistent results for the same URL
func GenerateShortURLHash(originalURL string) string {
	// Create MD5 hasher instance
	hasher := md5.New()

	// Convert original URL string to byte slice and hash it
	hasher.Write([]byte(originalURL))
	hashBytes := hasher.Sum(nil)

	// Convert hash bytes to hexadecimal string and take first 8 characters
	fullHash := hex.EncodeToString(hashBytes)
	shortHash := fullHash[:8]

	fmt.Printf("Generated hash for URL '%s': %s\n", originalURL, shortHash)
	return shortHash
}
