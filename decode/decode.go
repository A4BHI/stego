package decode

import (
	"encoding/binary"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

type FileMetaData struct {
	Datalength int
	Extlength  int
	Extname    string
	CurrIndex  int
}

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
	filemetadata := getDatalenandExtLen(pixels)
	getFileExtension(&filemetadata, pixels)
	DecodeData(&filemetadata, pixels)

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

func DecodeData(filemetadata *FileMetaData, pixels []uint8) {
	bitsRead := 0
	var sliceofdata []byte
	var currbyte byte
	bitcount := 0
	for bitsRead < filemetadata.Datalength {
		bit := pixels[filemetadata.CurrIndex] & 1
		currbyte = (currbyte << 1) | bit
		filemetadata.CurrIndex++
		bitcount++
		bitsRead++

		if bitcount == 8 {
			sliceofdata = append(sliceofdata, currbyte)
			currbyte = 0
			bitcount = 0
		}
	}
	filename := "decoded" + filemetadata.Extname
	err := os.WriteFile(filename, sliceofdata, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
