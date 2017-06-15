package main

import (
	"fmt"
)

type Fetcher interface {
	Fetch(url string) (body string, url []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth < 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("http://golang.ort/", 4, fetcher)
}
