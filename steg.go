package main

import (
	_ "image/png"
	"steg/decode"
)

func main() {

	//encode.Encode("tes.png")
	decode.Decode("Output.png")
}
