package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Initialize all additional commands, flags, etc
func init() {}

var rootCmd = &cobra.Command{
	Use:   "asd",
	Short: "asd (A Scanner Darkly) is a simple port scanner written in Go",
	Long: `
asd (A Scanner Darkly) is a simple and fast port scanner. 
`,
	Args: cobra.MinimumNArgs(1),
	Run:  runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Unknown command. Try `asd help` for more information")
}

// Execute runs a user's command. On error, it will print the error and cause
// the program to exit with status code 1
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
