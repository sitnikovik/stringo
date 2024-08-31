package stringo

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

func TestToSnakeCase(t *testing.T) {
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
				s: "hello world",
			},
			want: "hello_world",
		},
		{
			args: args{
				s: "hello, World 123",
			},
			want: "hello_world_123",
		},
		{
			args: args{
				s: "привет мир",
			},
			want: "привет_мир",
		},
		{
			args: args{
				s: "Аполон13 в космосе",
			},
			want: "аполон13_в_космосе",
		},
		{
			args: args{
				s: "Аполон_13 в космосе",
			},
			want: "аполон_13_в_космосе",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("ToSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := ToSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestToKebabCase(t *testing.T) {
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
				s: "hello world",
			},
			want: "hello-world",
		},
		{
			args: args{
				s: "hello, World 123",
			},
			want: "hello-world-123",
		},
		{
			args: args{
				s: "привет мир",
			},
			want: "привет-мир",
		},
		{
			args: args{
				s: "Аполон-13 в космосе",
			},
			want: "аполон-13-в-космосе",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("ToKebabCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := ToKebabCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

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
