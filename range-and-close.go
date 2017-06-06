package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	/*
		送り手のチャネルcだけ、クローズする
		closeしたチャネルに送信するとパニック(panic)が発生する
		チャネルはファイルと異なり、closeする必要が無い
		closeするのは、これ以上値が来ないことを受け手が知る必要があるときのみ(rangeループの終了等)
	*/
	close(c)
}

func main() {
	c := make(chan int, 10)
	/*
		fibonacci関数の引数の評価(cap(c)及びc)は、current goroutineで実施
		新しいgoroutineでfibonacci関数を実行
	*/
	go fibonacci(cap(c), c)
	// sendされた回数だけループする
	for i := range c {
		// 受け手のチャネルcは、クローズしない
		fmt.Println(i)
		/*
			0
			1
			1
			2
			3
			5
			8
			13
			21
			34
		*/
	}
}
