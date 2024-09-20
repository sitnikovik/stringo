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
			name: "ok on dot case",
			args: args{
				s: "this.is.dot.case",
			},
			want: "this.is.dot.case",
		},
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, ToDotCase(tt.args.s))
		})
	}
}

func TestFromDotToPascalCase(t *testing.T) {
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
				s: "Hello.world",
			},
			want: "HelloWorld",
		},
		{
			args: args{
				s: "Привет.мир.на.дворе.2024.год",
			},
			want: "ПриветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "а.что.если.будет.1.запятая.и.1точка",
			},
			want: "АЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5.рублей.и.10.копеек",
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
		t.Run(tt.args.s, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, FromDotToPascalCase(tt.args.s))
		})
	}
}

func TestFromDotToCamelCase(t *testing.T) {
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
				s: "Hello.world",
			},
			want: "helloWorld",
		},
		{
			args: args{
				s: "Привет.мир.на.дворе.2024.год",
			},
			want: "приветМирНаДворе2024Год",
		},
		{
			args: args{
				s: "а,.что.если.будет.1.запятая.и.1точка",
			},
			want: "а,ЧтоЕслиБудет1ЗапятаяИ1точка",
		},
		{
			args: args{
				s: "5.рублей.и.10.копеек",
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
		t.Run(tt.args.s, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, FromDotToCamelCase(tt.args.s))
		})
	}
}

func TestFromDotToSnakeCase(t *testing.T) {
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
				s: "Hello.world",
			},
			want: "hello_world",
		},
		{
			args: args{
				s: "Привет.мир.на.дворе.2024.год",
			},
			want: "привет_мир_на_дворе_2024_год",
		},
		{
			args: args{
				s: "а,.что.если.будет.1.запятая.и.1точка",
			},
			want: "а,_что_если_будет_1_запятая_и_1точка",
		},
		{
			args: args{
				s: "5.рублей.и.10.копеек",
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
		t.Run(tt.args.s, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, FromDotToSnakeCase(tt.args.s))
		})
	}
}

func TestFromDotToScreamingSnakeCase(t *testing.T) {
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
				s: "Hello.world",
			},
			want: "HELLO_WORLD",
		},
		{
			args: args{
				s: "Привет.мир.на.дворе.2024.год",
			},
			want: "ПРИВЕТ_МИР_НА_ДВОРЕ_2024_ГОД",
		},
		{
			args: args{
				s: "а,.что.если.будет.1.запятая.и.1точка",
			},
			want: "А,_ЧТО_ЕСЛИ_БУДЕТ_1_ЗАПЯТАЯ_И_1ТОЧКА",
		},
		{
			args: args{
				s: "5.рублей.и.10.копеек",
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
		t.Run(tt.args.s, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, FromDotToScreamingSnakeCase(tt.args.s))
		})
	}
}

func TestFromDotToTrainCase(t *testing.T) {
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
				s: "hello.world",
			},
			want: "Hello-World",
		},
		{
			args: args{
				s: "Привет.мир.на.дворе.2024.год",
			},
			want: "Привет-Мир-На-Дворе-2024-Год",
		},
		{
			args: args{
				s: "а,.что.если.будет.1.запятая.и.1точка",
			},
			want: "А,-Что-Если-Будет-1-Запятая-И-1точка",
		},
		{
			args: args{
				s: "5.рублей.и.10.копеек",
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

			got := FromDotToTrainCase(tt.args.s)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFromDotToKebabCase(t *testing.T) {
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
				s: "Hello.world",
			},
			want: "hello-world",
		},
		{
			args: args{
				s: "Привет.мир.на.дворе.2024.год",
			},
			want: "привет-мир-на-дворе-2024-год",
		},
		{
			args: args{
				s: "а,.что.если.будет.1.запятая.и.1точка",
			},
			want: "а,-что-если-будет-1-запятая-и-1точка",
		},
		{
			args: args{
				s: "5.рублей.и.10.копеек",
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
		t.Run(tt.args.s, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, FromDotToKebabCase(tt.args.s))
		})
	}
}
