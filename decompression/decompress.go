package decompression

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
)

func Decompress(plaintext []byte) []byte {

	r, err := gzip.NewReader(bytes.NewReader(plaintext))
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	decompressedData, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return decompressedData

}
