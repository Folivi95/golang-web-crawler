package crawler

import "testing"

func TestCrawler_CrawlLink(t *testing.T) {

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
