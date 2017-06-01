package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {

	// 開始点 z
	z = 1.0
	// xの平方根を10回のループで求める
	for i := 0; i < 10; i++ {
		z -= (math.Pow(z, 2) - x) / (2 * z)
	}

	return
}

func main() {
	fmt.Println(Sqrt(float64(2)))
	fmt.Println(Sqrt(float64(3)))
}
