package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(req.URL)

	resp := ShortenResponse{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func generateShortURL(originalURL string) string {
	// Placeholder: Simply return a dummy shortened URL for now
	return fmt.Sprintf("http://localhost:8080/%s", "abcd123")
}
