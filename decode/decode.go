package decode

import (
	"encoding/binary"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"steg/config"
	"steg/decryption"
)

type FileMetaData struct {
	Datalength int
	Extlength  int
	Extname    string
	CurrIndex  int
}

func Decode(cfg *config.Config) {
	inputimg, err := os.Open(cfg.OutputImage)
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
	filemetadata := getDatalenandExtLen(pixels)
	getFileExtension(&filemetadata, pixels)
	salt, nonce := GetNonceandSalt(&filemetadata, pixels)
	ciphertext, filename := DecodeData(&filemetadata, pixels)
	plaintext := decryption.Decrypt(ciphertext, salt, nonce)

}

func getDatalenandExtLen(pixels []uint8) FileMetaData {
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

		if bitcount == 8 && bitsRead <= 32 {
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

	lengthBytes := binary.BigEndian.Uint32(byteslice)

	return FileMetaData{
		Datalength: int(lengthBytes) * 8,
		Extlength:  int(extlength) * 8,
		CurrIndex:  index,
	}
}

func getFileExtension(filemetadata *FileMetaData, pixels []uint8) {
	bitsRead := 0
	var currbyte byte
	bitcount := 0
	var sliceofext []byte
	for bitsRead < filemetadata.Extlength {
		bit := pixels[filemetadata.CurrIndex] & 1
		currbyte = (currbyte << 1) | bit
		filemetadata.CurrIndex++
		bitcount++
		bitsRead++

		if bitcount == 8 {
			sliceofext = append(sliceofext, currbyte)
			currbyte = 0
			bitcount = 0
		}
	}

	filemetadata.Extname = string(sliceofext)

}
func GetNonceandSalt(filemetadata *FileMetaData, pixels []uint8) ([]byte, []byte) {
	bitsRead := 0

	var saltslice []byte
	var nonceslice []byte

	var currbyte byte
	bitcount := 0

	for bitsRead < 224 { //salt * nonce means 16 * 12 = 224 bits we have to iterate till 224

		bit := pixels[filemetadata.CurrIndex] & 1

		currbyte = (currbyte << 1) | bit
		bitcount++
		bitsRead++
		filemetadata.CurrIndex++

		if bitcount == 8 && bitsRead <= 128 { //first 128 bits = salt
			saltslice = append(saltslice, currbyte)
			currbyte = 0
			bitcount = 0
		}

		if bitsRead > 128 && bitcount == 8 { //rest of the bits till 224 belongs to salt
			nonceslice = append(nonceslice, currbyte)
			currbyte = 0
			bitcount = 0
		}
	}

	return saltslice, nonceslice
}
func DecodeData(filemetadata *FileMetaData, pixels []uint8) ([]byte, string) {
	bitsRead := 0
	var sliceofdata []byte
	var currbyte byte
	bitcount := 0
	for bitsRead < filemetadata.Datalength {
		bit := pixels[filemetadata.CurrIndex] & 1
		currbyte = (currbyte << 1) | bit
		filemetadata.CurrIndex++
		bitcount++

		if bitcount == 8 {
			sliceofdata = append(sliceofdata, currbyte)
			currbyte = 0
			bitcount = 0
		}
	}
	filename := "decoded" + filemetadata.Extname
	// err := os.WriteFile(filename, sliceofdata, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return sliceofdata, filename

}
