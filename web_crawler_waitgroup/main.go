package main

import (
	"fmt"
	"golang/web_crawler/url_fetcher"
	"sync"
)

var (
	cache = make(map[string]bool)
	mu    sync.Mutex
	wg    sync.WaitGroup
)

// Crawl function to crawl pages recursively
func Crawl(url string, depth int, fetcher url_fetcher.Fetcher) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	mu.Lock()
	if cache[url] {
		mu.Unlock()
		return
	}
	cache[url] = true
	mu.Unlock()

	urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("found: ", url)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
}

func main() {
	url := "https://go.dev/"
	fetcher := url_fetcher.NewHttpFetcher()
	wg.Add(1)
	go Crawl(url, 3, fetcher)
	wg.Wait()

	fmt.Printf("len(cache): %v\n", len(cache))
}
