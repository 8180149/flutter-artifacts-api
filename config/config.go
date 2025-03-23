package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var APIKey string
var Port string
var ArtifactDir string

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	APIKey = os.Getenv("API_KEY")
	if APIKey == "" {
		log.Fatal("API_KEY is required in .env file")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	ArtifactDir = "artifacts"
}
