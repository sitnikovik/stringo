package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToPascalCase(t *testing.T) {
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
				s: "helloworld",
			},
			want: "Helloworld",
		},
		{
			args: args{
				s: "helloWorld",
			},
			want: "HelloWorld",
		},
		{
			args: args{
				s: "hello, World 123",
			},
			want: "HelloWorld123",
		},
		{
			args: args{
				s: "привет мир",
			},
			want: "ПриветМир",
		},
		{
			args: args{
				s: "Аполон13 в космосе",
			},
			want: "Аполон13ВКосмосе",
		},
		{
			args: args{
				s: "camelCase",
			},
			want: "CamelCase",
		},
		{
			args: args{
				s: "snake_case",
			},
			want: "SnakeCase",
		},
		{
			args: args{
				s: "kebab-case",
			},
			want: "KebabCase",
		},
		{
			args: args{
				s: "PascalCase",
			},
			want: "PascalCase",
		},
		{
			args: args{
				s: "Train-Case",
			},
			want: "TrainCase",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("ToPascalCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := ToPascalCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromPascalToSnakeCase(t *testing.T) {
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
				s: "HelloWorld",
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
				s: "HelloWorld123",
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
		t.Run("FromPascalToSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromPascalToSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromPascalToKebabCase(t *testing.T) {
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
				s: "HelloWorld123",
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
		t.Run("FromPascalToKebabCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromPascalToKebabCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromPascalToCamelCase(t *testing.T) {
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
				s: "HelloWorld",
			},
			want: "helloWorld",
		},
		{
			args: args{
				s: "ПриветМир",
			},
			want: "приветМир",
		},
		{
			args: args{
				s: "HelloWorld123",
			},
			want: "helloWorld123",
		},
		{
			args: args{
				s: "ПриветМир123",
			},
			want: "приветМир123",
		},
		{
			args: args{
				s: "Snake_case_changed_with_first_to_lower",
			},
			want: "snake_case_changed_with_first_to_lower",
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
		t.Run("FromPascalToCamelCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromPascalToCamelCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromPascalToTrainCase(t *testing.T) {
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
			args: args{
				s: "helloWorld",
			},
			want: "Hello-World",
		},
		{
			args: args{
				s: "ПриветМир",
			},
			want: "Привет-Мир",
		},
		{
			args: args{
				s: "helloWorld123",
			},
			want: "Hello-World123",
		},
		{
			args: args{
				s: "ПриветМир123",
			},
			want: "Привет-Мир123",
		},
		{
			args: args{
				s: "Аполлон13ВКосмосе",
			},
			want: "Аполлон13-В-Космосе",
		},
		{
			args: args{
				s: "string with spaces to be up with first",
			},
			want: "String with spaces to be up with first",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FromPascalToTrainCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}
