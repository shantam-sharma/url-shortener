package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/shantam-sharma/url-shortener/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

type CreateURLRequest struct {
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

type CreateURLResponse struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

func (h *URLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateURLRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	url, err := h.service.Create(req.URL, req.Alias)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidURL):
			http.Error(w, err.Error(), http.StatusBadRequest)

		case errors.Is(err, service.ErrInvalidAlias):
			http.Error(w, err.Error(), http.StatusBadRequest)

		case errors.Is(err, service.ErrAliasAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)

		default:
			http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		}

		return
	}

	baseURL := os.Getenv("BASE_URL")

	if baseURL == "" {
		baseURL = "http://127.0.0.1:8080"
	}

	response := CreateURLResponse{
		OriginalURL: url.OriginalURL,
		ShortCode:   url.ShortCode,
		ShortURL:    baseURL + "/" + url.ShortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/" {
		http.NotFound(w, r)
		return
	}

	shortCode := r.URL.Path[1:]

	url, err := h.service.Resolve(shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
