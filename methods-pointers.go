package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
  ポインタレシーバ
  Vertex型への構文*Vertexがあることを意味する
  レシーバが指す変数を変更出来る
  レシーバ自身を更新することが多いため、変数レシーバより一般的(変数レシーバは、レシーバの指す変数を変更出来ない)
*/
func (v *Vertex) Scala(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// *intのようにポインタ自身を取ることが出来ない
	v.Scala(10)
	fmt.Println(v.Abs())
}
