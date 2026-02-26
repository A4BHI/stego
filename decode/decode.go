package decode

import (
	"encoding/binary"
	"fmt"
	"image"
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
	rgba := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	pixels := rgba.Pix
	var byteslice []byte
	var currbyte byte
	bitcount := 0
	length := make([]byte, 4)
	for i := 0; i < 32; i++ {
		bit := pixels[i] & 1

		currbyte = (currbyte << 1) | bit
		bitcount++

		if bitcount == 8 {
			byteslice = append(byteslice, currbyte)
			bitcount = 0
			currbyte = 0
		}

	}

	lengthBytes := binary.BigEndian.Uint32(length)
	fmt.Println(lengthBytes)

}
