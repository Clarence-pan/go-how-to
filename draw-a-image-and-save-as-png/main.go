package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {

	log.Printf("Creating the image...")
	img := image.NewRGBA(image.Rect(0, 0, 300, 400))

	log.Printf("Drawing the image...")
	imgBounds := img.Bounds()

	for i := imgBounds.Min.Y; i < imgBounds.Max.Y; i++ {
		log.Printf("Drawing line %d", i)
		for j := imgBounds.Min.X; j < imgBounds.Max.X; j++ {
			img.SetRGBA(j, i, color.RGBA{
				R: uint8(i % 255),
				G: uint8(max(0, i-255) % 255),
				B: uint8(j % 255),
				A: 255,
			})
		}
	}

	log.Printf("Saving the image...")
	// create a file to save the image
	f, err := os.OpenFile("test.png", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// write the image as png
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

	log.Printf("File written")

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
