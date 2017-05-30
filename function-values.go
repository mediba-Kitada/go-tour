package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	// 関数値 戻り値として利用出来る
	return fn(3, 4)
}

func main() {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13

	// 関数値 引数として利用出来る
	fmt.Println(compute(hypot))    // 5
	fmt.Println(compute(math.Pow)) // 81
}
