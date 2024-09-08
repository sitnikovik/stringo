package cases

import (
	"regexp"
	"strings"
)

var screamingSnakeCaseRE = regexp.MustCompile("^[A-Z]+(_[A-Z]+)*$")

// ToScreamingSnakeCase converts a string to SCREAMING_SNAKE_CASE
func ToScreamingSnakeCase(s string) string {
	if c := DefineStringCase(s); c != NormalCase {
		switch c {
		case ScreamingSnakeCase:
			return s
		case PascalCase:
			return FromPascalToScreamingSnakeCase(s)
		case CamelCase:
			return FromCamelToScreamingSnakeCase(s)
		case KebabCase:
			return FromKebabToScreamingSnakeCase(s)
		case SnakeCase:
			return FromSnakeToScreamingSnakeCase(s)
		}
	}

	return mapWordsAndJoin(s, "_", func(s string, _ int) string {
		return strings.ToUpper(s)
	})
}

// MatchScreamingSnakeCase defines if the string matches the SCREAMING_SNAKE_CASE
func MatchScreamingSnakeCase(s string) bool {
	return screamingSnakeCaseRE.MatchString(s)
}

// FromScreamingSnakeToPascalCase converts a SCREAMING_SNAKE_CASE string to PascalCase.
// Keep in mind that it skips spaces cause of these does not match SCREAMING_SNAKE_CASE.
func FromScreamingSnakeToPascalCase(s string) string {
	return rejoin(s, "_", "", func(s string, idx int) string {
		return ToUpperFirst(strings.ToLower(s))
	})
}

// FromScreamingSnakeToCamelCase converts a SCREAMING_SNAKE_CASE string to camelCase.
// Keep in mind that it skips spaces cause of these does not match SCREAMING_SNAKE_CASE.
func FromScreamingSnakeToCamelCase(s string) string {
	return rejoin(s, "_", "", func(s string, idx int) string {
		w := strings.ToLower(s)
		if idx > 0 {
			w = ToUpperFirst(w)
		}
		return w
	})
}

// FromScreamingSnakeToKebabCase converts a SCREAMING_SNAKE_CASE string to kebab-case.
// Keep in mind that it skips spaces cause of these does not match SCREAMING_SNAKE_CASE.
func FromScreamingSnakeToKebabCase(s string) string {
	return rejoin(s, "_", "-", func(s string, _ int) string {
		return strings.ToLower(s)
	})
}

// FromScreamingSnakeToSnakeCase converts a SCREAMING_SNAKE_CASE string to snake_case.
// Keep in mind that it skips spaces cause of these does not match SCREAMING_SNAKE_CASE.
func FromScreamingSnakeToSnakeCase(s string) string {
	return rejoin(s, "_", "_", func(s string, _ int) string {
		return strings.ToLower(s)
	})
}

// FromTrainToScreamingSnakeCase converts a SCREAMING_SNAKE_CASE string to Train-Case.
// Keep in mind that it skips spaces cause of these does not match SCREAMING_SNAKE_CASE.
func FromScreamingCaseToTrainCase(s string) string {
	return rejoin(s, "_", "-", func(s string, _ int) string {
		return ToUpperFirst(strings.ToLower(s))
	})
}
