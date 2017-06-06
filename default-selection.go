package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		// チャネルtickから受信出来る
		case <-tick:
			fmt.Println("tick.")
		// チャネルboomから受信出来る
		case <-boom:
			fmt.Println("BOOM!")
			// 受信できた場合、returnし関数を終了
			return
		/*
			いずれのcaseも準備できていないのであれば、defaultを実行
			goroutineをブロックしない
		*/
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	/*
			   .
			   .
			tick.
			   .
			   .
			tick.
			   .
			   .
		    tick.
			   .
			   .
			tick.
			   .
			   .
			tick.
			BOOM!
	*/
}
