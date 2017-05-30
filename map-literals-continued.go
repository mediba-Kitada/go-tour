package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// mapの型が単純な場合、リテラルの要素から推測するので、make関数が不要となる
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
