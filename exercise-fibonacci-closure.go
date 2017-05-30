package main

import (
	"fmt"
)

// fibonacci is a function that return a function that returns an int
func fibonacci() func() int {

	c := 0
	p := make(map[int]int)
	return func() int {
		switch c {
		case 0:
			// 初期条件1
			p[0] = 0
			c++
			return 0
		case 1:
			// 初期条件2
			p[1] = 1
			c++
			return 1
		default:
			p[c] = p[c-2] + p[c-1]
			c++
			return p[c-1]
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
