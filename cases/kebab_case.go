package cases

import (
	"regexp"
	"strings"
)

var kebabCaseRE = regexp.MustCompile("^[a-z]+(-[a-z]+)*$")

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
		case ScreamingSnakeCase:
			return FromScreamingSnakeToKebabCase(s)
		case TrainCase:
			return FromTrainToKebabCase(s)
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

// FromKebabToScreamingSnakeCase converts a kebab-case string to SCREAMING_SNAKE_CASE.
// Keep in mind that it skips spaces cause of these does not match kebab-case.
func FromKebabToScreamingSnakeCase(s string) string {
	return rejoin(s, "-", "_", func(s string, _ int) string {
		return strings.ToUpper(s)
	})
}

// FromKebabToTrainCase converts a kebab-case string to Train-Case.
// Keep in mind that it skips spaces cause of these does not match kebab-case.
func FromKebabToTrainCase(s string) string {
	return rejoin(s, "-", "-", func(s string, _ int) string {
		return ToUpperFirst(s)
	})
}
