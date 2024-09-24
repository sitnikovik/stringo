package stringo

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func TestSplitFunc(t *testing.T) {
	t.Parallel()

	type args struct {
		s      string
		cond   func(rune, int) bool
		incSep bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "nil func provied",
			args: args{
				s:    "hello world",
				cond: nil,
			},
			want: []string{"hello world"},
		},
		{
			name: "split by space",
			args: args{
				s:    "hello world",
				cond: func(r rune, i int) bool { return unicode.IsSpace(r) },
			},
			want: []string{"hello", "world"},
		},
		{
			name: "split camelCase",
			args: args{
				s:      "helloWorldItWorks",
				cond:   func(r rune, i int) bool { return unicode.IsUpper(r) },
				incSep: true,
			},
			want: []string{"hello", "World", "It", "Works"},
		},
		{
			name: "split camelCase in string with spaces",
			args: args{
				s:      "helloWorld itWorks ",
				cond:   func(r rune, i int) bool { return unicode.IsUpper(r) || unicode.IsSpace(r) },
				incSep: true,
			},
			want: []string{"hello", "World", " it", "Works", " "},
		},
		{
			name: "split camelCase string but only with space and include separator",
			args: args{
				s:      "helloWorld itWorks ",
				cond:   func(r rune, i int) bool { return unicode.IsSpace(r) },
				incSep: true,
			},
			want: []string{"helloWorld", " itWorks", " "},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := SplitFunc(tt.args.s, tt.args.cond, tt.args.incSep)

			require.Equal(t, tt.want, got)
		})
	}
}
