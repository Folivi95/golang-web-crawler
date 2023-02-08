package main

import (
	"crypto/tls"
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

	// set up a new crawler
	crawlingService := crawler.NewCrawler(graphMap, httpClient, hasCrawled, urlQueue)

	return &App{
		HttpConfig: httpConfig,
		Transport:  transport,
		HttpClient: httpClient,
		GraphMap:   graphMap,
		HasCrawled: hasCrawled,
		UrlQueue:   urlQueue,
		Crawler:    crawlingService,
	}, nil
}
