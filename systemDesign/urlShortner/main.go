package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sync"
)

var (
	urlMapping = make(map[string]string) // map[shortURL]longURL
	baseURL    = "http://short.url/"
	mutex      sync.Mutex
)

// HashFunction: Create a hash of the original URL to generate a short key
func HashFunction(longURL string) string {
	hash := sha256.New()
	hash.Write([]byte(longURL))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))[:8] // Use first 8 characters
}

// GenerateShortURL: Create a short URL from a long URL
func GenerateShortURL(longURL string) string {
	mutex.Lock()
	defer mutex.Unlock()

	shortKey := HashFunction(longURL)
	shortURL := baseURL + shortKey

	if _, exists := urlMapping[shortURL]; !exists {
		urlMapping[shortURL] = longURL
	}

	return shortURL
}

// RetrieveLongURL: Retrieve the original long URL using the short URL
func RetrieveLongURL(shortURL string) (string, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	longURL, exists := urlMapping[shortURL]
	return longURL, exists
}

func main() {
	longURL := "https://example.com/some/very/long/url/that/needs/shortening"

	// Generate a short URL
	shortURL := GenerateShortURL(longURL)
	fmt.Println("Short URL:", shortURL)

	// Retrieve the original long URL
	if longURL, found := RetrieveLongURL(shortURL); found {
		fmt.Println("Original URL:", longURL)
	} else {
		fmt.Println("URL not found")
	}
}
