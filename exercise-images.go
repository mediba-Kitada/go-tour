package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

/*
Imageインターフェースは、imageパッケージで定義されているを実装する
type Image interface {
	ColorModel() color.Model
	Bounds() Rectangle
	At(x,y int) color.Color
}
*/
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

// ひとつの色を返却
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x + y), 255, 3, uint8(x)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
