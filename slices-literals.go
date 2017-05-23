package main

import (
	"fmt"
)

func main() {

	// スライスのリテラル
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// 無名の構造体を要素とするスライス
	// 構造体のプロパティの初期設定も同時に実施
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s)
}
