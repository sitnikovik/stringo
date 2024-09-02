package cases

import (
	"strings"
)

// rejoin splits provided string with fromSep, maps each splitted
// part of the string with callback and join resulted with toSep
func rejoin(s, fromSep, toSep string, f func(s string, idx int) string) string {
	if s == "" {
		return ""
	}

	words := strings.Split(s, fromSep)
	for i, word := range words {
		words[i] = f(word, i)
	}

	return strings.Join(words, toSep)
}

// split maps provided string with separator on its condition function returns true
// and do the callback for each rune of string
func split(s, sep string, sepcond func(r rune, idx int) bool, f func(r rune) rune) string {
	if s == "" {
		return ""
	}

	sb := strings.Builder{}
	for i, r := range s {
		if sep != "" && sepcond(r, i) {
			sb.WriteString(sep)
		}
		sb.WriteString(string(f(r)))
	}

	return sb.String()
}
