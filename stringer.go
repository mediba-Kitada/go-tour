package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

/* 変数を文字列で出力するためのインターフェース
* fmtパッケージその他多くのパッケージで標準的に提供されている
 */
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)                     // Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
	fmt.Printf("value: %v type:%T", a, a) // value: Arthur Dent (42 years) type:main.Person
}
