package main

import "golang.org/x/tour/pic"

type Image struct{}

/*
Imageインターフェースは、imageパッケージで定義されているを実装する
type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x,y int) color.Color
}

color.Colorとcolor.Modelはインターフェースだが、定義済みのcolo.RGBAとcolor.RGBAmodelを使うことで無視出来る
*/

func main() {
	m := Image{}
	pic.ShowImage(m)
}
