package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
ポインタレシーバを使うべき理由
- メソッドがレシーバが指す先の変数を変更するため
- メソッドの呼び出し毎に変数のコピーを避けるため

一般的に変数レシーバまたは、ポインタレシーバのどちからですべてのメソッドを与え、混在させるべきではない
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 構造体のプロパティを変更する必要はないが、ポインタレシーバとする
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
