package main

import (
	"fmt"
	"os"
)

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	args := os.Args[1:]

	app, err := newApp()
	app.Logger.LogError("Failed to create new application", err)

	if len(args) == 0 {
		app.Logger.LogError("URL is missing, e.g. webscrapper http://js.org/", nil)
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

	app.Logger.LogInfo("===========================================================")
	app.Logger.LogInfo(fmt.Sprintf("Done crawling host: %s\n", baseUrl))
}

//func SignalHandler() {
//	gracefulShutdown := make(chan os.Signal, 1)
//	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)
//
//	select {
//	case <-gracefulShutdown:
//		fmt.Println("shutting down gracefully")
//	}
//
//	//for s := <-c; ; s = <-c {
//	//	switch s {
//	//	case os.Interrupt:
//	//		fmt.Println("^C received")
//	//		fmt.Println("<----------- ----------- ----------- ----------->")
//	//		fmt.Println("<----------- ----------- ----------- ----------->")
//	//		graphMap.CreatePath("https://youtube.com/jsfunc", "https://youtube.com/YouTubeRedOriginals")
//	//		os.Exit(0)
//	//	case os.Kill:
//	//		fmt.Println("SIGKILL received")
//	//		os.Exit(1)
//	//	}
//	//}
//}
