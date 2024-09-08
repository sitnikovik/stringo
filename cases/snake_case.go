package cases

import (
	"regexp"
	"strings"
)

var snakeCaseRE = regexp.MustCompile("^[a-z]+(_[a-z]+)*$")

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
		case ScreamingSnakeCase:
			return FromScreamingSnakeToSnakeCase(s)
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

// FromSnakeToScreamingSnakeCase converts a snake_case string to SCREAMING_SNAKE_CASE.
// Keep in mind that it skips spaces cause of these does not match snake_case.
func FromSnakeToScreamingSnakeCase(s string) string {
	return rejoin(s, "_", "_", func(s string, _ int) string {
		return strings.ToUpper(s)
	})
}

// FromSnakeToTrainCase converts a snake_case string to Train-Case.
// Keep in mind that it skips spaces cause of these does not match snake_case.
func FromSnakeToTrainCase(s string) string {
	return rejoin(s, "_", "-", func(s string, _ int) string {
		return ToUpperFirst(strings.ToLower(s))
	})
}
