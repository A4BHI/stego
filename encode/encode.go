package encode

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"path/filepath"
	"steg/compression"
	"steg/encryption"

	"log"
	"os"
)

func Encode(imgfile string, secretfile string) {
	file, err := os.Open(imgfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
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

	data := compression.Compress(secretfile)
	ciphertext, nonce, salt := encryption.Encrypt(data, "test")
	index := 0
	length := len(ciphertext)
	fmt.Println("Encoded length:", len(ciphertext))
	ext := filepath.Ext(secretfile)
	extdata := []byte(ext)

	extbytes := byte(len(extdata))
	lengthBytes := make([]byte, 4)

	binary.BigEndian.PutUint32(lengthBytes, uint32(length))

	payload := append(lengthBytes, extbytes)
	payload = append(payload, extdata...)
	payload = append(payload, salt...)
	payload = append(payload, nonce...)
	payload = append(payload, ciphertext...)
	totalbits := len(payload) * 8

	if totalbits > len(pixels) {
		log.Fatal("Not enough space in image.")
	}
	for i := 0; i < len(payload); i++ {
		for j := 7; j >= 0; j-- {

			bit := (payload[i] >> j) & 1

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
