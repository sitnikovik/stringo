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
				s: "string with spaces is not convertable",
			},
			want: "String with spaces is not convertable",
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
