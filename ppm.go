package ppm

import (
	"fmt"
	"math"

	"github.com/Lunarisnia/vecm"
)

type PPM struct {
	Width       int
	Height      int
	AspectRatio float64
	pixelIndex  int
	imageSize   int
}

func CreateImage(width int, height int) *PPM {
	fmt.Println("P3")
	fmt.Printf("%v %v\n", width, height)
	fmt.Println(255)
	return &PPM{
		Width:       width,
		Height:      height,
		AspectRatio: float64(width) / float64(height),
		imageSize:   width * height,
	}
}

func (p *PPM) Draw(color vecm.Vector3f) {
	if p.pixelIndex >= p.imageSize {
		return
	}
	c := vecm.Vector3f{
		X: math.Round(color.X * 255.0),
		Y: math.Round(color.Y * 255.0),
		Z: math.Round(color.Z * 255.0),
	}

	fmt.Printf("%v %v %v\n", c.X, c.Y, c.Z)
	p.pixelIndex++
}
