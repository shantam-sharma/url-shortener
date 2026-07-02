package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/shantam-sharma/url-shortner/internal/database"
	"github.com/shantam-sharma/url-shortner/internal/handlers"
	"github.com/shantam-sharma/url-shortner/internal/repository"
	"github.com/shantam-sharma/url-shortner/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
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

	// Start HTTP server
	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
