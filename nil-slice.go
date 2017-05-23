package main

import (
	"fmt"
)

func main() {
	/*
		スライスのゼロ値はnil
		nilスライスは、0の長さと容量を保持している
		元となる配列も持っていない
	*/
	var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!")
	}
}
