package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}

	// プロパティの値を指定して初期化
	v2 = Vertex{X: 1}

	v3 = Vertex{}

	// &オペレータを指定すると構造体へのポインタを返却する
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3) // {1 2} &{1 2} {1 0} {0 0}
}
