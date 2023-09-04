package main

import (
	"errors"
	"fmt"
	"sync"
)

type SafeCrawlerCache struct {
    mutex sync.Mutex
    urlCache map[string]error
}

func (crawlerCache* SafeCrawlerCache) IsInCache(url string) bool {
    crawlerCache.mutex.Lock()
    defer crawlerCache.mutex.Unlock()

    _, ok := crawlerCache.urlCache[url]

    return ok
}

func (crawlerCache* SafeCrawlerCache) AddUrl(url string, err error ) {
    crawlerCache.mutex.Lock()
    defer crawlerCache.mutex.Unlock()

    crawlerCache.urlCache[url] = err
}

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }
    if crawlerCache.IsInCache(url) {
        fmt.Printf("<- Done with %v, already fetched! \n", url)
        return
    }
    crawlerCache.AddUrl(url, errors.New("Loading..."))
    body, urls, err := fetcher.Fetch(url)
    crawlerCache.AddUrl(url, err)

    if err != nil {
        fmt.Printf("<- Error on %v: %v\n", url, err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    done := make(chan bool)
    for i, u := range urls {
        fmt.Printf("-> Crawing child %v/%v of %v : %v\n", i+1, len(urls), url, u)
        go func(url string) {
            Crawl(url, depth-1, fetcher)
            done <- true
        }(u)
    }

    for i, url := range urls {
        fmt.Printf("<-[%v] Waiting for child %v/%v \n", url, i+1, len(urls))
        <-done
    }

   fmt.Printf("<-Done with %v\n", url) 
}

func main() {
    Crawl("https://golang.org/", 4, fetcher)

    fmt.Println("Fetching Stats\n------------------")

    for url, err := range crawlerCache.urlCache {
        if err != nil {
            fmt.Printf("%v failed: %v\n", url, err)
        } else {
            fmt.Printf("%v, was fetched\n", url)
        }
    }
}

type fakeFetcher map[string]* fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }

    return "", nil, fmt.Errorf("not found %s", url)
}

var crawlerCache = SafeCrawlerCache{
    urlCache: make(map[string]error),
}

var fetcher = fakeFetcher {
    "https://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string {
            "https://golang.org/pkg/",
            "https://golang.org/cmd/",
        },
    },
    "https://golang.org/pkg/": &fakeResult {
        "Packages",
        []string {
            "https://golang.org/",
            "https://golang.org/cmd/",
            "https://golang.org/pkg/fmt/",
            "https://golang.org/pkg/os/",
        },
    },
    "https://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string {
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
    "https://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string {
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
}
