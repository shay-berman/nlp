package nlp

import (
	"regexp"
	"strings"
)

var (
	wordRe = regexp.MustCompile("[[:alpha:]]+")
)

// Tokenzied returns a slice of tokens found in text -> this is the way go docs will see it well.
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		// TODO: stem
		token := strings.ToLower(w)
		tokens = append(tokens, token)
	}
	return tokens
}
