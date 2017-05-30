package main

import (
	"fmt"
	"math"
)

// 任意の数値型 MyFloat
type MyFloat float64

// MyFloat型のレシーバを引き取るAbsメソッドを定義
// MyFloat型のAbsメソッド
// レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要がある
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	// MyFloat型のAbsメソッドを実行
	fmt.Println(f.Abs())
}
