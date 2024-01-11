package utils

import (
	"fmt"

	"golang.org/x/exp/slog"
)

var (
	mutations, awsEndpoints, gcpEndpoints, azureEndpoints []string
)

type Mutations struct {
	Mutations []string
}

func CreateProviderMutations(provider string, mutations []string) {
	switch provider {
	case "all":

	case "gcp":
	// run gcp
	case "aws":

	}
}

func (m *Mutations) QuickScanMutations(keywords []string) Mutations {
	for _, keyword := range keywords {
		m.Mutations = append(mutations, fmt.Sprintf("%s.%s", keyword))
	}
	return *m
}

func CreateMutations(keywords []string, dictionaryPath string, logger *slog.Logger) []string {
	words := ReadDictionaryFile(logger, dictionaryPath)

	for _, word := range words {
		for _, keyword := range keywords {
			// Appends
			mutations = append(mutations, fmt.Sprintf("%s%s", word, keyword))
			mutations = append(mutations, fmt.Sprintf("%s.%s", word, keyword))
			mutations = append(mutations, fmt.Sprintf("%s-%s", word, keyword))

			// Prepends
			mutations = append(mutations, fmt.Sprintf("%s%s", keyword, word))
			mutations = append(mutations, fmt.Sprintf("%s.%s", keyword, word))
			mutations = append(mutations, fmt.Sprintf("%s-%s", keyword, word))
		}
	}

	return mutations
}
