package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("s: %s i: %d\n", s, i)
	}
}

// goroutine(ゴルーチン)は、Goのランタイムに管理される軽量なスレッド
// goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要がある
func main() {
	// say()及び"world"の評価は、実行元(current)で実行される
	// say()の実行は、新しいgoroutineで実行される
	go say("world")
	say("hello")
	/*
		s: hello i: 0 引数sは、currentで評価され、say()関数の実行は、currentで実行される
		s: world i: 0 引数sは、currentで評価され、say()関数の実行は、新しいgoroutineで実行される
		s: world i: 1
		s: hello i: 1
		s: hello i: 2
		s: world i: 2
		s: world i: 3
		s: hello i: 3
		s: hello i: 4
		s: world i: 4

			world
			hello
			hello
			world
			world
			hello
			hello
			world
			world
			hello
	*/
}
