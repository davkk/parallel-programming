package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func HSBToRGB(h, s, v float64) (uint8, uint8, uint8) {
	if s == 0 {
		rgb := uint8(v * 255)
		return rgb, rgb, rgb
	}

	h = h * 6
	if h == 6 {
		h = 0
	}
	i := math.Floor(h)
	v1 := v * (1 - s)
	v2 := v * (1 - s*(h-i))
	v3 := v * (1 - s*(1-(h-i)))

	var r, g, b float64
	switch int(i) {
	case 0:
		r, g, b = v, v3, v1
	case 1:
		r, g, b = v2, v, v1
	case 2:
		r, g, b = v1, v, v3
	case 3:
		r, g, b = v1, v2, v
	case 4:
		r, g, b = v3, v1, v
	default:
		r, g, b = v, v1, v2
	}

	return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}

type Mandelbrot struct {
	minRe   float64
	maxRe   float64
	minIm   float64
	maxIm   float64
	maxIter int
}

func (mb *Mandelbrot) calculate(x, y float64) color.Color {
	c := complex(x, y)
	z := complex(0, 0)
	for i := 0; i < mb.maxIter; i++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			hue := float64(i) / float64(mb.maxIter) * 2.5
			r, g, b := HSBToRGB(hue, 1.0, 1.0)
			return color.RGBA{r, g, b, 255}
		}
		z = z*z + c
	}
	return color.Black
}

func (mb *Mandelbrot) render(img *image.RGBA, size int) {
	dx := (mb.maxRe - mb.minRe) / float64(size)
	dy := (mb.maxIm - mb.minIm) / float64(size)
	for py := 0; py < size; py++ {
		y := mb.minIm + float64(py)*dy
		for px := 0; px < size; px++ {
			x := mb.minRe + float64(px)*dx
			col := mb.calculate(x, y)
			img.Set(px, py, col)
		}
	}
}

func main() {
	size := 2048
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	mandelbrot := Mandelbrot{
		minRe:   -2.0,
		maxRe:   1.0,
		minIm:   -1.5,
		maxIm:   1.5,
		maxIter: 200,
	}

	mandelbrot.render(img, size)

	file, _ := os.Create("output.png")
	defer file.Close()

	png.Encode(file, img)
}
