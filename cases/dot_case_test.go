package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatchDotCase(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				s: "this.is.dot.case",
			},
			want: true,
		},
		{
			name: "false on simple sentences",
			args: args{
				s: "Hello there. Good to see you.",
			},
			want: false,
		},
		{
			name: "false on snake_case",
			args: args{
				s: "not_dot_case",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, MatchDotCase(tt.args.s))
		})
	}
}

func TestToDotCase(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple sentence",
			args: args{
				s: "Hello there. Good to see you.",
			},
			want: "hello.there.good.to.see.you",
		},
		{
			name: "snake_case",
			args: args{
				s: "this_is_snake_case",
			},
			want: "this.is.snake.case",
		},
		{
			name: "kebab-case",
			args: args{
				s: "this-is-kebab-case",
			},
			want: "this.is.kebab.case",
		},
		{
			name: "camelCase",
			args: args{
				s: "thisIsCamelCase",
			},
			want: "this.is.camel.case",
		},
		{
			name: "PascalCase",
			args: args{
				s: "ThisIsPascalCase",
			},
			want: "this.is.pascal.case",
		},
		{
			name: "SCREAMING_SNAKE_CASE",
			args: args{
				s: "THIS_IS_SCREAMING_SNAKE_CASE",
			},
			want: "this.is.screaming.snake.case",
		},
		{
			name: "Train-Case",
			args: args{
				s: "THIS-IS-TRAIN-CASE",
			},
			want: "this.is.train.case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, ToDotCase(tt.args.s))
		})
	}
}
