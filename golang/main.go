package main

import (
	"image"
	"image/color"
	"image/png"
	"mandelbrot-go/internal"
	"os"
)

type Params struct {
	size          int
	minRe         float64
	maxRe         float64
	minIm         float64
	maxIm         float64
	maxIterations int
}

func mandelbrot(img *image.RGBA, y0 int, params Params) {
	width := img.Bounds().Dx()
	pxRe := (params.maxRe - params.minRe) / float64((width - 1))
	pxIm := (params.maxIm - params.minIm) / float64((width - 1))

	for y := y0; y < min(y0+width, params.size); y++ {
		cIm := params.maxIm - float64(y)*pxIm

		for x := 0; x < width; x++ {
			cRe := params.minRe + float64(x)*pxRe

			zRe := 0.0
			zIm := 0.0
			iterations := 0

			for zRe*zRe+zIm*zIm <= 4 && iterations < params.maxIterations {
				nextRe := zRe*zRe - zIm*zIm + cRe
				nextIm := 2*zRe*zIm + cIm

				zRe = nextRe
				zIm = nextIm

				iterations++
			}

			hue := float64(iterations) / float64(params.maxIterations) * 2.5
			saturation := 1.0
			brightness := 1.0
			if iterations == params.maxIterations {
				brightness = 0.0
			}

			r, g, b := hsb.HSBToRGB(hue, saturation, brightness)
			img.SetRGBA(x, y, color.RGBA{r, g, b, 255})
		}
	}

}

func main() {
	size := 8196
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	mandelbrot(img, 0, Params{
		size:          size,
		minRe:         -2.0,
		maxRe:         1.0,
		minIm:         -1.5,
		maxIm:         1.5,
		maxIterations: 200,
	})

	file, _ := os.Create("output.png")
	defer file.Close()

	png.Encode(file, img)
}
