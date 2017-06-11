package main

/*
チャネルがgoroutineとコミュケーションするための素晴らしい方法だが、コミュケーションが不要な場合もある
あるgoroutineが変数へのコンフリクトを避けるためだけに一時的に変数へのアクセスをできるようにしたい場合は、排他制御を実装する必要がある
排他制御(mutual exclusion)を表現するデータ構造をmutex(ミューテックス)と呼ぶ
*/
import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	// 排他制御の実施
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Unlockの保証のため、deferを利用
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
