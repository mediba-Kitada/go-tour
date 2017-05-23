package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// 構造体Vertexオペランドへのポインタを&オペレータを用いて代入
	p := &v

	// ポインタを通して構造体のプロパティにアクセス、値を設定
	p.X = 1e9
	fmt.Println(v)

	// 本来は、*p型であるが、上記の通り省略出来る
	fmt.Println((*p).X)
}
