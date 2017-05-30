package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("value:", m["answer"])

	m["answer"] = 48
	fmt.Println("value:", m["answer"])

	delete(m, "answer")
	fmt.Println("value:", m["answer"])

	// 存在しないキーを指定した場合、値はその型のゼロ値となる
	v, ok := m["answer"]
	fmt.Println("value:", v, "present?", ok)
}
