package stringo

import (
	"strings"
	"unicode"
)

// ToPascalCase converts a string to PascalCase
func ToPascalCase(s string) string {
	words := SplitToWords(s)
	for i, word := range words {
		words[i] = ToUpperFirst(word)
	}

	return strings.Join(words, "")
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	for i, word := range words {
		w := strings.ToLower(word)
		if i > 0 {
			w = ToUpperFirst(w)
		}
		words[i] = w
	}

	return strings.Join(words, "")
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "_")
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "-")
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

// FromSnakeToPascalCase converts a snake_case string to PascalCase
func FromSnakeToPascalCase(s string) string {
	words := SplitToWords(s)
	for i, word := range words {
		words[i] = ToUpperFirst((word))
	}

	return strings.Join(words, "")
}
