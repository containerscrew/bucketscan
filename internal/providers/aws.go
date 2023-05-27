package providers

import (
	"fmt"

	"github.com/containerscrew/bucketscan/internal/utils"
	"golang.org/x/exp/slog"
)

func AWSMutations(keywords []string, quickScan bool, logger *slog.Logger, dictionaryPath string) []string {
	var mutations []string

	if quickScan {
		for _, keyword := range keywords {
			mutations = append(mutations, fmt.Sprintf("https://%s.%s", keyword, S3URL))
		}
		return mutations
	}

	// If quickScan not selected, then create mutatiosn using your keywords and fuzz.txt file or your custom dictionary
	words := utils.ReadFuzzFile(logger, dictionaryPath)

	for _, word := range words {
		for _, keyword := range keywords {
			// Appends
			mutations = append(mutations, fmt.Sprintf("https://%s%s.%s", word, keyword, S3URL))
			mutations = append(mutations, fmt.Sprintf("https://%s.%s.%s", word, keyword, S3URL))
			mutations = append(mutations, fmt.Sprintf("https://%s-%s.%s", word, keyword, S3URL))

			// Prepends
			mutations = append(mutations, fmt.Sprintf("https://%s%s.%s", keyword, word, S3URL))
			mutations = append(mutations, fmt.Sprintf("https://%s.%s.%s", keyword, word, S3URL))
			mutations = append(mutations, fmt.Sprintf("https://%s-%s.%s", keyword, word, S3URL))
		}
	}

	return mutations
}
