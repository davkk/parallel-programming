package main

import (
	"image"
	"image/color"
	"image/png"
	"mandelbrot-go/internal"
	"os"
)

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
			r, g, b := hsb.HSBToRGB(hue, 1.0, 1.0)
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
