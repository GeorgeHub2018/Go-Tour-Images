package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	w := 300
	h := 255
	return image.Rect(0, 0, w, h)
}

func (i Image) At(x, y int) color.Color {
	if x <= 100 {
		return color.RGBA{uint8(x), uint8(y), 255, 255}
	} else if x <= 200 {
		return color.RGBA{uint8(x + y/2), uint8(y), 255, 255}
	} else {
		return color.RGBA{uint8(x + y/2), uint8(y - x/2), 255, 255}
	}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}