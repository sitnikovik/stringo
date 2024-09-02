package cases

import (
	"regexp"
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
)

type StringCase int8

var (
	pascalCaseRE         = regexp.MustCompile("^[A-Z][a-z]+(?:[A-Z][a-z]+)*$")
	camelCaseRE          = regexp.MustCompile("^[a-z]+(?:[A-Z][a-z]+)*$")
	snakeCaseRE          = regexp.MustCompile("^[a-z]+(_[a-z]+)*$")
	kebabCaseRE          = regexp.MustCompile("^[a-z]+(-[a-z]+)*$")
	screamingSnakeCaseRE = regexp.MustCompile("^[A-Z]+(_[A-Z]+)*$")
)

// ToPascalCase converts a string to PascalCase.
func ToPascalCase(s string) string {
	if c := DefineStringCase(s); c != NormalCase {
		switch c {
		case PascalCase:
			return s
		case CamelCase:
			return FromCamelToPascalCase(s)
		case KebabCase:
			return FromKebabToPascalCase(s)
		case SnakeCase:
			return FromSnakeToPascalCase(s)
		}
	}

	words := SplitToWords(s)
	for i, word := range words {
		words[i] = ToUpperFirst(word)
	}

	return strings.Join(words, "")
}

// MatchPascalCase defines if the string matches the PascalCase
func MatchPascalCase(s string) bool {
	return pascalCaseRE.MatchString(s)
}

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

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	if c := DefineStringCase(s); c != NormalCase {
		switch c {
		case SnakeCase:
			return s
		case CamelCase:
			return FromCamelToSnakeCase(s)
		case KebabCase:
			return FromKebabToSnakeCase(s)
		case PascalCase:
			return FromPascalToSnakeCase(s)
		}
	}

	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "_")
}

// MatchSnakeCase defines if the string matches the snake_case
func MatchSnakeCase(s string) bool {
	return snakeCaseRE.MatchString(s)
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	if c := DefineStringCase(s); c != NormalCase {
		switch c {
		case KebabCase:
			return s
		case PascalCase:
			return FromPascalToKebabCase(s)
		case CamelCase:
			return FromCamelToKebabCase(s)
		case SnakeCase:
			return FromSnakeToKebabCase(s)
		}
	}

	words := SplitToWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "-")
}

// MatchKebabCase defines if the string matches the kebab-case
func MatchKebabCase(s string) bool {
	return kebabCaseRE.MatchString(s)
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

// FromSnakeToPascalCase converts a snake_case string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match snake_kase.
func FromSnakeToPascalCase(s string) string {
	return rejoin(s, "_", "", func(s string, idx int) string {
		return ToUpperFirst(strings.ToLower(s))
	})
}

// FromSnakeToCamelCase converts a snake_case string to camelCase.
// Keep in mind that it skips spaces cause of these does not match snake_case.
func FromSnakeToCamelCase(s string) string {
	return rejoin(s, "_", "", func(s string, idx int) string {
		w := strings.ToLower(s)
		if idx > 0 {
			w = ToUpperFirst(w)
		}
		return w
	})
}

// FromSnakeToKebabCase converts a snake_case string to kebab-case.
// Keep in mind that it skips spaces cause of these does not match snake_case.
func FromSnakeToKebabCase(s string) string {
	return rejoin(s, "_", "-", func(s string, _ int) string {
		return strings.ToLower(s)
	})
}

// FromKebabToPascalCase converts a kebab-case string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match kebab-case.
func FromKebabToPascalCase(s string) string {
	return rejoin(s, "-", "", func(s string, idx int) string {
		return ToUpperFirst(strings.ToLower(s))
	})
}

// FromKebabToCamelCase converts a kebab-case string to camelCase.
// Keep in mind that it skips spaces cause of these does not match kebab-case.
func FromKebabToCamelCase(s string) string {
	return rejoin(s, "-", "", func(s string, idx int) string {
		w := strings.ToLower(s)
		if idx > 0 {
			w = ToUpperFirst(w)
		}
		return w
	})
}

// FromKebabToSnakeCase converts a kebab-case string to snake_case.
// Keep in mind that it skips spaces cause of these does not match kebab-case.
func FromKebabToSnakeCase(s string) string {
	return rejoin(s, "-", "_", func(s string, _ int) string {
		return strings.ToLower(s)
	})
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
		func(r rune) rune {
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
		func(r rune) rune {
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

// FromPascalToSnakeCase converts a PascalCase string to snake_case.
// Keep in mind that it skips spaces cause of these does not match PascalCase.
func FromPascalToSnakeCase(s string) string {
	return split(
		s,
		"_",
		func(r rune, idx int) bool {
			return unicode.IsUpper(r) && idx > 0
		},
		func(r rune) rune {
			return unicode.ToLower(r)
		},
	)
}

// FromPascalToKebabCase converts a PascalCase string to kebab-case.
// Keep in mind that it skips spaces cause of these does not match PascalCase.
func FromPascalToKebabCase(s string) string {
	return split(
		s,
		"-",
		func(r rune, idx int) bool {
			return unicode.IsUpper(r) && idx > 0
		},
		func(r rune) rune {
			return unicode.ToLower(r)
		},
	)
}

// FromPascalToCamelCase converts a PascalCase string to camelCase.
// Keep in mind that it skips spaces cause of these does not match PascalCase.
func FromPascalToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	return ToLowerFirst(s)
}
