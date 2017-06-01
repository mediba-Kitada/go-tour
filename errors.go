package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// error型はfmt.Sringerに似た組み込みのインターフェース
/*
type error interface {
	Error() string
}
*/

// Errorインターフェース Error関数を実装
// fmt.Stringerと同様、fmtパッケージは変数を文字列で出力する際にerrorインターフェースを確認する
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// 関数はerror型を返却するパターンがある。呼び出し元は、エラー(error型)がnilかどうかを確認することでエラーをハンドルする
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
