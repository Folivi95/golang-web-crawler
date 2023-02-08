package crawler

import (
	"context"
	"fmt"
	"golang-web-crawler/internal/adapters/extract_links"
	"golang-web-crawler/internal/adapters/graph"
	"net/http"
	"net/url"
)

type Crawler struct {
	GraphMap   graph.Graph
	HttpClient *http.Client
	HasCrawled map[string]bool
	UrlQueue   chan string
}

func NewCrawler(graphMap graph.Graph,
	httpClient *http.Client,
	hasCrawled map[string]bool,
	urlQueue chan string) *Crawler {

	return &Crawler{
		GraphMap:   graphMap,
		HttpClient: httpClient,
		HasCrawled: hasCrawled,
		UrlQueue:   urlQueue,
	}
}

func (c *Crawler) CrawlLink(ctx context.Context, baseHref string) {
	addVertexStatus := c.GraphMap.AddVertex(baseHref)
	fmt.Println("addVertexStatus: ", addVertexStatus)

	c.HasCrawled[baseHref] = true

	fmt.Println("Crawling... ", baseHref)

	resp, err := c.HttpClient.Get(baseHref)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	links, err := extract_links.All(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	for _, l := range links {
		if l.Href == "" {
			continue
		}
		fixedUrl := toFixedUrl(l.Href, baseHref)
		if baseHref != fixedUrl {
			addEdgeStatus := c.GraphMap.AddEdge(baseHref, fixedUrl)
			fmt.Println("addEdgeStatus: ", addEdgeStatus)
		}

		go func(url string) {
			c.UrlQueue <- url
		}(fixedUrl)
	}
}

func toFixedUrl(href, base string) string {
	uri, err := url.Parse(href)

	if err != nil || uri.Scheme == "mailto" || uri.Scheme == "tel" {
		return base
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}

	if uri.Host != baseUrl.Host {
		return base
	}

	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}
