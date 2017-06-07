package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	for {
		select {
		case ch <- t.Left.Value:
		case ch <- t.Right.Value:
		default:
			ch <- t.Value
		}
	}
}

/*
func Same(t1, t2 *tree.Tree) bool {
}
*/

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
