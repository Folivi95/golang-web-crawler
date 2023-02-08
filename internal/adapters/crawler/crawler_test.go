package crawler

import "testing"

func TestToFixedUrl(t *testing.T) {
	fixedUrl := toFixedUrl("http://example.com/aboutus.html", "http://example.com/")
	if fixedUrl != "http://example.com/aboutus.html" {
		t.Error("toFixedUrl did not get expected href")
	}

	mailToUrl := toFixedUrl("mailto:ajinkya@gmail.com", "http://example.com/")
	if mailToUrl != "http://example.com/" {
		t.Error("expected baseUrl instead of mailto link")
	}

	telephoneUrl := toFixedUrl("tel://9820098200", "http://example.com/")
	if telephoneUrl != "http://example.com/" {
		t.Error("expected baseUrl instead of telephone link")
	}
}
