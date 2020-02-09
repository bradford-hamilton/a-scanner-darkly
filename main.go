package main

import (
	"fmt"

	"github.com/bradford-hamilton/a-scanner-darkly/cmd"
)

const aScannerDarkly = `
eeeee      eeeee eeee eeeee eeeee eeeee eeee eeeee       eeeee eeeee eeeee  e   e  e   e    e
8   8      8   " 8  8 8   8 8   8 8   8 8    8   8       8   8 8   8 8   8  8   8  8   8    8
8eee8 eeee 8eeee 8e   8eee8 8e  8 8e  8 8eee 8eee8e eeee 8e  8 8eee8 8eee8e 8eee8e 8e  8eeee8
88  8         88 88   88  8 88  8 88  8 88   88   8      88  8 88  8 88   8 88   8 88    88
88  8      8ee88 88e8 88  8 88  8 88  8 88ee 88   8      88ee8 88  8 88   8 88   8 88eee 88
`

// printlnGreen simply prints a green string
func printlnGreen(text string) {
	fmt.Printf("\033[32m%s\033[0m\n", text)
}

func main() {
	printlnGreen(aScannerDarkly)
	cmd.Execute()
}
