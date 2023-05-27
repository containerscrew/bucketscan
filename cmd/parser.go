package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

type Arguments struct {
	Workers        *int
	Keywords       *[]string
	QuickScan      *bool
	DictionaryPath *string
	Nameserver     *string
	LogLevel       *string
}

func ParseArgs() Arguments {
	var arguments Arguments
	parser := argparse.NewParser("bucketscan", "Simple Golang bucket brute force scanner (AWS|GCP|AZURE)")

	arguments.Workers = parser.Int("w", "workers", &argparse.Options{Required: false, Help: "Number of workers (threads)", Default: 5})
	arguments.Keywords = parser.StringList("k", "keyword", &argparse.Options{Required: true, Help: "Keyword for url mutations"})
	arguments.QuickScan = parser.Flag("q", "quick-scan", &argparse.Options{Required: false, Default: false, Help: "Quick scan, do not create mutations from fuzz.txt file"})
	arguments.DictionaryPath = parser.String("d", "dictionary", &argparse.Options{Required: true, Help: "Dictionary path"})
	arguments.Nameserver = parser.String("n", "nameserver", &argparse.Options{Required: false, Help: "Custom nameserver", Default: "8.8.8.8 "})
	arguments.LogLevel = parser.Selector("l", "log-level", []string{"info", "debug"}, &argparse.Options{Required: false, Default: "info", Help: "Log level of the application"})
	err := parser.Parse(os.Args)

	if err != nil {
		log.Fatal(fmt.Println(parser.Usage(err)))
	}

	return arguments
}
