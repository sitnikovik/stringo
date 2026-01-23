package stringo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfRune(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
		r rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "rune present in string",
			args: args{s: "hello", r: 'e'},
			want: 1,
		},
		{
			name: "rune not present in string",
			args: args{s: "hello", r: 'a'},
			want: -1,
		},
		{
			name: "newline rune present in string",
			args: args{s: "hello\n", r: '\n'},
			want: 5,
		},
		{
			name: "first rune in string",
			args: args{s: "world", r: 'w'},
			want: 0,
		},
		{
			name: "last rune in string",
			args: args{s: "world", r: 'd'},
			want: 4,
		},
		{
			name: "empty string",
			args: args{s: "", r: 'a'},
			want: -1,
		},
		{
			name: "unicode rune present in string",
			args: args{s: "こんにちは", r: 'に'},
			want: 6,
		},
		{
			name: "unicode rune not present in string",
			args: args{s: "こんにちは", r: 'さ'},
			want: -1,
		},
		{
			name: "multiple occurrences of rune",
			args: args{s: "banana", r: 'a'},
			want: 1,
		},
		{
			name: "space rune present in string",
			args: args{s: "hello world", r: ' '},
			want: 5,
		},
		{
			name: "tab rune present in string",
			args: args{s: "hello\tworld", r: '\t'},
			want: 5,
		},
		{
			name: "special character rune present in string",
			args: args{s: "hello@world", r: '@'},
			want: 5,
		},
		{
			name: "digit rune present in string",
			args: args{s: "version1.0", r: '1'},
			want: 7,
		},
		{
			name: "emoji rune present in string",
			args: args{s: "hello😊world", r: '😊'},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := IndexOfRune(tt.args.s, tt.args.r)
			assert.Equal(t, tt.want, got)
		})
	}
}
