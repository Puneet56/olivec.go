package main

import (
	"bytes"
	"fmt"
	"image/color"
	"os"
)

const (
	WIDTH  = 800
	HEIGHT = 600
)

func main() {
	pixels := make([]color.RGBA, WIDTH*HEIGHT)
	fillColor(pixels, WIDTH, HEIGHT, color.RGBA{255, 0, 0, 0})
	writePixelsToPPM("output.ppm", pixels, WIDTH, HEIGHT)
}

func fillColor(pixels []color.RGBA, width, height int, color color.RGBA) {
	for i := 0; i < WIDTH*HEIGHT; i++ {
		pixels[i] = color
	}
}

func writePixelsToPPM(filename string, pixels []color.RGBA, width, height int) error {
	b := bytes.Buffer{}

	header := fmt.Sprintf("P6\n%d %d\n255\n", WIDTH, HEIGHT)
	if _, err := b.Write([]byte(header)); err != nil {
		return err
	}

	for _, p := range pixels {
		pb := []byte{p.R, p.G, p.B}
		if _, err := b.Write(pb); err != nil {
			return err
		}
	}

	if err := os.WriteFile(filename, b.Bytes(), 0644); err != nil {
		return err
	}

	fmt.Printf("%s PPM file written successfully\n", filename)
	return nil
}
