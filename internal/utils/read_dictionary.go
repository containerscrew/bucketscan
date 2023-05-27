package utils

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

func ReadFuzzFile(logger *slog.Logger, dictionaryPath string) []string {
	var words []string
	readFile, err := os.Open(dictionaryPath)

	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}

	defer readFile.Close()

	return words
}
