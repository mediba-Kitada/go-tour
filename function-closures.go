package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	// クロージャを返却
	return func(x int) int {
		// クロージャとしては、sum変数にバインドされている
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			// クロージャ(関数値)に値を渡す
			pos(i),
			neg(-2*1),
		)
		/*
			sum関数にバインドされているので、インクリメントされた数値で処理される
				0 -2
				1 -4
				3 -6
				6 -8
				10 -10
				15 -12
				21 -14
				28 -16
				36 -18
				45 -20
		*/
	}
}
