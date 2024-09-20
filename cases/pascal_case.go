package cases

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/sitnikovik/stringo"
)

var pascalCaseRE = regexp.MustCompile("^[A-Z][a-z]+(?:[A-Z][a-z]+)*$")

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
		case ScreamingSnakeCase:
			return FromScreamingSnakeToPascalCase(s)
		case TrainCase:
			return FromTrainToPascalCase(s)
		case DotCase:
			return FromDotToPascalCase(s)
		}
	}

	return mapWordsAndJoin(s, "", func(s string, _ int) string {
		return ToUpperFirst(strings.ToLower(s))
	})
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
		func(r rune, _ int) rune {
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
		func(r rune, _ int) rune {
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

// FromPascalToScreamingSnakeCase converts a PascalCase string to SCREAMING_SNAKE_CASE.
// Keep in mind that it skips spaces cause of these does not match PascalCase.
func FromPascalToScreamingSnakeCase(s string) string {
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

// FromPascalToTrainCase converts a PascalCase string to Train-Case.
// Keep in mind that it skips spaces cause of these does not match PascalCase.
func FromPascalToTrainCase(s string) string {
	ss := stringo.SplitFunc(s, func(r rune, idx int) bool {
		return idx > 0 && unicode.IsUpper(r)
	}, true)

	for i, w := range ss {
		ss[i] = ToUpperFirst(strings.ToLower(w))
	}

	return strings.Join(ss, "-")
}

// FromPascalToDotCase converts a PascalCase string to dot.case.
// Keep in mind that it skips spaces cause of these does not match PascalCase.
func FromPascalToDotCase(s string) string {
	return split(
		s,
		".",
		func(r rune, idx int) bool {
			return unicode.IsUpper(r) && idx > 0
		},
		func(r rune, _ int) rune {
			return unicode.ToLower(r)
		},
	)
}
