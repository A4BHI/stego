package encode

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/png"
	_ "image/png"

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
	for i := 3; i < len(pixels); i += 4 {
		pixels[i] = 255
	}

	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal("Error Reading Data: ", err)
	}

	index := 0
	length := len(data)
	fmt.Println("Encoded length:", len(data))
	lengthBytes := make([]byte, 4)

	binary.BigEndian.PutUint32(lengthBytes, uint32(length))
	//fmt.Println(lengthBytes)
	// lenindex := len(lengthBytes)
	//fmt.Println(data)

	// for k := 0; k < len(lengthBytes); k++ {
	// 	for l := 7; l >= 0; l-- {
	// 		bit := (lengthBytes[k] >> l) & 1
	// 		pixels[index] = (pixels[index] & 254) | bit
	// 		index++
	// 	}
	// } inefficient

	payload := append(lengthBytes, data...) //more better way instead of two loops
	totalbits := len(payload) * 8

	if totalbits > len(pixels) {
		log.Fatal("Not enough space in image.")
	}
	for i := 0; i < len(payload); i++ {
		for j := 7; j >= 0; j-- {

			bit := (payload[i] >> j) & 1

			for index%4 == 3 {
				index++
			}
			pixels[index] = (pixels[index] & 254) | bit
			index++
		}

	}
	OutFile, err := os.Create("Output.png")
	if err != nil {
		log.Fatal(err)
	}
	index = 0
	bitsPrinted := 0
	fmt.Print("Encode bits: ")

	for bitsPrinted < 32 {

		if index%4 == 3 {
			index++
			continue
		}

		fmt.Print(pixels[index] & 1)
		bitsPrinted++
		index++
	}
	fmt.Println()
	err = png.Encode(OutFile, rgba)
	if err != nil {
		log.Fatal(err)
	}
}
