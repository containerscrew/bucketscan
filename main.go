package main

import (
	"fmt"

	"github.com/containerscrew/bucketscan/cmd"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

// printBanner will print an ascii banner with colors
func printBanner() {
	templ := `{{ .AnsiColor.Green  }} {{ .Title "bucketscan" "" 2 }}{{ .AnsiColor.Default }}
   Author: github.com/containerscrew
   Now: {{ .Now "Monday, 2 Jan 2006" }}`
	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	fmt.Printf("\n\n")
}

func main() {
	//printBanner()
	cmd.Execute()
}
