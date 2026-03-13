package main

import (
	_ "image/png"
	"steg/config"
	"steg/decode"
	"steg/encode"
)

func main() {
	cfg := config.Config{
		InputImage:  "tes.png",
		SecretFile:  "test.txt",
		OutputImage: "decoded",
		DecodedFile: "decoded",
		Password:    "test123",
	}
	encode.Encode("tes.png", "test.txt")
	decode.Decode("Output.png")
}
