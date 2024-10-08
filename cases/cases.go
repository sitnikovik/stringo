package cases

import (
	"strings"
	"unicode"
)

const (
	// NormalCase describes that string built in normal case
	NormalCase StringCase = iota + 1
	// SnakeCase describes that string built in snake_case
	SnakeCase
	// CamelCase describes that string built in camelCase
	CamelCase
	// PascalCase describes that string built in PascalCase
	PascalCase
	// KebabCase describes that string built in kebab-case
	KebabCase
	// ScreamingSnakeCase describes that string built in SCREAMING_SNAKE_CASE
	ScreamingSnakeCase
	// TrainCase describes that string built in Train-Case
	TrainCase
	// DotCase describes that string built in dot.case
	DotCase
)

type StringCase int8

// MatchPascalCase defines if the string matches the PascalCase
func MatchPascalCase(s string) bool {
	return pascalCaseRE.MatchString(s)
}

// DefineStringCase returns case type for the string
func DefineStringCase(s string) StringCase {
	if MatchKebabCase(s) {
		return KebabCase
	}
	if MatchSnakeCase(s) {
		return SnakeCase
	}
	if MatchPascalCase(s) {
		return PascalCase
	}
	if MatchCamelCase(s) {
		return CamelCase
	}
	if MatchScreamingSnakeCase(s) {
		return ScreamingSnakeCase
	}
	if MatchTrainCase(s) {
		return TrainCase
	}
	if MatchDotCase(s) {
		return DotCase
	}

	return NormalCase
}

// ToUpperFirst converts the first character of a string to uppercase
func ToUpperFirst(s string) string {
	return caseFirstFunc(s, unicode.ToUpper)
}

// ToLowerFirst converts the first character of a string to lowercase
func ToLowerFirst(s string) string {
	return caseFirstFunc(s, unicode.ToLower)
}

// caseFirstFunc converts the first character of a string using the provided function
func caseFirstFunc(s string, f func(rune) rune) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	runes[0] = f(runes[0])

	if len(runes) == 1 {
		return string(runes)
	}

	return string(runes[0]) + string(runes[1:])
}

// SplitToWords splits a string into words
func SplitToWords(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

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
func split(s, sep string, sepcond func(r rune, idx int) bool, f func(r rune, idx int) rune) string {
	if s == "" {
		return ""
	}

	sb := strings.Builder{}
	for i, r := range s {
		if sep != "" && sepcond(r, i) {
			sb.WriteString(sep)
		}
		sb.WriteString(string(f(r, i)))
	}

	return sb.String()
}

// join joins strings with separator and apply function to each string
func join(ss []string, sep string, f func(s string, idx int) string) string {
	if len(ss) == 0 {
		return ""
	}

	if f == nil {
		return strings.Join(ss, sep)
	}

	sb := strings.Builder{}
	for i, s := range ss {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(f(s, i))
	}

	return sb.String()
}

// mapWordsAndJoin splits string to words, do the callback for each word
// and join results into new string with separator
func mapWordsAndJoin(s, sep string, f func(s string, idx int) string) string {
	words := SplitToWords(s)
	for i, word := range words {
		words[i] = f(word, i)
	}

	return strings.Join(words, sep)
}
