package extract_links

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	extractLinks := ExtractLinks{}
	htmlBody := strings.NewReader(`
		<html>
		<body>
		    <h1>Hello!</h1>
			<a href="/another-page">
				A link to
				<span>internal page</span>
			</a>
		    <a href="https://example.com/about/">Link with slash</a>
			<a href="https://example.com/about">Same link without slash</a>
			<a href="https://youtube.com/jsfunc/">A link to external page</a>
		</body>
		</html>
	`)

	links, err := extractLinks.All(htmlBody)

	if err != nil {
		t.Error("Err should be nil")
	}

	if len(links) != 3 {
		t.Error("Links count should be 3, we are removing duplicates")
	}

	for i, link := range links {
		if i == 0 {
			if link.Href != "/another-page" {
				t.Error("Anchor link href is invalid")
			}
			if link.Text != "A link to internal page" {
				t.Error("Anchor link text is invalid")
			}
		}

		if i == 1 {
			//since we are removing trailing slash
			if link.Href != "https://example.com/about" {
				t.Error("Anchor link href is invalid")
			}
			if link.Text != "Link with slash" {
				t.Error("Anchor link text is invalid")
			}
		}
	}
}

func TestRemoveTrailingSlash(t *testing.T) {
	href := "https://example.js.org/about/"
	cleanHref := removeTrailingSlash(href)

	if strings.Compare(cleanHref, "https://example.js.org/about") != 0 {
		t.Error("Trailing slash should be removed")
	}
}
