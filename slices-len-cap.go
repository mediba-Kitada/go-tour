package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// スライスsの長さを0にスライス(slice)
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// スライスsの長さを4に設定(expand)
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// スライスsの頭2つの要素を削除(drop)
	// 容量も削らている
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]

	fmt.Println(s) // [5 7]

	// 削除したスライスに対しては、拡張は出来ない
	s = s[:6] //panic: runtime error: slice bounds out of range
	printSlice(s)
}
