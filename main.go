package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
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

	imd := imdraw.New(nil)

	imd.Color(pixel.RGB(1, 0, 0))
	imd.Push(pixel.V(200, 100))
	imd.Color(pixel.RGB(0, 1, 0))
	imd.Push(pixel.V(800, 100))
	imd.Color(pixel.RGB(0, 0, 1))
	imd.Push(pixel.V(500, 700))
	imd.Polygon(0)

	imd.Color(colornames.Blueviolet)
	imd.EndShape(imdraw.RoundEndShape)
	imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	imd.EndShape(imdraw.SharpEndShape)
	imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	imd.Line(30)

	imd.Color(colornames.Limegreen)
	imd.Push(pixel.V(500, 325))
	imd.Circle(300, 50)
	imd.Color(colornames.Navy)
	imd.Push(pixel.V(200, 500), pixel.V(800, 500))
	imd.Ellipse(pixel.V(120, 80), 0)

	imd.Color(colornames.Red)
	imd.EndShape(imdraw.RoundEndShape)
	imd.Push(pixel.V(500, 350))
	imd.CircleArc(150, -math.Pi, 0, 30)

	for !win.Closed() {
		win.Clear(colornames.Aquamarine)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
