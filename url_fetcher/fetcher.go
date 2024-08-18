package url_fetcher

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// Fetcher interface definition
type Fetcher interface {
	Fetch(url string) (urls []string, err error)
}

// HTTPFetcher struct definition
type HTTPFetcher struct {
	Client *http.Client
}

// HTTPFetcher.Fetch method implementation
func (f *HTTPFetcher) Fetch(url string) ([]string, error) {
	resp, err := f.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	urls, err := extractURLs(url, string(body))
	if err != nil {
		return nil, fmt.Errorf("failed to extract URLs: %v", err)
	}

	return urls, nil
}

// extractURLs function to extract URLs from HTML body
func extractURLs(baseURL string, body string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	var urls []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					u, err := url.Parse(attr.Val)
					if err != nil || !u.IsAbs() {
						continue
					}
					base, err := url.Parse(baseURL)
					if err != nil {
						continue
					}
					urls = append(urls, base.ResolveReference(u).String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return urls, nil
}

// Timeouts on 10 second mark
func NewHttpFetcher() *HTTPFetcher {
	return &HTTPFetcher{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
