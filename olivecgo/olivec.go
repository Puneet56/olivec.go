package olivecgo

import (
	"bytes"
	"fmt"
	"image/color"
	"os"
)

type Pixel color.RGBA

func FillColor(pixels []Pixel, width, height int, color Pixel) {
	for i := 0; i < width*height; i++ {
		pixels[i] = color
	}
}

func WritePixelsToPPM(filename string, pixels []Pixel, width, height int) error {
	b := bytes.Buffer{}

	header := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
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

func BlendColor(base, top Pixel) Pixel {
	a1 := float64(base.A) / 255
	a2 := float64(top.A) / 255

	a := a1 + a2*(1-a1)

	r1 := float64(base.R)
	r2 := float64(top.R)

	g1 := float64(base.G)
	g2 := float64(top.G)

	b1 := float64(base.B)
	b2 := float64(top.B)

	r := (a*r1*a1*(1-a2) + r2*a2) / a
	g := (a*g1*a1*(1-a2) + g2*a2) / a
	b := (a*b1*a1*(1-a2) + b2*a2) / a

	return Pixel{uint8(r), uint8(g), uint8(b), uint8(a * 255)}
}

func FillRect(pixels []Pixel, width, height, posX, posY, rw, rh int, color Pixel) {
	for dx := 0; dx < rw; dx++ {
		for dy := 0; dy < rh; dy++ {
			x := posX + dx
			y := posY + dy
			pixels[y*width+x] = BlendColor(pixels[y*width+x], color)
		}
	}
}

func FillCircle(pixels []Pixel, width, height, posX, posY, r int, color Pixel) {
	x1 := posX - r
	y1 := posY - r
	x2 := posX + r
	y2 := posY + r

	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			dx := x - posX
			dy := y - posY
			if dx*dx+dy*dy <= r*r {
				pixels[y*width+x] = BlendColor(pixels[y*width+x], color)
			}
		}
	}
}
