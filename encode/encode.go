package encode

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func Encode(targetfile string) {
	file, err := os.Open(targetfile)
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
	length := len(data)

	lengthBytes := make([]byte, 4)

	binary.BigEndian.PutUint32(lengthBytes, uint32(length))
	fmt.Println(lengthBytes)
	// lenindex := len(lengthBytes)
	//fmt.Println(data)

	for k := 0; k < len(lengthBytes); k++ {
		for l := 7; l >= 0; l-- {
			bit := (lengthBytes[k] >> l) & 1
			pixels[index] = (pixels[index] & 254) | bit
			index++
		}
	}

	for i := 0; i < len(data); i++ {
		for j := 7; j >= 0; j-- {

			bit := (data[i] >> j) & 1
			pixels[index] = (pixels[index] & 254) | bit
			index++
		}

	}
	OutFile, err := os.Create("Output.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(OutFile, rgba)
	if err != nil {
		log.Fatal(err)
	}
}
