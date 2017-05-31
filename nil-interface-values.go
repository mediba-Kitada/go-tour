package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i) // nilインターフェースの値は、値も具体的な型も保持しない (<nil>, <nil>)
	i.M()       // 呼び出す具体的なメソッドを示す型がインターフェースのタプル内に存在しないため、nilインターフェースメソッドを呼び出すとランタイムエラーになる
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x1088b9a]
	*/

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
