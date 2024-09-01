package stringo

import (
	"net/url"
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
		{
			args: args{
				s: "userId",
			},
			want: "user_id",
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

func TestUrlValuesToSnakeCase(t *testing.T) {
	t.Parallel()

	type args struct {
		vals url.Values
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		{
			args: args{
				vals: url.Values{
					"userId":       []string{"userId"},
					"age-category": []string{"age-category"},
					"UserRole":     []string{"UserRole"},
					"trace-id":     []string{"trace-id"},
				},
			},
			want: url.Values{
				"user_id":      []string{"userId"},
				"age_category": []string{"age-category"},
				"user_role":    []string{"UserRole"},
				"trace_id":     []string{"trace-id"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := UrlValuesToSnakeCase(tt.args.vals)

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
