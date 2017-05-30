package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// mapのゼロ値は、nilなのでmakeにより指定された型の初期化を行う必要がある
	// make関数は、初期化され、利用できるようになったmapを返却する
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
