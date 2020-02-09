package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Retrieve the currently installed asd CLI version",
	Long:  "Simply run `asd version` to get your current asd CLI version",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("The version command does not take any arguments")
			os.Exit(1)
		}

		fmt.Println(Version)
	},
}
