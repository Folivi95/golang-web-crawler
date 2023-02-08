package main

import (
	"fmt"
	"os"
)

//var (
//	config    = &tls.Config{InsecureSkipVerify: true}
//	transport = &http.Transport{
//		TLSClientConfig: config,
//	}
//	netClient *http.Client
//	graphMap  = graph.NewGraph()
//)

//func init() {
//	netClient = &http.Client{
//		Transport: transport,
//	}
//	go SignalHandler()
//}

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("URL is missing, e.g. webscrapper http://js.org/")
		os.Exit(1)
	}

	baseUrl := args[0]

	app, err := newApp()
	checkErr(err)

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

	fmt.Println("===========================================================")
	fmt.Printf("Done crawling host: %s\n", baseUrl)
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

//func crawlLink(baseHref string) {
//	addVertexStatus := graphMap.AddVertex(baseHref)
//	fmt.Println("addVertexStatus: ", addVertexStatus)
//
//	hasCrawled[baseHref] = true
//
//	fmt.Println("Crawling... ", baseHref)
//
//	resp, err := netClient.Get(baseHref)
//	checkErr(err)
//	defer resp.Body.Close()
//
//	links, err := extract_links.All(resp.Body)
//	checkErr(err)
//
//	for _, l := range links {
//		if l.Href == "" {
//			continue
//		}
//		fixedUrl := toFixedUrl(l.Href, baseHref)
//		if baseHref != fixedUrl {
//			addEdgeStatus := graphMap.AddEdge(baseHref, fixedUrl)
//			fmt.Println("addEdgeStatus: ", addEdgeStatus)
//		}
//
//		go func(url string) {
//			urlQueue <- url
//		}(fixedUrl)
//	}
//}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

//func toFixedUrl(href, base string) string {
//	uri, err := url.Parse(href)
//
//	if err != nil || uri.Scheme == "mailto" || uri.Scheme == "tel" {
//		return base
//	}
//	baseUrl, err := url.Parse(base)
//	if err != nil {
//		return ""
//	}
//
//	if uri.Host != baseUrl.Host {
//		return base
//	}
//
//	uri = baseUrl.ResolveReference(uri)
//	return uri.String()
//}
