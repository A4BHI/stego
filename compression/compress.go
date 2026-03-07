package compression

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func Compress(file string) []byte {
	data, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	if err != nil {
		fmt.Println(err)
	}
	var b bytes.Buffer

	gzip, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		log.Fatal(err)
	}
	defer gzip.Close()

	io.Copy(gzip, data)

	return b.Bytes()

}
