import (
	"strings"
	"unicode"
)

func countWordOccurrences(text string) map[string]int {
	wordCounts := make(map[string]int)

	words := strings.FieldsFunc(text, func(c rune) bool {
		return !unicode.IsLetter(c)
	})

	for _, word := range words {
		word = strings.ToLower(word)

		wordCounts[word]++
	}

	return wordCounts
}

