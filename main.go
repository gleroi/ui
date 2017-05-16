package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "ui",
		Bounds:    pixel.R(0, 0, 1024, 768),
		Resizable: true,
		VSync:     true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	b := Button(func() {
		log.Print("button clicked!")
	}, Box(50, 50, 150, 100))

	var result interface{} = b
	done := false
	for !done {
		switch t := result.(type) {
		case Element:
			result = t.Render()
		case *imdraw.IMDraw:
			done = true
		}
	}
	imd := result.(*imdraw.IMDraw)

	for !win.Closed() {
		if win.JustReleased(pixelgl.MouseButtonLeft) {
			clickPosition := win.MousePosition()
			for elt, handler := range backend {
				if elt.Bounds().Contains(clickPosition) {
					handler()
				}
			}
		}

		win.Clear(colornames.White)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
