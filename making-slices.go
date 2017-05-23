package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
func main() {
	// make関数でゼロ化された配列を割り当て、その配列を指すスライスを返却
	a := make([]int, 5)
	printSlice("a", a) // len=5 cap=5 [0 0 0 0 0]

	// make関数の第三引数でスライスの容量を指定
	b := make([]int, 0, 5)
	printSlice("b", b) // b len=0 cap=5 []

	// スライスbの長さを拡張
	c := b[:2]
	printSlice("c", c) // c len=2 cap=5 [0 0]

	// スライスcをスライス
	d := c[2:5]
	printSlice("d", d) // d len=3 cap=3 [0 0 0]
}
