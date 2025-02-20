package main

import (
	"DatabasePipeline/config"
	"DatabasePipeline/internal/creator"
	"DatabasePipeline/internal/db"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func startCreating(gormDB *gorm.DB) {

	log.Println("Starting video fetcher...") // Add logging

	creator.GenerateAndStoreVideos(gormDB)
}

func main() {
	// Load configuration

	cfg := config.LoadConfig()

	// Initialize the database with configuration
	db.InitDB(cfg)
	gormDB := db.DB

	// Auto-migrate the Video model
	if err := gormDB.AutoMigrate(&creator.Video{}); err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	// Start background YouTube data fetching
	go startCreating(gormDB)

	// Start the HTTP server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
