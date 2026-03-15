package main

import (
	_ "image/png"
	"steg/config"
	"steg/stego"
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
	stego.Encode(&cfg)
	stego.Decode(&cfg)
}
