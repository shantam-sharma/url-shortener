package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shantam-sharma/url-shortner/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

type CreateURLRequest struct {
	URL string `json:"url"`
}

type CreateURLResponse struct {
	ShortCode string `json:"short_code"`
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

	url, err := h.service.Create(req.URL)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	response := CreateURLResponse{
		ShortCode: url.ShortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
