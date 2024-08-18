# Go Web Crawler

This project contains two versions of a simple web crawler implemented in Go. The crawler extracts URLs from web pages recursively.

## Project Structure

- `url_fetcher/`
  - `fetcher.go`: Contains the `Fetcher` interface and the `HTTPFetcher` struct with methods to fetch and extract URLs from a web page.
- `web_crawler_synchronous/`
  - `web_crawler_synchronous.go`: Synchronous version of the web crawler.
- `web_crawler_waitgroup/`
  - `web_crawler_waitgroup.go`: Concurrent version of the web crawler using goroutines and `sync.WaitGroup`.

## Dependencies

- `golang.org/x/net/html`: Used for parsing HTML documents.

## How to Run

### Prerequisites

Make sure you have Go installed on your system. You can download and install Go from [the official Go website](https://golang.org/dl/).

### Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/yourusername/your-repository.git
   cd your-repository
   ```

2. Install Dependencies

   ```go mod tidy```

## Running the Synchronous Web Crawler

The synchronous crawler processes URLs sequentially. To run it:

```sh
go run web_crawler_synchronous.go
```

This will start crawling from https://go.dev/ with a depth of 3.

## Running the Concurrent and Parallel Web Crawler

The concurrent crawler uses goroutines to process URLs in parallel. To run it:

```sh
go run web_crawler_synchronous.go
```

This will start crawling from https://go.dev/ with a depth of 3.

# Code Explanation

### `url_fetcher/fetcher.go`

- **Fetcher Interface**: Defines the Fetch method that any fetcher implementation must provide.
- **HTTPFetcher**: Implements the Fetcher interface. It uses an HTTP client to fetch web pages and extract URLs from them.
- **NewHttpFetcher**: Constructs a new HTTPFetcher with a timeout for HTTP requests.

### `web_crawler_synchronous.go`

- **Synchronous Crawl**: This version processes URLs in a single-threaded manner. It recursively fetches and processes URLs until the specified depth is reached.

### `web_crawler_waitgroup.go`

- **Concurrent Crawl**: This version uses goroutines and sync.WaitGroup to process URLs concurrently. It allows for parallel fetching and processing of URLs, improving performance for large crawls.

# Notes

- Adjust the `url` variable and `depth` in both `web_crawler_synchronous.go` and `web_crawler_waitgroup.go` to test different starting points and crawl depths.
- Be cautious with large depth values as they may result in high resource usage and may take significant time to complete. May implement semaphores or channels in future to fix this.
- The concurrent version uses a `sync.Mutex` to protect shared state (cache), ensuring thread safety.

# License
This project is licensed under the MIT License. See the LICENSE file for details.

