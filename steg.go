package main

import (
	"image"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	image.Decode()
	rgba := image.NewRGBA(image.Bounds())

}
