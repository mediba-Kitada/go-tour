package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// インターフェイスの値iがstring型を保持し、基となる具体的な値を変数sに代入
	s := i.(string)
	fmt.Println(s)

	// インターフェイスの値が特定の型を保持しているかどうかをテストするため、型アサーションは2つの値(元になる値とアサーションが成功したかどうかを報告するブール)を返却
	s, ok := i.(string)
	fmt.Println(s, ok) // sがstring型を保持していれば、変数okはtrue

	// 型アサーションを実施すれば、iがfloat64型を保持していなければ、変数fはfloat64型のゼロ値となりpanicは発生しない
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// 型アサーションを実施せず、iがfloat64型を保持していなければ、panicが発生する
	f = i.(float64)
	fmt.Println(f)
}
