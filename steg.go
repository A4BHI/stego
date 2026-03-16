package main

import (
	"fmt"
	_ "image/png"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/a4bhi/stego/ui"
)

func main() {
	// cmd.Execute()
	// cfg := config.Config{
	// 	InputImage:  "tes.png",
	// 	SecretFile:  "test.txt",
	// 	OutputImage: "encoded",
	// 	DecodedFile: "decoded",
	// 	Password:    "test123",
	// }
	// stego.Encode(&cfg)
	// stego.Decode(&cfg)

	p := tea.NewProgram(ui.InitialModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
