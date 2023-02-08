package main

import "os"

type AppConfig struct {
	LogLevel string
	BaseUrl  string
}

func loadAppConfig() (AppConfig, error) {
	baseUrl := os.Getenv("BASE_URL")
	logLevel := os.Getenv("LOG_LEVEL")

	return AppConfig{
		LogLevel: logLevel,
		BaseUrl:  baseUrl,
	}, nil
}
