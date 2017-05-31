package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

/*
インターフェースの実装の具体的な値がnilの場合、メソッドはnilをレシーバとして呼び出す
具体的な値がnilではない場合、それ自体がnilではない
*/
func (t *T) M() {
	// golangでは、nilをレシーバとして呼び出されても適切に処理するメソッドを記述するのが一般的
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	i = t
	describe(i) // (<nil>, *main.T)
	i.M()       // <nil>

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T)
	i.M()       // hello
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
