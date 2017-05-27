package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.X * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)        // ポインタレシーバを持つ、Scaleは変数vのステートメントであるが、利便性のため(&v).Scale(2)として解釈される
	ScaleFunc(&v, 10) // ScaleFuncは、ポインタ型の引数を持つので、ポインタを引数に渡す

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
