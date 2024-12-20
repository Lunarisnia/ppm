package ppm

import (
	"fmt"

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
		X: color.X * 255.0,
		Y: color.Y * 255.0,
		Z: color.Z * 255.0,
	}

	fmt.Printf("%v %v %v\n", int(c.X), int(c.Y), int(c.Z))
	p.pixelIndex++
}
