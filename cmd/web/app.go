package main

import (
	"crypto/tls"
	"golang-web-crawler/internal/adapters/graph"
	"golang.org/x/net/context"
	"net/http"
)

const URLQueueChannelCapacity = 10

var (
	urlQueue   = make(chan string, URLQueueChannelCapacity)
	hasCrawled = make(map[string]bool)
)

type App struct {
	HttpConfig *tls.Config
	Transport  *http.Transport
	HttpClient *http.Client
	GraphMap   *graph.Graph
}

func newApp(applicationContext context.Context) (*App, error) {
	httpConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{
		TLSClientConfig: httpConfig,
	}
	httpClient := &http.Client{
		Transport: transport,
	}
	graphMap := graph.NewGraph()

	return &App{
		HttpConfig: httpConfig,
		Transport:  transport,
		HttpClient: httpClient,
		GraphMap:   graphMap,
	}, nil
}
