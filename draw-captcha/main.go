package main

import (
	"image/png"
	"log"
	"os"

	"github.com/afocus/captcha"
)

func main() {
	cap := captcha.New()

	// set font:
	cap.SetFont("comic.ttf")

	img, str := cap.Create(4, captcha.NUM)

	log.Printf("generated captcha: %s", str)

	imgFile, err := os.OpenFile("captcha.png", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	defer imgFile.Close()

	png.Encode(imgFile, img)
	log.Printf("saved captcha image file.")

}
