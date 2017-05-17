package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Element interface {
	Bounds() pixel.Rect
	Render() interface{}
}

var backend = make(map[Element]func())

type BoxElement struct {
	X, Y          float64
	Width, Height float64
}

func (b *BoxElement) Bounds() pixel.Rect {
	return pixel.Rect{
		Min: pixel.V(b.X, b.Y),
		Max: pixel.V(b.X+b.Width, b.Y+b.Height),
	}
}

func Box(x, y float64, width, height float64) *BoxElement {
	return &BoxElement{
		X: x, Y: y,
		Width:  width,
		Height: height,
	}
}

func (b *BoxElement) Render() interface{} {
	imd := imdraw.New(nil)

	foreground := pixel.RGB(1, 0, 0)
	background := pixel.RGB(1, 1, 0)
	borderThickness := 5.5

	x := b.X
	y := b.Y
	width := b.Width
	height := b.Height

	imd.Color(foreground)
	imd.Push(pixel.V(x-borderThickness, y-borderThickness),
		pixel.V(x+width+borderThickness, y-borderThickness),
		pixel.V(x+width+borderThickness, y+height+borderThickness),
		pixel.V(x-borderThickness, y+height+borderThickness))
	imd.Polygon(0)

	imd.Color(background)
	imd.Push(pixel.V(x, y),
		pixel.V(x+width, y),
		pixel.V(x+width, y+height),
		pixel.V(x, y+height))
	imd.Polygon(0)

	return imd
}

func Button(onClick func(), content Element) *ButtonElement {
	// record onClick for Button region
	button := &ButtonElement{
		Content: content,
		OnClick: onClick,
	}
	backend[button] = onClick
	return button
}

type ButtonElement struct {
	Content Element
	OnClick func()
}

func (b *ButtonElement) Bounds() pixel.Rect {
	return b.Content.Bounds()
}

func (b *ButtonElement) Render() interface{} {
	return b.Content.Render()
}
