package main

import (
	"fmt"
	"math"
)

// Vertex型
type Vertex struct {
	X, Y float64
}

// レシーバ引数を関数に取らせることで、型にメソッドを定義出来る
// Vertex型にAbsメソッドを定義
// ポインタレシーバと比較して、変数レシーバと呼ばれる
// 変数レシーバでは、レシーバのVertex変数をコピーを操作する
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// Vertex型の変数vを代入
	v := Vertex{3, 4}
	// Vertex型のAbsメソッドを実行
	fmt.Println(v.Abs())
}
