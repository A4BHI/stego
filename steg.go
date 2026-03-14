package main

import (
	_ "image/png"
	"steg/cmd"
)

func main() {
	cmd.Execute()
	// cfg := config.Config{
	// 	InputImage:  "tes.png",
	// 	SecretFile:  "test.txt",
	// 	OutputImage: "encoded",
	// 	DecodedFile: "decoded",
	// 	Password:    "test123",
	// }
	// encode.Encode(&cfg)
	// decode.Decode(&cfg)
}
