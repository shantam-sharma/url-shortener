package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/shantam-sharma/url-shortner/internal/database"
	"github.com/shantam-sharma/url-shortner/internal/handlers"
	"github.com/shantam-sharma/url-shortner/internal/middleware"
	"github.com/shantam-sharma/url-shortner/internal/repository"
	"github.com/shantam-sharma/url-shortner/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	// Connect to PostgreSQL
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	log.Println("Connected to PostgreSQL")

	// Initialize application layers
	urlRepository := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepository)
	urlHandler := handlers.NewURLHandler(urlService)

	// Register routes
	http.HandleFunc("/api/v1/urls", urlHandler.CreateURL)
	http.HandleFunc("/", urlHandler.RedirectURL)
	// Start HTTP server
	// Start HTTP server
	log.Println("Server running on :8080")

	handler := middleware.CORS(http.DefaultServeMux)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
