package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	// dyの長さ、容量を持ち、[]uint8型の要素を持つスライスを作成
	pic := make([][]uint8, dy)

	// dyの長さ分ループする
	for y := range pic {
		// それぞれの要素に dxの長さ、容量を持ちuint8型の要素を持つスライスを作成
		pic[y] = make([]uint8, dx)
		// dxの長さ分ループする
		for x := range pic[y] {
			// それぞれの要素にuint8型の数値を代入
			pic[y][x] = uint8((x + y) / 2)
		}
	}

	return pic
}

func main() {
	// todo Pic関数はどの様にコールされている??
	pic.Show(Pic)
}
