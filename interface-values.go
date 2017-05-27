package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	// インターフェイスの値は、特定の規程になる具体的な型の値を保持する
	i = &T{"Hello"}
	// インターフェイスの値は、値と具体的な型のタプルと捉える事が出来る
	describe(i) // (&{Hello},*main.T)
	// インターフェイスの値のメソッドを呼び出すと、その規定型の同じ名前のメソッドが実行される
	i.M() // Hello

	i = F(math.Pi)
	describe(i) // (3.141592653589793,main.F)
	i.M()       // 3.141592653589793
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}
