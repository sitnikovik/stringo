package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToUpperFirst(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				s: "привет мир",
			},
			want: "Привет мир",
		},
		{
			args: args{
				s: "hello world",
			},
			want: "Hello world",
		},
		{
			args: args{
				s: "h",
			},
			want: "H",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("ToUpperFirst:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := ToUpperFirst(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestToLowerFirst(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				s: "Привет мир",
			},
			want: "привет мир",
		},
		{
			args: args{
				s: "Hello world",
			},
			want: "hello world",
		},
		{
			args: args{
				s: "H",
			},
			want: "h",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("ToLowerFirst:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := ToLowerFirst(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestMatchSnakeCase(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				s: "some_string",
			},
			want: true,
		},
		{
			args: args{
				s: "someString",
			},
			want: false,
		},
		{
			args: args{
				s: "some_string with space",
			},
			want: false,
		},
		{
			args: args{
				s: "SomeString",
			},
			want: false,
		},
		{
			args: args{
				s: "some-string_withUnderScore",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("MatchSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := MatchSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestMatchPascalCase(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{

		{
			args: args{
				s: "RealPascalCase",
			},
			want: true,
		},
		{
			args: args{
				s: "CAPSLOCK",
			},
		},
		{
			args: args{
				s: "camelCase",
			},
		},
		{
			args: args{
				s: "snake_case",
			},
		},
		{
			args: args{
				s: "kebab-case",
			},
		},
		{
			args: args{
				s: "mix_stringWith-Many VARIANTS",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("MatchPascalCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := MatchPascalCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestMatchCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{

		{
			args: args{
				s: "PascalCase",
			},
		},
		{
			args: args{
				s: "CAPSLOCK",
			},
		},
		{
			args: args{
				s: "camelCase",
			},
			want: true,
		},
		{
			args: args{
				s: "snake_case",
			},
		},
		{
			args: args{
				s: "kebab-case",
			},
		},
		{
			args: args{
				s: "mix_stringWith-Many VARIANTS",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("MatchCamelCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := MatchCamelCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestMatchKebabCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				s: "some_string",
			},
		},
		{
			args: args{
				s: "someString",
			},
		},
		{
			args: args{
				s: "some-string",
			},
			want: true,
		},
		{
			args: args{
				s: "some_string with space",
			},
		},
		{
			args: args{
				s: "SomeString",
			},
		},
		{
			args: args{
				s: "some-string_withUnderScore and Spaces",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("MatchKebabCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := MatchKebabCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestDefineStringCase(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want StringCase
	}{
		{
			name: "define KebabCase",
			args: args{
				s: "kebab-case",
			},
			want: KebabCase,
		},
		{
			name: "define SnakeCase",
			args: args{
				s: "snake_case",
			},
			want: SnakeCase,
		},
		{

			name: "define CamelCase",
			args: args{
				s: "camelCase",
			},
			want: CamelCase,
		},
		{

			name: "define PascalCase",
			args: args{
				s: "PascalCase",
			},
			want: PascalCase,
		},
		{

			name: "define NormalCase on a sentence",
			args: args{
				s: "Who wants to be a millionaire",
			},
			want: NormalCase,
		},
		{
			name: "define NormalCase on mixes",
			args: args{
				s: "some-string_withUnderScore and Spaces",
			},
			want: NormalCase,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := DefineStringCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}
