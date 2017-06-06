package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		/*
			selectステートメントは、goroutineを複数の通信操作で待たせる
			複数あるcaseのいずれかが準備できるようになるまでブロックし、準備が出来たcaseを実行する
			もし、複数のcaseの準備が出来ている場合、caseはランダムに選択される
		*/
		select {
		// 変数xの値をチャネルcに送信
		case c <- x:
			x, y = y, x+y
		// チャネルquitから受信
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
