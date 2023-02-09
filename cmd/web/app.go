package main

import (
	"crypto/tls"
	"go.uber.org/zap"
	"golang-web-crawler/internal/adapters/crawler"
	"golang-web-crawler/internal/adapters/graph"
	"golang-web-crawler/internal/application/ports"
	"net/http"
)

const URLQueueChannelCapacity = 10

type App struct {
	HttpConfig *tls.Config
	Transport  *http.Transport
	HttpClient *http.Client
	GraphMap   ports.GraphStructure
	HasCrawled map[string]bool
	Crawler    ports.Crawl
	UrlQueue   chan string
	Logger     *zap.Logger
}

func newApp() (*App, error) {
	httpConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{
		TLSClientConfig: httpConfig,
	}
	httpClient := &http.Client{
		Transport: transport,
	}
	graphMap := graph.NewGraph()

	hasCrawled := make(map[string]bool)

	urlQueue := make(chan string, URLQueueChannelCapacity)

	// add logger
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	if err != nil {
		logger.Error("Failed to create logger", zap.Error(err))
	}

	// set up a new crawler
	crawlingService := crawler.NewCrawler(graphMap, httpClient, hasCrawled, urlQueue, logger)

	return &App{
		HttpConfig: httpConfig,
		Transport:  transport,
		HttpClient: httpClient,
		GraphMap:   graphMap,
		HasCrawled: hasCrawled,
		UrlQueue:   urlQueue,
		Crawler:    crawlingService,
		Logger:     logger,
	}, nil
}
