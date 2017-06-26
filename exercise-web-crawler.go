package main

import (
	"fmt"
	"sync"
	"time"
)

// fetcher is a populated fakeFetcher.
var (
	fetcher = fakeFetcher{
		"http://golang.org/": &fakeResult{
			"The Go Programing Language",
			[]string{"http://golang.org/pkg/", "http://golang.org/cmd/"},
		},
		"http://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{"http://golang.org/", "http://golang.org/cmd/", "http://golang.org/pkg/fmt/"},
		},
		"http://golang.org/pkg/fmt/": &fakeResult{
			"Package fmt",
			[]string{"http://golang.org/", "http://golang.org/pkg/"},
		},
		"http://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{"http://golang.org/", "http://golang.org/pkg/"},
		},
	}
	fetchedUrls = &FetchedUrls{
		v: make(map[string]int),
	}
)

// Fetcher Fetcherインターフェース
type Fetcher interface {
	// Fetch returns the body or URL and a slice of URLs found on the page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {

	fetchedUrls.Save(url)

	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found:%s", url)
}

// FetchedUrls フェッチ済みのURLとフェッチ回数を保持する構造体
type FetchedUrls struct {
	v   map[string]int
	mux sync.Mutex
}

// Save フェッチ済みURLをフェッチ回数を保存する関数
func (fu *FetchedUrls) Save(url string) {
	fu.mux.Lock()
	defer fu.mux.Unlock()
	fu.v[url]++
}

// Load フェッチ済みURLをキーにフェッチ回数を取得する関数
func (fu *FetchedUrls) Load(url string) int {
	fu.mux.Lock()
	defer fu.mux.Unlock()
	return fu.v[url]
}

// Crawl uses fetcher to recursively crawl pages starting with url,to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Don't fetch the same URL twice.
	// TODO: Fetch URLs in parallel.
	// This implementation dose't do either:

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if fetchedUrls.Load(url) >= 2 {
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	time.Sleep(time.Second)
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}
