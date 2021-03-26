package main

import (
	"image/color"
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"
)

const (
	BLACK_AND_WHITE = iota
	HUE
)

func (d *pixelData) computeColor(maxIterations uint64, colorMode int) {
	switch colorMode {
	case BLACK_AND_WHITE:
		var c uint8
		if d.iterations == maxIterations {
			c = 0
		} else {
			c = uint8(
				math.Round(
					float64(maxIterations) * math.Sqrt(
						float64(d.iterations)/float64(maxIterations),
					),
				),
			)
		}
		d.color = color.RGBA{
			R: c, G: c, B: c, A: 0xff,
		}
		break
	case HUE:
		hue := 360 * float64(d.iterations) / float64(maxIterations)
		saturation := 1
		value := 1
		if d.iterations == maxIterations {
			value = 0
		}

		d.color = colorful.Hsv(hue, float64(saturation), float64(value))
	}
}

func (d *pixelData) computeColorContinuous(maxIterations uint64, colorMode int, lengthSqr float64, escape float64) {
	// var logz float64 = math.Logb(length)
	// var logr float64 = math.Logb(escape)

	// var adjustedIterations float64 = float64(d.iterations) - math.Logb(logz/logr)

	var log_zn float64 = math.Log(lengthSqr) / 2
	var nu float64 = math.Log2(log_zn / math.Log(escape))

	var adjustedIterations float64 = float64(d.iterations) - nu

	// print(d.iterations)
	// print(" ")
	// print(lengthSqr)
	// print(" ")
	// print(log_zn)
	// print(" ")
	// print(nu)
	// print(" ")
	// print(adjustedIterations)
	// println()

	switch colorMode {
	case BLACK_AND_WHITE:
		var c uint8

		if d.iterations == maxIterations {
			c = 0
		} else {

			c = uint8(
				math.Round(
					float64(maxIterations) * math.Sqrt(
						adjustedIterations/float64(maxIterations),
					),
				),
			)
		}
		d.color = color.RGBA{
			R: c, G: c, B: c, A: 0xff,
		}
		break
	case HUE:
		hue := 222 + 360*(adjustedIterations/float64(maxIterations))
		saturation := 0.77
		value := 0.85
		if d.iterations == maxIterations {
			value = 0
		}

		d.color = colorful.Hsv(hue, float64(saturation), float64(value))
	}
}
