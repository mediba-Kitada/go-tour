package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

/*
チャネル型(Channel)型は、チャネルオペレータ <- を用いて、値の送受信が出来る通り道
データは、矢印の方向へ流れる
	ch <- v vをチャネルchに送信する
	w := <-ch チャネルchから受信した変数をvへアロケートする
*/
func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// チャネルの生成
	c := make(chan int)
	go sum(s[:len(s)/2], c) // -9+4
	go sum(s[len(s)/2:], c) // 7+2+8

	/*
		通常、チャネルの片方(送信及び受信)の準備が出来るまで、送受信はブロックされる
		これにより、明確なロックや条件変数がなくてもgoroutineの同期を可能にする
	*/
	x, y := <-c, <-c       // receive from c
	fmt.Println(x, y, x+y) // -5 17 12
}
