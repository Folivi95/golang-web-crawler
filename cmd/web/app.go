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
	GraphMap   ports.GraphStructure
	HasCrawled map[string]bool
	Crawler    ports.Crawl
	UrlQueue   chan string
	Logger     *zap.Logger
}

func newApp() (*App, error) {
	// define http client
	httpConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{
		TLSClientConfig: httpConfig,
	}
	httpClient := &http.Client{
		Transport: transport,
	}

	// instantiate new graph
	graphMap := graph.NewGraph()

	// create map of crawled URLs
	hasCrawled := make(map[string]bool)

	// create URL queue
	urlQueue := make(chan string, URLQueueChannelCapacity)

	// set crawl external links
	crawlExternalLinks := false

	// add logger
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	if err != nil {
		logger.Error("Failed to create logger", zap.Error(err))
		return nil, err
	}

	// set up a new crawler
	crawlingService := crawler.NewCrawler(graphMap, httpClient, hasCrawled,
		urlQueue, crawlExternalLinks, logger)

	return &App{
		GraphMap:   graphMap,
		HasCrawled: hasCrawled,
		UrlQueue:   urlQueue,
		Crawler:    crawlingService,
		Logger:     logger,
	}, nil
}
