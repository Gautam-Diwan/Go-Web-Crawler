package main

import (
	"fmt"
	"golang/web_crawler/url_fetcher"
)

var (
	cache = make(map[string]bool)
)

// Crawl function to crawl pages recursively
func Crawl(url string, depth int, fetcher url_fetcher.Fetcher) {

	if depth <= 0 {
		return
	}

	if cache[url] {
		return
	}
	cache[url] = true

	urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("found: ", url)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
}

func main() {
	url := "https://go.dev/"
	fetcher := url_fetcher.NewHttpFetcher()
	Crawl(url, 3, fetcher)
}
