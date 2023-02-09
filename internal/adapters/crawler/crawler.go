package crawler

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"golang-web-crawler/internal/adapters/extract_links"
	"golang-web-crawler/internal/adapters/graph"
	"golang-web-crawler/internal/application/ports"
	"io"
	"net/http"
	"net/url"
)

var CrawlExternalLinks = false

type Crawler struct {
	GraphMap      *graph.Graph
	HttpClient    *http.Client
	HasCrawled    map[string]bool
	UrlQueue      chan string
	LinkExtractor ports.LinksExtractor
	Logger        *zap.Logger
}

func NewCrawler(graphMap *graph.Graph,
	httpClient *http.Client,
	hasCrawled map[string]bool,
	urlQueue chan string,
	crawlExternalLinks bool,
	logger *zap.Logger) *Crawler {

	linkExtractor := extract_links.ExtractLinks{}
	CrawlExternalLinks = crawlExternalLinks

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

	c.Logger.Info(fmt.Sprintf("Crawling... %s", baseHref))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseHref, nil)
	if err != nil {
		c.Logger.Error("Failed to create http request with context", zap.Error(err))
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Failed to make GET request to %s", baseHref), zap.Error(err))
	}

	// close response body reader
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.Logger.Error("Error closing http response body", zap.Error(err))
		}
	}(resp.Body)

	links, err := c.LinkExtractor.All(resp.Body)
	if err != nil {
		c.Logger.Error("Failed to extract links", zap.Error(err))
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
	// parse href URL
	uri, err := url.Parse(href)

	// check if href URL is an email address or phone number
	if err != nil || uri.Scheme == "mailto" || uri.Scheme == "tel" {
		return base
	}

	// parse baseUrl
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}

	// check if crawler should crawl external links
	if !CrawlExternalLinks && uri.Host != baseUrl.Host {
		return base
	}

	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}
