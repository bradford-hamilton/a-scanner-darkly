package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version constant that represents the current build version
const Version = "v.0.1.0"

// Ports will be hydrated with it's value if a user runs the scan cmd with the
// flag --ports (or -p)
var Ports string

// Initialize all additional commands, flags, etc
func init() {
	// Attach all commands
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(versionCmd)

	// Hydrate scanCmd flag variables with the (if any) user input
	scanCmd.Flags().StringVarP(
		&Ports,
		"ports",
		"p",
		"1-1024",
		"choose port numbers to scan",
	)
}

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
