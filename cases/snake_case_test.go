package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
		{
			args: args{
				s: "camelCase",
			},
			want: "camel_case",
		},
		{
			args: args{
				s: "snake_case",
			},
			want: "snake_case",
		},
		{
			args: args{
				s: "kebab-case",
			},
			want: "kebab_case",
		},
		{
			args: args{
				s: "PascalCase",
			},
			want: "pascal_case",
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

func TestFromSnakeToPascalCase(t *testing.T) {
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
				s: "Hello_world",
			},
			want: "HelloWorld",
		},
		{
			args: args{
				s: "Привет_мир_на_дворе_2024_год",
			},
			want: "ПриветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "а_что_если_будет_1_запятая_и_1точка",
			},
			want: "АЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5_рублей_и_10_копеек",
			},
			want: "5РублейИ10Копеек",
		},
		{
			args: args{
				s: "string with spaces is not convertable",
			},
			want: "String with spaces is not convertable",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromSnakeToPascalCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromSnakeToPascalCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromSnakeToCamelCase(t *testing.T) {
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
				s: "Hello_world",
			},
			want: "helloWorld",
		},
		{
			args: args{
				s: "Привет_мир_на_дворе_2024_год",
			},
			want: "приветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "а,_что_если_будет_1_запятая_и_1точка",
			},
			want: "а,ЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5_рублей_и_10_копеек",
			},
			want: "5РублейИ10Копеек",
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
		t.Run("FromSnakeToCamelCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromSnakeToCamelCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromSnakeToKebabCase(t *testing.T) {
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
				s: "Hello_world",
			},
			want: "hello-world",
		},
		{
			args: args{
				s: "Привет_мир_на_дворе_2024_год",
			},
			want: "привет-мир-на-дворе-2024-год",
		},
		{
			args: args{
				s: "а,_что_если_будет_1_запятая_и_1точка",
			},
			want: "а,-что-если-будет-1-запятая-и-1точка",
		},
		{
			args: args{
				s: "5_рублей_и_10_копеек",
			},
			want: "5-рублей-и-10-копеек",
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
		t.Run("FromSnakeToKebabCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromSnakeToKebabCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromSnakeToScreamingSnakeCase(t *testing.T) {
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
				s: "Hello_world",
			},
			want: "HELLO_WORLD",
		},
		{
			args: args{
				s: "Привет_мир_на_дворе_2024_год",
			},
			want: "ПРИВЕТ_МИР_НА_ДВОРЕ_2024_ГОД",
		},
		{
			args: args{
				s: "а,_что_если_будет_1_запятая_и_1точка",
			},
			want: "А,_ЧТО_ЕСЛИ_БУДЕТ_1_ЗАПЯТАЯ_И_1ТОЧКА",
		},
		{
			args: args{
				s: "5_рублей_и_10_копеек",
			},
			want: "5_РУБЛЕЙ_И_10_КОПЕЕК",
		},
		{
			args: args{
				s: "string with spaces to be up",
			},
			want: "STRING WITH SPACES TO BE UP",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("FromSnakeToScreamingSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromSnakeToScreamingSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromSnakeToTrainCase(t *testing.T) {
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
				s: "Hello_world",
			},
			want: "Hello-World",
		},
		{
			args: args{
				s: "Привет_мир_на_дворе_2024_год",
			},
			want: "Привет-Мир-На-Дворе-2024-Год",
		},
		{
			args: args{
				s: "а,_что_если_будет_1_запятая_и_1точка",
			},
			want: "А,-Что-Если-Будет-1-Запятая-И-1точка",
		},
		{
			args: args{
				s: "5_рублей_и_10_копеек",
			},
			want: "5-Рублей-И-10-Копеек",
		},
		{
			args: args{
				s: "string with spaces to be up with first letter",
			},
			want: "String with spaces to be up with first letter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromSnakeToTrainCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}
