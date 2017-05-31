package main

import "fmt"

func do(i interface{}) {
	// 型switchの宣言は、型アサーションi.(T)と同じ構文を持つが、特定の型Tはキーワードtypeに置き換えられる
	// 変数vはiによって保持される値を保持
	switch v := i.(type) {
	case int:
		// iはint型を保持しているか
		fmt.Printf("Twice %v is %v\n", v, v*2)
		// iはstring型を保持しているか
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		// 一致する型を保持していない場合、変数vは同じインターフェイス型で値はiとなる
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {
	do(21)      // Twice 21 is 42
	do("hello") // "hello" is 5 bytes long
	do(true)    // I don't know about type bool!
}
