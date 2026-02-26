package decode

import (
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

	for y := bounds.Min.Y;i<

}
