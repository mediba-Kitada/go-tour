package main

import (
	"fmt"
)

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
func main() {
	var s []int
	printSlice("s", s) //s len=0 cap=0 []

	// ゼロ値なスライスに対して要素を追加
	// 長さと容量が追加されている
	s = append(s, 0)
	printSlice("appended s", s) // appended s len=1 cap=1 [0]

	// 一度に複数要素を追加出来る
	s = append(s, 2, 3, 4)
	printSlice("appended s take02", s) // appended s take02 len=4 cap=4 [0 2 3 4]

}
