package main

import (
	_ "image/png"
	"steg/decode"
	"steg/encode"
)

func main() {

	encode.Encode("tes.png", "test.txt")
	decode.Decode("Output.png")
}
