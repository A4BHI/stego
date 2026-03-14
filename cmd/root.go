package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testcmd = cobra.Command{
	Use:   "Test",
	Short: "testing cobra",
	Example: `
			Nothing nigga just testing
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TESTTTTT PASSEEDDDD")
	},
}

func Execute() error {
	return testcmd.Execute()
}
