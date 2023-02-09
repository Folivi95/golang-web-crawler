package main

import (
	"go.uber.org/zap"
	"os"
)

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	// create new application
	app, err := newApp()
	if err != nil {
		app.Logger.Error("Failed to create new application", zap.Error(err))
	}

	// load application config
	appConfig, err := loadAppConfig()
	if err != nil {
		app.Logger.Error("Failed to load appConfig", zap.Error(err))
	}

	// check if base URL is set in app config
	if appConfig.BaseUrl != "" {
		// send base URL to channel
		app.UrlQueue <- appConfig.BaseUrl
	} else {
		// check if base URL is passed as command line argument
		args := os.Args[1:]
		if len(args) == 0 {
			app.Logger.Error("URL is missing, e.g. webscrapper https://js.org/", zap.Error(err))
			os.Exit(1)
		}

		// set base URL
		baseUrl := args[0]
		app.UrlQueue <- baseUrl
	}

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
	app.Logger.Info("Done crawling host")
}
