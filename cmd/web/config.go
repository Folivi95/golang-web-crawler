package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	BaseUrl            string
	CrawlExternalLinks bool
}

func loadAppConfig() (AppConfig, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// set base URL
	baseUrl := os.Getenv("BASE_URL")

	// set crawl external links value
	crawlExternalLinks := os.Getenv("CRAWL_EXTERNAL_LINKS")
	CrawlExternalLinks, _ = strconv.ParseBool(crawlExternalLinks)

	return AppConfig{
		BaseUrl: baseUrl,
	}, nil
}
