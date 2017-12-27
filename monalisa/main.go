package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	draw(336, 500)
}

func load(filePath string) *image.RGBA {
	i, _ := draw2dimg.LoadFromPngFile(filePath)
	return i.(*image.RGBA)
}

func draw(w int, h int) *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.MoveTo(300, 300) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()

	return dest
}

func diff(a, b *image.RGBA) (d int64) {
	d = 0.0

	return
}
