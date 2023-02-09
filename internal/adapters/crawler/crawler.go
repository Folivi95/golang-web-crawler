package crawler

import (
	"context"
	"fmt"
	"golang-web-crawler/internal/adapters/extract_links"
	"golang-web-crawler/internal/adapters/graph"
	"golang-web-crawler/internal/application/ports"
	"io"
	"net/http"
	"net/url"
)

type Crawler struct {
	GraphMap      *graph.Graph
	HttpClient    *http.Client
	HasCrawled    map[string]bool
	UrlQueue      chan string
	LinkExtractor ports.LinksExtractor
	Logger        ports.Logger
}

func NewCrawler(graphMap *graph.Graph,
	httpClient *http.Client,
	hasCrawled map[string]bool,
	urlQueue chan string,
	logger ports.Logger) *Crawler {

	linkExtractor := extract_links.ExtractLinks{}

	return &Crawler{
		GraphMap:      graphMap,
		HttpClient:    httpClient,
		HasCrawled:    hasCrawled,
		UrlQueue:      urlQueue,
		LinkExtractor: linkExtractor,
		Logger:        logger,
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
	_ = c.GraphMap.AddVertex(baseHref)

	c.HasCrawled[baseHref] = true

	c.Logger.LogInfo(fmt.Sprintf("Crawling... %s", baseHref))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseHref, nil)
	if err != nil {
		c.Logger.LogError("Failed to create http request with context", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.LogError(fmt.Sprintf("Failed to make GET request to %s", baseHref), err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.Logger.LogError("Error closing http response body", err)
		}
	}(resp.Body)

	links, err := c.LinkExtractor.All(resp.Body)
	if err != nil {
		c.Logger.LogError("Failed to extract links", err)
	}

	for _, l := range links {
		if l.Href == "" {
			continue
		}
		fixedUrl := toFixedUrl(l.Href, baseHref)
		if baseHref != fixedUrl {
			_ = c.GraphMap.AddEdge(baseHref, fixedUrl)
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
