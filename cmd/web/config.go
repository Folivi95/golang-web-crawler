package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	BaseUrl string
}

func loadAppConfig() (AppConfig, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// set base URL
	baseUrl := os.Getenv("BASE_URL")

	return AppConfig{
		BaseUrl: baseUrl,
	}, nil
}
