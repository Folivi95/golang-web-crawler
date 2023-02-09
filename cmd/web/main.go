package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	args := os.Args[1:]

	app, err := newApp()
	if err != nil {
		app.Logger.Error("Failed to create new application", zap.Error(err))
	}

	if len(args) == 0 {
		app.Logger.Error("URL is missing, e.g. webscrapper https://js.org/", zap.Error(err))
		os.Exit(1)
	}

	baseUrl := args[0]

	// pass base url to crawler to begin processing of base url
	app.Crawler.ProcessBaseUrl(ctx, baseUrl)

	for href := range app.UrlQueue {
		if !app.HasCrawled[href] {
			app.Crawler.CrawlLink(ctx, href)
		}

		// close channel if it has stopped receiving unprocessed URLs
		if len(app.UrlQueue) == 0 {
			close(app.UrlQueue)
		}
	}

	app.Logger.Info("===========================================================")
	app.Logger.Info(fmt.Sprintf("Done crawling host: %s\n", baseUrl))
}
