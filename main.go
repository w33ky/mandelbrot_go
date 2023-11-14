package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"strconv"
)

type vec = [2]float64

func main() {
	if len(os.Args) < 3 {
		panic("you have to pass x and y")
	}

	resX, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	resY, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	fmt.Println("resX:", resX, "resY:", resY)

	iteration := 50
	if len(os.Args) >= 4 {
		iteration, err = strconv.Atoi(os.Args[3])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("iterations:", iteration)

	posX := float64(0)
	posY := float64(0)
	if len(os.Args) >= 6 {
		posX, err = strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			panic(err)
		}
		posY, err = strconv.ParseFloat(os.Args[5], 64)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("posx:", posX, "posY:", posY)

	scaleX := float64(1)
	scaleY := float64(1)
	if len(os.Args) >= 8 {
		scaleX, err = strconv.ParseFloat(os.Args[6], 64)
		if err != nil {
			panic(err)
		}
		scaleY, err = strconv.ParseFloat(os.Args[7], 64)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("scaleX:", scaleX, "scaleY:", scaleY)

	img := image.NewRGBA(image.Rect(0, 0, resX, resY))

	for x := 0; x < resX; x++ {
		for y := 0; y < resY; y++ {
			px := calcPos(x, resX)
			py := calcPos(y, resY)

			v := vec{px*scaleX + posX, py*scaleY + posY}
			m := mandelbrot(iteration, v)

			red := uint8((255 / iteration) * m)
			green := 0
			if m > (iteration / 2) {
				green = ((255 / iteration) * (m - (iteration / 2)))
			}
			blue := 0

			img.Set(x, y, color.RGBA{red, uint8(green), uint8(blue), 255})
		}
	}

	f, err := os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, img, nil); err != nil {
		log.Printf("failed to encode: %v", err)
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
