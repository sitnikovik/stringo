package cases

import (
	"regexp"
	"unicode"
)

var pascalCaseRE = regexp.MustCompile("^[A-Z][a-z]+(?:[A-Z][a-z]+)*$")

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
