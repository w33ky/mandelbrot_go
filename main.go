package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

type vec = [2]float64

func main() {
	var resX int
	var resY int
	var iterations int
	var posX float64
	var posY float64
	var scaleX float64
	var scaleY float64
	var colorPreset string
	var useJpeg bool
	var printHelp bool

	flag.IntVar(&resX, "resX", 800, "width of the image in pixels")
	flag.IntVar(&resY, "resY", 600, "height of the image in pixels")
	flag.IntVar(&iterations, "iterations", 50, "maximum number of iterations to calculate")
	flag.Float64Var(&posX, "posX", 0, "x positions of the center of the resulting image")
	flag.Float64Var(&posY, "posY", 0, "y positions of the center of the resulting image")
	flag.Float64Var(&scaleX, "scaleX", 1, "x scale of the resulting image")
	flag.Float64Var(&scaleY, "scaleY", 1, "y scale of the resulting image")
	flag.StringVar(&colorPreset, "colorPreset", "default", "choose a color preset: default, red")
	flag.BoolVar(&useJpeg, "jpeg", false, "write jpeg imagte file")
	flag.BoolVar(&printHelp, "help", false, "print help")

	flag.Parse()

	if printHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Println("resX:", resX, "resY:", resY)
	fmt.Println("iterations:", iterations)
	fmt.Println("posx:", posX, "posY:", posY)
	fmt.Println("scaleX:", scaleX, "scaleY:", scaleY)
	fmt.Println("colorPreset:", colorPreset)
	fmt.Println("write file as jpeg:", useJpeg)

	img := image.NewRGBA(image.Rect(0, 0, resX, resY))

	for x := 0; x < resX; x++ {
		for y := 0; y < resY; y++ {
			px := calcPos(x, resX)
			py := calcPos(y, resY)

			v := vec{px*scaleX + posX, py*scaleY + posY}
			m := mandelbrot(iterations, v)

			img.Set(x, y, calcColor(m, iterations, colorPreset))
		}
	}

	if useJpeg {
		f, err := os.Create("img.jpg")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err = jpeg.Encode(f, img, &jpeg.Options{Quality: 95}); err != nil {
			log.Printf("failed to encode: %v", err)
		}
	} else {
		f, err := os.Create("img.png")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err = png.Encode(f, img); err != nil {
			log.Printf("failed to encode: %v", err)
		}
	}
}

func mandelbrot(n int, z vec) int {
	var z1 vec
	z2 := vec{0, 0}
	for i := 0; i < n; i++ {
		z1 = z2
		z2 = vecAdd(vecQuad(z1), z)
		if vecAbs(z2) >= 2 {
			return i
		}
	}
	return 0
}

func calcColor(m int, iterations int, colorPreset string) color.RGBA {
	switch colorPreset {
	case "red":
		red := uint8((255 / iterations) * m)
		green := 0
		if m > (iterations / 2) {
			green = ((255 / iterations) * (m - (iterations / 2)))
		}
		blue := 0
		return color.RGBA{red, uint8(green), uint8(blue), 255}
	}

	return color.RGBA{uint8((255 / iterations) * m), uint8((255 / iterations) * m), uint8((255 / iterations) * m), 255}
}

func vecAdd(a vec, b vec) vec {
	return vec{a[0] + b[0], a[1] + b[1]}
}

func vecQuad(z vec) vec {
	return vec{z[0]*z[0] - z[1]*z[1], z[0]*z[1] + z[1]*z[0]}
}

func vecAbs(z vec) float64 {
	return math.Sqrt(z[0]*z[0] + z[1]*z[1])
}

func calcPos(n int, res int) float64 {
	return (2/float64(res))*float64(n) - 1
}
