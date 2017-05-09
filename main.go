package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Pixel Rocks!",
		Bounds:    pixel.R(0, 0, 1024, 768),
		Resizable: true,
		VSync:     true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	b := Box(50, 50, 150, 100)
	imd := b.Render()

	for !win.Closed() {
		win.Clear(colornames.Aquamarine)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
