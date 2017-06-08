package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {

	/* 二分木を探索し、チャネルに送信する処理を関数型の変数に代入する
	探索方法は、片方の部分木を調べ、根を調べ、次いで反対の部分木を調べる通りがけ順(in-order)を採用している
	*/
	var walker func(*tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}

	// 二分木を探索
	walker(t)

	// rangeでループしたいので、チャネルをクローズ
	close(ch)
	/*
		fmt.Println(t) // ((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)
		selectを用いて簡潔に記述したかったが、ニルポでチャネルに送信出来ず
		そもそも探索方法が誤っていた...
		for {
			select {
			case ch <- t.Right.Value:
				t = t.Right
			case ch <- t.Left.Value:
				t = t.Left
			default:
				ch <- t.Value
				break
			}
		}
	*/
}

func Same(t1, t2 *tree.Tree) bool {

	ch01 := make(chan int, 10)
	go Walk(t1, ch01)

	ch02 := make(chan int, 10)
	go Walk(t2, ch02)

	s01 := make([]int, 10)
	s02 := make([]int, 10)

	for i := range ch01 {
		s01 = append(s01, i)
	}

	for i := range ch02 {
		s02 = append(s02, i)
	}

	for k, v := range s01 {
		if v != s02[k] {
			return false
		}
	}

	return true
}

func main() {

	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Same!")
	}

	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("Not same!")
	}
}
