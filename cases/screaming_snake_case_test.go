package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToScreamingSnakeCase(t *testing.T) {
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
			want: "HELLO_WORLD",
		},
		{
			args: args{
				s: "hello, World 123",
			},
			want: "HELLO_WORLD_123",
		},
		{
			args: args{
				s: "привет мир",
			},
			want: "ПРИВЕТ_МИР",
		},
		{
			args: args{
				s: "Аполлон13 в космосе",
			},
			want: "АПОЛЛОН13_В_КОСМОСЕ",
		},
		{
			args: args{
				s: "Аполлон_13 в космосе",
			},
			want: "АПОЛЛОН_13_В_КОСМОСЕ",
		},
		{
			args: args{
				s: "camelCase",
			},
			want: "CAMEL_CASE",
		},
		{
			args: args{
				s: "snake_case",
			},
			want: "SNAKE_CASE",
		},
		{
			args: args{
				s: "kebab-case",
			},
			want: "KEBAB_CASE",
		},
		{
			args: args{
				s: "PascalCase",
			},
			want: "PASCAL_CASE",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("ToScreamingSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := ToScreamingSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromScreamingSnakeToPascalCase(t *testing.T) {
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
				s: "HELLO_WORLD",
			},
			want: "HelloWorld",
		},
		{
			args: args{
				s: "ПРИВЕТ_МИР_НА_ДВОРЕ_2024_ГОД",
			},
			want: "ПриветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "А_ЧТО_ЕСЛИ_БУДЕТ_1_ЗАПЯТАЯ_И_1ТОЧКА",
			},
			want: "АЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5_РУБЛЕЙ_И_10_КОПЕЕК",
			},
			want: "5РублейИ10Копеек",
		},
		{
			args: args{
				s: "STRING WITH SPACES IS NOT CONVERTABLE",
			},
			want: "String with spaces is not convertable",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromScreamingSnakeToPascalCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromScreamingSnakeToPascalCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromScreamingSnakeToCamelCase(t *testing.T) {
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
				s: "HELLO_WORLD",
			},
			want: "helloWorld",
		},
		{
			args: args{
				s: "ПРИВЕТ_МИР_НА_ДВОРЕ_2024_ГОД",
			},
			want: "приветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "А,_ЧТО_ЕСЛИ_БУДЕТ_1_ЗАПЯТАЯ_И_1ТОЧКА",
			},
			want: "а,ЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5_РУБЛЕЙ_И_10_КОПЕЕК",
			},
			want: "5РублейИ10Копеек",
		},
		{
			args: args{
				s: "STRING WITH SPACES IS NOT CONVERTABLE",
			},
			want: "string with spaces is not convertable",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromScreamingSnakeToCamelCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromScreamingSnakeToCamelCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromScreamingSnakeToKebabCase(t *testing.T) {
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
				s: "HELLO_WORLD",
			},
			want: "hello-world",
		},
		{
			args: args{
				s: "ПРИВЕТ_МИР_НА_ДВОРЕ_2024_ГОД",
			},
			want: "привет-мир-на-дворе-2024-год",
		},
		{
			args: args{
				s: "А,_ЧТО_ЕСЛИ_БУДЕТ_1_ЗАПЯТАЯ_И_1ТОЧКА",
			},
			want: "а,-что-если-будет-1-запятая-и-1точка",
		},
		{
			args: args{
				s: "5_РУБЛЕЙ_И_10_КОПЕЕК",
			},
			want: "5-рублей-и-10-копеек",
		},
		{
			args: args{
				s: "STRING WITH SPACES IS NOT CONVERTABLE",
			},
			want: "string with spaces is not convertable",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromScreamingSnakeToKebabCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromScreamingSnakeToKebabCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}
