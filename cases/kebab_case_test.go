package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
		{
			args: args{
				s: "camelCase",
			},
			want: "camel-case",
		},
		{
			args: args{
				s: "snake_case",
			},
			want: "snake-case",
		},
		{
			args: args{
				s: "kebab-case",
			},
			want: "kebab-case",
		},
		{
			args: args{
				s: "PascalCase",
			},
			want: "pascal-case",
		},
		{
			args: args{
				s: "SCREAMING_SNAKE_CASE",
			},
			want: "screaming-snake-case",
		},
		{
			args: args{
				s: "Train-Case",
			},
			want: "train-case",
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

func TestFromKebabToPascalCase(t *testing.T) {
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
				s: "Hello-world",
			},
			want: "HelloWorld",
		},
		{
			args: args{
				s: "Привет-мир-на-дворе-2024-год",
			},
			want: "ПриветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "а-что-если-будет-1-запятая-и-1точка",
			},
			want: "АЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5-рублей-и-10-копеек",
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
		t.Run("FromKebabToPascalCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromKebabToPascalCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromKebabToCamelCase(t *testing.T) {
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
				s: "Hello-world",
			},
			want: "helloWorld",
		},
		{
			args: args{
				s: "Привет-мир-на-дворе-2024-год",
			},
			want: "приветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "а,-что-если-будет-1-запятая-и-1точка",
			},
			want: "а,ЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5-рублей-и-10-копеек",
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
		t.Run("FromKebabToCamelCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromKebabToCamelCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromKebabToSnakeCase(t *testing.T) {
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
				s: "Hello-world",
			},
			want: "hello_world",
		},
		{
			args: args{
				s: "Привет-мир-на-дворе-2024-год",
			},
			want: "привет_мир_на_дворе_2024_год",
		},
		{
			args: args{
				s: "а,-что-если-будет-1-запятая-и-1точка",
			},
			want: "а,_что_если_будет_1_запятая_и_1точка",
		},
		{
			args: args{
				s: "5-рублей-и-10-копеек",
			},
			want: "5_рублей_и_10_копеек",
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
		t.Run("FromKebabToSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromKebabToSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromKebabToScreamingSnakeCase(t *testing.T) {
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
				s: "Hello-world",
			},
			want: "HELLO_WORLD",
		},
		{
			args: args{
				s: "Привет-мир-на-дворе-2024-год",
			},
			want: "ПРИВЕТ_МИР_НА_ДВОРЕ_2024_ГОД",
		},
		{
			args: args{
				s: "а,-что-если-будет-1-запятая-и-1точка",
			},
			want: "А,_ЧТО_ЕСЛИ_БУДЕТ_1_ЗАПЯТАЯ_И_1ТОЧКА",
		},
		{
			args: args{
				s: "5-рублей-и-10-копеек",
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
		t.Run("FromKebabToScreamingSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromKebabToScreamingSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromKebabToTrainCase(t *testing.T) {
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
				s: "hello-world",
			},
			want: "Hello-World",
		},
		{
			args: args{
				s: "Привет-мир-на-дворе-2024-год",
			},
			want: "Привет-Мир-На-Дворе-2024-Год",
		},
		{
			args: args{
				s: "а,-что-если-будет-1-запятая-и-1точка",
			},
			want: "А,-Что-Если-Будет-1-Запятая-И-1точка",
		},
		{
			args: args{
				s: "5-рублей-и-10-копеек",
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FromKebabToTrainCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}
