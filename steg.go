package main

import (
	"fmt"
	"image"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, _ := image.Decode(file)
	rgba, _ := img.(*image.RGBA)

	pixels := rgba.Pix
	fmt.Println(pixels)

}
