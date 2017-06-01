package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
	}
	return ""
}

func Sqrt(x float64) (float64, error) {

	// ErrNegativeSqrt型は、Error関数を実装しているので、error型を返却することが出来る
	if e := ErrNegativeSqrt(x); e.Error() != "" {
		return x, e
	}

	// 開始点 z
	z := 1.0
	// xの平方根を10回のループで求める
	for i := 0; i < 10; i++ {
		z -= (math.Pow(z, 2) - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
