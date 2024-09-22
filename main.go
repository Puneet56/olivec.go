package main

import "olivec.go/olivecgo"

const (
	WIDTH  = 800
	HEIGHT = 600
)

func main() {
	pixels := make([]olivecgo.Pixel, WIDTH*HEIGHT)
	olivecgo.FillColor(pixels, WIDTH, HEIGHT, olivecgo.Pixel{255, 255, 255, 255})
	olivecgo.FillCircle(pixels, WIDTH, HEIGHT, WIDTH/2, HEIGHT/2, 150, olivecgo.Pixel{255, 0, 0, 150})
	olivecgo.FillRect(pixels, WIDTH, HEIGHT, WIDTH/2, HEIGHT/2, 100, 100, olivecgo.Pixel{0, 255, 0, 200})

	olivecgo.WritePixelsToPPM("output.ppm", pixels, WIDTH, HEIGHT)
}
