package decode

import (
	"encoding/binary"
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

	length := make([]byte, 4)
	for i := 0; i < 32; i++ {
		for j := 7; j >= 0; j-- {
			length = append(length, (pixels[i]>>j)&1)

		}
	}

	lengthBytes := binary.BigEndian.Uint32(length)

}
