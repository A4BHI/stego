package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

func main() {
	file, err := os.Open("tes.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	pixels := rgba.Pix

	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal("Error Reading Data: ", err)
	}
	index := 0

	fmt.Println(data)
	for i := 0; i < len(data); i++ {
		for j := 7; j >= 0; j-- {

			bit := (data[i] >> j) & 1
			pixels[index] = (pixels[index] & 254) | bit
			index++
		}

	}

	// for i := 0; i < len(pixels); i++ {
	// 	fmt.Println(i, ":", pixels[i], "\n")
	// }

	//fmt.Println(pixels[6289403])

	//fmt.Println(pixels)
	//bitwise niggaaaa
}
