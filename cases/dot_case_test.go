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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, ToDotCase(tt.args.s))
		})
	}
}
