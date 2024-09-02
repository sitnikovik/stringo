package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToCamelCase(t *testing.T) {
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
			want: "приветМир",
		},
		{
			args: args{
				s: "hello, World 123",
			},
			want: "helloWorld123",
		},
		{
			args: args{
				s: "hello world",
			},
			want: "helloWorld",
		},
		{
			args: args{
				s: "Аполон13 в космосе",
			},
			want: "аполон13ВКосмосе",
		},
		{
			args: args{
				s: "camelCase",
			},
			want: "camelCase",
		},
		{
			args: args{
				s: "snake_case",
			},
			want: "snakeCase",
		},
		{
			args: args{
				s: "SCREAMING_SNAKE_CASE",
			},
			want: "screamingSnakeCase",
		},
		{
			args: args{
				s: "kebab-case",
			},
			want: "kebabCase",
		},
		{
			args: args{
				s: "PascalCase",
			},
			want: "pascalCase",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("ToCamelCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := ToCamelCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromCamelToSnakeCase(t *testing.T) {
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
				s: "helloWorld",
			},
			want: "hello_world",
		},
		{
			args: args{
				s: "ПриветМир",
			},
			want: "привет_мир",
		},
		{
			args: args{
				s: "helloWorld123",
			},
			want: "hello_world123",
		},
		{
			args: args{
				s: "ПриветМир123",
			},
			want: "привет_мир123",
		},
		{
			args: args{
				s: "Аполон13ВКосмосе",
			},
			want: "аполон13_в_космосе",
		},
		{
			args: args{
				s: "string with spaces is not convertable",
			},
			want: "string with spaces is not convertable",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromCamelToSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromCamelToSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromCamelToKebabCase(t *testing.T) {
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
				s: "helloWorld",
			},
			want: "hello-world",
		},
		{
			args: args{
				s: "ПриветМир",
			},
			want: "привет-мир",
		},
		{
			args: args{
				s: "helloWorld123",
			},
			want: "hello-world123",
		},
		{
			args: args{
				s: "ПриветМир123",
			},
			want: "привет-мир123",
		},
		{
			args: args{
				s: "Аполон13ВКосмосе",
			},
			want: "аполон13-в-космосе",
		},
		{
			args: args{
				s: "string with spaces is not convertable",
			},
			want: "string with spaces is not convertable",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromCamelToKebabCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromCamelToKebabCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromCamelToPascalCase(t *testing.T) {
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
				s: "helloWorld",
			},
			want: "HelloWorld",
		},
		{
			args: args{
				s: "приветМир",
			},
			want: "ПриветМир",
		},
		{
			args: args{
				s: "helloWorld123",
			},
			want: "HelloWorld123",
		},
		{
			args: args{
				s: "приветМир123",
			},
			want: "ПриветМир123",
		},
		{
			args: args{
				s: "snake_case_changed_with_first_to_upper",
			},
			want: "Snake_case_changed_with_first_to_upper",
		},
		{
			args: args{
				s: "string with spaces is convertable but the first is to upper",
			},
			want: "String with spaces is convertable but the first is to upper",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromCamelToPascalCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromCamelToPascalCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromCamelToScreamingSnakeCase(t *testing.T) {
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
				s: "helloWorld",
			},
			want: "HELLO_WORLD",
		},
		{
			args: args{
				s: "ПриветМир",
			},
			want: "ПРИВЕТ_МИР",
		},
		{
			args: args{
				s: "helloWorld123",
			},
			want: "HELLO_WORLD123",
		},
		{
			args: args{
				s: "ПриветМир123",
			},
			want: "ПРИВЕТ_МИР123",
		},
		{
			args: args{
				s: "Аполлон13ВКосмосе",
			},
			want: "АПОЛЛОН13_В_КОСМОСЕ",
		},
		{
			args: args{
				s: "string with spaces will be upped",
			},
			want: "STRING WITH SPACES WILL BE UPPED",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromCamelToScreamingSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromCamelToScreamingSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}
