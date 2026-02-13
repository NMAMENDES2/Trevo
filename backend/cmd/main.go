package main

import (
	"log"
	"os"

	"github.com/NMAMENDES2/Trevo/api"
	"github.com/NMAMENDES2/Trevo/db"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found... Using defaults")
	}

	database, err := db.New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer database.Close()

	server := api.NewServer(database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s\n", port)
	if err := server.Start(":" + port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
