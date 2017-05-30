package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// インターフェイスを実装することを明示的に宣言する必要は無い(implementsキーワードは必要無い)
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// インターフェイスI型の変数にT型を代入
	var i I = T{"hello"}
	// 実装したインターフェイスを呼び出し
	i.M()
}
