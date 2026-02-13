package main

import (
	"log"
	"os"

	"github.com/NMAMENDES2/Trevo/api"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found... Using defaults")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	server := api.NewServer(port)

	log.Printf("Starting server on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
