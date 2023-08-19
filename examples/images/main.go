package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width  int
	Height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	value := uint8(x * y)

	return color.RGBA{value, value, 255, 255}
}

func main() {
	m := Image{
		Width:  256,
		Height: 256,
	}
	pic.ShowImage(m)
}
