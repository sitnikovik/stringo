package cases

import (
	"regexp"
	"strings"
	"unicode"
)

var camelCaseRE = regexp.MustCompile("^[a-z]+(?:[A-Z][a-z]+)*$")

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	if c := DefineStringCase(s); c != NormalCase {
		switch c {
		case CamelCase:
			return s
		case PascalCase:
			return FromPascalToCamelCase(s)
		case KebabCase:
			return FromKebabToCamelCase(s)
		case SnakeCase:
			return FromSnakeToCamelCase(s)
		case ScreamingSnakeCase:
			return FromScreamingSnakeToCamelCase(s)
		case TrainCase:
			return FromTrainToCamelCase(s)
		}
	}

	words := SplitToWords(s)
	for i, word := range words {
		w := strings.ToLower(word)
		if i > 0 {
			w = ToUpperFirst(w)
		}
		words[i] = w
	}

	return strings.Join(words, "")
}

// MatchCamelCase defines if the string matches the camelCase
func MatchCamelCase(s string) bool {
	return camelCaseRE.MatchString(s)
}

// FromCamelToSnakeCase converts a camelCase string to snake_case.
// Keep in mind that it skips spaces cause of these does not match camelCase.
func FromCamelToSnakeCase(s string) string {
	return split(
		s,
		"_",
		func(r rune, idx int) bool {
			return unicode.IsUpper(r) && idx > 0
		},
		func(r rune, _ int) rune {
			return unicode.ToLower(r)
		},
	)
}

// FromCamelToSnakeCase converts a camelCase string to kebab-case.
// Keep in mind that it skips spaces cause of these does not match camelCase.
func FromCamelToKebabCase(s string) string {
	return split(
		s,
		"-",
		func(r rune, idx int) bool {
			return unicode.IsUpper(r) && idx > 0
		},
		func(r rune, _ int) rune {
			return unicode.ToLower(r)
		},
	)
}

// FromCamelToPascalCase converts a camelCase string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match camelCase.
func FromCamelToPascalCase(s string) string {
	if s == "" {
		return ""
	}

	return ToUpperFirst(s)
}

// FromCamelToScreamingSnakeCase converts a camelCase string to SCREAMING_SNAKE_CASE.
// Keep in mind that it skips spaces cause of these does not match camelCase.
func FromCamelToScreamingSnakeCase(s string) string {
	return split(
		s,
		"_",
		func(r rune, idx int) bool {
			return unicode.IsUpper(r) && idx > 0
		},
		func(r rune, _ int) rune {
			return unicode.ToUpper(r)
		},
	)
}

// FromCamelToTrainCase converts a camelCase string to Train-Case.
// Keep in mind that it skips spaces cause of these does not match camelCase.
func FromCamelToTrainCase(s string) string {
	return FromPascalToTrainCase(s) // It is the same in this case
}
