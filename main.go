package main

import (
	"github.com/Puneet56/olivec.go/olivecgo"
)

const (
	WIDTH  = 800
	HEIGHT = 600
)

func main() {
	pixels := make([]olivecgo.Pixel, WIDTH*HEIGHT)
	olivecgo.FillColor(pixels, WIDTH, HEIGHT, olivecgo.Pixel{0, 0, 0, 255})
	olivecgo.FillCircle(pixels, WIDTH, HEIGHT, WIDTH/2, HEIGHT/2, 100, olivecgo.Pixel{255, 255, 255, 150})
	olivecgo.FillRect(pixels, WIDTH, HEIGHT, WIDTH/2, HEIGHT/2, 100, 100, olivecgo.Pixel{255, 255, 255, 200})
	olivecgo.WritePixelsToTerminal("title", pixels, WIDTH, HEIGHT)
}
