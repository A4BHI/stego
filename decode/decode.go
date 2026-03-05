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
	var extlength byte
	for bitsRead < 40 {

		bit := pixels[index] & 1

		currbyte = (currbyte << 1) | bit
		bitcount++
		bitsRead++
		index++

		if bitcount == 8 && bitsRead < 32 {
			byteslice = append(byteslice, currbyte)
			currbyte = 0
			bitcount = 0
		}

		if bitsRead >= 32 && bitcount == 8 {
			extlength = currbyte
			currbyte = 0
			bitcount = 0
		}
	}
	fmt.Println(extlength)
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

	err = os.WriteFile("decoded.txt", sliceofdata, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
