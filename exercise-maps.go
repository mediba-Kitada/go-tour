package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	wc := make(map[string]int)
	for _, k := range strings.Fields(s) {
		if i, ok := wc[k]; ok {
			wc[k] = i + 1
		} else {
			wc[k] = 1
		}
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
