package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func Button(onClick func(), content func() *imdraw.IMDraw) *imdraw.IMDraw {
	// record onClick for Button region
	r := content()
	return r
}

type BoxElement struct {
	X, Y          float64
	Width, Height float64
}

func Box(x, y float64, width, height float64) BoxElement {
	return BoxElement{
		X: x, Y: y,
		Width:  width,
		Height: height,
	}
}
func (b *BoxElement) Render() *imdraw.IMDraw {
	imd := imdraw.New(nil)

	x := b.X
	y := b.Y
	width := b.Width
	height := b.Height

	foreground := pixel.RGB(1, 0, 0)
	background := pixel.RGB(1, 1, 0)
	borderThickness := 5.5

	imd.Color(background)
	imd.Push(pixel.V(x, y),
		pixel.V(x+width, y),
		pixel.V(x+width, y+height),
		pixel.V(x, y+height))
	imd.Polygon(0)

	imd.Color(foreground)
	imd.Push(pixel.V(x, y),
		pixel.V(x+width, y),
		pixel.V(x+width, y+height),
		pixel.V(x, y+height),
		pixel.V(x, y))
	imd.EndShape(imdraw.NoEndShape)
	imd.Line(borderThickness)

	return imd
}
