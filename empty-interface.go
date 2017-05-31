package main

import "fmt"

func main() {
	// 空のインターフェース(The empty interface) ゼロ個のメソッドを指定されたインターフェース
	// 空のインターフェースは、未知の型の値を扱うコードで利用される
	// e.g. fmt.Printf()は、interface{}型の任意の数の引数を受け取る
	var i interface{}
	describe(i) //(<nil>, <nil>)

	// 空のインターフェースは任意の型の値を保持出来る(全ての方は、少なくともゼロ個のメソッドを実装している)
	i = 42
	describe(i) //(42, int)

	i = "hello"
	describe(i) //(hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
