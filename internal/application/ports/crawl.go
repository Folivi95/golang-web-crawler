package ports

import "context"

type Crawl interface {
	CrawlLink(ctx context.Context, baseHref string)
}
