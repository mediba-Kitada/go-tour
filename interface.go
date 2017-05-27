package main

import (
	"fmt"
	"math"
)

// interface型は、メソッドのシグニチャの集まりで定義する
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// メソッドの集まりを実装した値をinterface型の変数aに持たせることが出来る
	a = f
	a = &v

	// ポインタレシーバなので、変数をレシーバが受け取ってくれない(インターフェイスを実装していないということになる)
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
