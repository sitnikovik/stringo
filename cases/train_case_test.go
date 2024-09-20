package cases

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToTrainCase(t *testing.T) {
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
			want: "Hello-World",
		},
		{
			args: args{
				s: "hello, World 123",
			},
			want: "Hello-World-123",
		},
		{
			args: args{
				s: "привет мир",
			},
			want: "Привет-Мир",
		},
		{
			args: args{
				s: "Аполон-13 в космосе",
			},
			want: "Аполон-13-В-Космосе",
		},
		{
			args: args{
				s: "camelCase",
			},
			want: "Camel-Case",
		},
		{
			args: args{
				s: "snake_case",
			},
			want: "Snake-Case",
		},
		{
			args: args{
				s: "kebab-case",
			},
			want: "Kebab-Case",
		},
		{
			args: args{
				s: "PascalCase",
			},
			want: "Pascal-Case",
		},
		{
			args: args{
				s: "SCREAMING_SNAKE_CASE",
			},
			want: "Screaming-Snake-Case",
		},
		{
			args: args{
				s: "Train-Case",
			},
			want: "Train-Case",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("ToTrainCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := ToTrainCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromTrainToPascalCase(t *testing.T) {
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
				s: "Hello-World",
			},
			want: "HelloWorld",
		},
		{
			args: args{
				s: "Привет-Мир-На-Дворе-2024-Год",
			},
			want: "ПриветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "А-Что-Если-Будет-1-Запятая-И-1точка",
			},
			want: "АЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5-Рублей-И-10-Копеек",
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
		t.Run("FromTrainToPascalCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromTrainToPascalCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromTraomToCamelCase(t *testing.T) {
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
				s: "Hello-World",
			},
			want: "helloWorld",
		},
		{
			args: args{
				s: "Привет-Мир-На-Дворе-2024-Год",
			},
			want: "приветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "А,-Что-Если-Будет-1-Запятая-И-1точка",
			},
			want: "а,ЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5-Рублей-И-10-Копеек",
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
		t.Run("FromTrainToCamelCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromTrainToCamelCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromTrainToSnakeCase(t *testing.T) {
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
				s: "Hello-World",
			},
			want: "hello_world",
		},
		{
			args: args{
				s: "Привет-Мир-На-Дворе-2024-Год",
			},
			want: "привет_мир_на_дворе_2024_год",
		},
		{
			args: args{
				s: "А,-Что-Если-Будет-1-Запятая-И-1точка",
			},
			want: "а,_что_если_будет_1_запятая_и_1точка",
		},
		{
			args: args{
				s: "5-Рублей-И-10-Копеек",
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
		t.Run("FromTrainToSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromTrainToSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromTrainbToScreamingSnakeCase(t *testing.T) {
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
				s: "Hello-World",
			},
			want: "HELLO_WORLD",
		},
		{
			args: args{
				s: "Привет-Мир-На-Дворе-2024-Год",
			},
			want: "ПРИВЕТ_МИР_НА_ДВОРЕ_2024_ГОД",
		},
		{
			args: args{
				s: "А,-Что-Если-Будет-1-Запятая-И-1точка",
			},
			want: "А,_ЧТО_ЕСЛИ_БУДЕТ_1_ЗАПЯТАЯ_И_1ТОЧКА",
		},
		{
			args: args{
				s: "5-Рублей-И-10-Копеек",
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
		t.Run("FromTrainToScreamingSnakeCase:"+tt.args.s, func(t *testing.T) {
			t.Parallel()

			got := FromTrainToScreamingSnakeCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromTrainToDotCase(t *testing.T) {
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
				s: "Hello-World",
			},
			want: "hello.world",
		},
		{
			args: args{
				s: "Привет-Мир-На-Дворе-2024-Год",
			},
			want: "привет.мир.на.дворе.2024.год",
		},
		{
			args: args{
				s: "А,-Что-Если-Будет-1-Запятая-И-1точка",
			},
			want: "а,.что.если.будет.1.запятая.и.1точка",
		},
		{
			args: args{
				s: "5-Рублей-И-10-Копеек",
			},
			want: "5.рублей.и.10.копеек",
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
		t.Run(tt.args.s, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, FromTrainToDotCase(tt.args.s))
		})
	}
}
