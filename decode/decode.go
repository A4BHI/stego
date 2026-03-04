package decode

import (
	"encoding/binary"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

func Decode(targetfile string) {
	inputimg, err := os.Open(targetfile)
	if err != nil {
		log.Fatal(err)
	}

	defer inputimg.Close()

	img, _, err := image.Decode(inputimg)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	rgba := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	pixels := rgba.Pix

	index := 0
	bitsRead := 0

	var byteslice []byte
	var currbyte byte
	bitcount := 0

	for bitsRead < 32 {

		bit := pixels[index] & 1

		currbyte = (currbyte << 1) | bit
		bitcount++
		bitsRead++
		index++

		if bitcount == 8 {
			byteslice = append(byteslice, currbyte)
			currbyte = 0
			bitcount = 0
		}
	}

	bitsRead = 0
	var sliceofdata []byte
	lengthBytes := binary.BigEndian.Uint32(byteslice)
	databits := int(lengthBytes) * 8
	for bitsRead < databits {
		bit := pixels[index] & 1
		currbyte = (currbyte << 1) | bit
		index++
		bitcount++
		bitsRead++

		if bitcount == 8 {
			sliceofdata = append(sliceofdata, currbyte)
			currbyte = 0
			bitcount = 0
		}
	}

	fmt.Println(sliceofdata)

	err = os.WriteFile("decoded.txt", sliceofdata, 0644)
	if err != nil {
		log.Fatal(err)
	}

	index = 0
	bitsPrinted := 0

	fmt.Print("Decode bits: ")

	for bitsPrinted < 32 {

		fmt.Print(pixels[index] & 1)
		bitsPrinted++
		index++
	}
	fmt.Println()
	// lengthBytes := binary.BigEndian.Uint32(byteslice)
	fmt.Println(lengthBytes)

}
