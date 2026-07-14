package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/shantam-sharma/url-shortener/internal/database"
	"github.com/shantam-sharma/url-shortener/internal/handlers"
	"github.com/shantam-sharma/url-shortener/internal/middleware"
	"github.com/shantam-sharma/url-shortener/internal/repository"
	"github.com/shantam-sharma/url-shortener/internal/service"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := middleware.CORS(http.DefaultServeMux)

	log.Printf("Server running on :%s", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
