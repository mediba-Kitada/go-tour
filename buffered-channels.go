package main

import "fmt"

func main() {
	// チャネルはバッファとして利用出来る
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	/*
		バッファが詰まった(バッファの長さが足りない場合)時は、チャネルへの送信をブロックする
		ch <- 3
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
			/Users/kitatuba/go/src/github.com/mediba-kitada/go-tour/buffered-channels.go:9 +0xd4
		exit status 2
	*/
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	/*
		//ch <- 2
		バッファが空(バッファの中身が無い)の時は、チャネルの受信をブロックする
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan receive]:
		main.main()
			/Users/kitatuba/go/src/github.com/mediba-kitada/go-tour/buffered-channels.go:20 +0x126
		exit status 2
	*/

}
