package crawler

import (
	"context"
	"fmt"
	"golang-web-crawler/internal/adapters/extract_links"
	"golang-web-crawler/internal/adapters/graph"
	"golang-web-crawler/internal/application/ports"
	"net/http"
	"net/url"
)

type Crawler struct {
	GraphMap      *graph.Graph
	HttpClient    *http.Client
	HasCrawled    map[string]bool
	UrlQueue      chan string
	LinkExtractor ports.LinksExtractor
}

func NewCrawler(graphMap *graph.Graph,
	httpClient *http.Client,
	hasCrawled map[string]bool,
	urlQueue chan string) *Crawler {

	linkExtractor := extract_links.ExtractLinks{}

	return &Crawler{
		GraphMap:      graphMap,
		HttpClient:    httpClient,
		HasCrawled:    hasCrawled,
		UrlQueue:      urlQueue,
		LinkExtractor: linkExtractor,
	}
}

func (c *Crawler) ProcessBaseUrl(ctx context.Context, baseHref string) {
	go func() {
		select {
		case <-ctx.Done():
			close(c.UrlQueue)
		default:
			c.UrlQueue <- baseHref
		}
	}()
}

func (c *Crawler) CrawlLink(ctx context.Context, baseHref string) {
	addVertexStatus := c.GraphMap.AddVertex(baseHref)
	fmt.Println("addVertexStatus: ", addVertexStatus)

	c.HasCrawled[baseHref] = true

	fmt.Println("Crawling... ", baseHref)

	req, err := http.NewRequestWithContext(ctx, "GET", "", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	links, err := c.LinkExtractor.All(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	for _, l := range links {
		if l.Href == "" {
			continue
		}
		fixedUrl := toFixedUrl(l.Href, baseHref)
		if baseHref != fixedUrl {
			c.GraphMap.AddEdge(baseHref, fixedUrl)
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
