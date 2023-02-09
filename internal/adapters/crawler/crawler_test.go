package crawler

import (
	"context"
	"go.uber.org/zap"
	"golang-web-crawler/internal/adapters/graph"
	"net/http"
	"testing"
)

func TestCrawler_CrawlLink(t *testing.T) {
	// http client
	httpClient := &http.Client{}

	// instantiate new graph
	graphMap := graph.NewGraph()

	// create map of crawled URLs
	hasCrawled := make(map[string]bool)

	// create URL queue
	urlQueue := make(chan string, 2)

	// add logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// given
	crawler := NewCrawler(graphMap, httpClient, hasCrawled, urlQueue,
		CrawlExternalLinks, logger)
	baseUrl := "https://example.com"

	// when link is crawled
	crawler.CrawlLink(context.Background(), baseUrl)

	// then length of channel should be greater than one
	select {
	case url := <-urlQueue:
		if url == "" {
			t.Error("Expected ", baseUrl, ", Got ", url)
		}
	}
	close(urlQueue)
}
func TestToFixedUrl(t *testing.T) {
	fixedUrl := toFixedUrl("https://example.com/aboutus.html", "https://example.com/")
	if fixedUrl != "https://example.com/aboutus.html" {
		t.Error("toFixedUrl did not get expected href")
	}

	mailToUrl := toFixedUrl("mailto:ajinkya@gmail.com", "https://example.com/")
	if mailToUrl != "https://example.com/" {
		t.Error("expected baseUrl instead of mailto link")
	}

	telephoneUrl := toFixedUrl("tel://9820098200", "https://example.com/")
	if telephoneUrl != "https://example.com/" {
		t.Error("expected baseUrl instead of telephone link")
	}
}
