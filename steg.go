package main

import (
	_ "image/png"
	"steg/config"
	"steg/decode"
	"steg/encode"
)

func main() {
	// cmd.Execute()
	cfg := config.Config{
		InputImage:  "tes.png",
		SecretFile:  "test.txt",
		OutputImage: "encoded",
		DecodedFile: "decoded",
		Password:    "test123",
	}
	encode.Encode(&cfg)
	decode.Decode(&cfg)
}
