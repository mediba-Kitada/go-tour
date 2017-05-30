package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// メソッドが変数レシーバの場合、呼び出し時に変数またはポイントのいずれかをレシーバとして受け取ることが出来る
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // ポインタもレシーバが受け取ってくれる (*p).Abs()として解釈される
	fmt.Println(AbsFunc(*p))
}
