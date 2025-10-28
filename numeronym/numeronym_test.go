package numeronym

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumeronymASCII(t *testing.T) {
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
			name: "kubernetes",
			args: args{
				s: "kubernetes",
			},
			want: "k8s",
		},
		{
			name: "cat",
			args: args{
				s: "cat",
			},
			want: "c1t",
		},
		{
			name: "INTERNATIONALIZATION",
			args: args{
				s: "INTERNATIONALIZATION",
			},
			want: "I18N",
		},
		{
			name: "accessibility",
			args: args{
				s: "accessibility",
			},
			want: "a11y",
		},
		{
			name: "localization",
			args: args{
				s: "localization",
			},
			want: "l10n",
		},
		{
			name: "to",
			args: args{
				s: "to",
			},
			want: "to",
		},
		{
			name: "t",
			args: args{
				s: "t",
			},
			want: "t",
		},
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "post-office",
			args: args{
				s: "post-office",
			},
			want: "p9e",
		},
		{
			name: "google.com",
			args: args{
				s: "google.com",
			},
			want: "g8m",
		},
		{
			name: "-kebab-case-",
			args: args{
				s: "-kebab-case-",
			},
			want: "-10-",
		},
		{
			name: "-123",
			args: args{
				s: "-123",
			},
			want: "-23",
		},
		{
			name: "Supercalifragilisticexpialidocious",
			args: args{
				s: "Supercalifragilisticexpialidocious",
			},
			want: "S32s",
		},
		{
			name: "мир",
			args: args{
				s: "мир",
			},
			want: "\xd04\x80",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NumeronymASCII(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNumeronym(t *testing.T) {
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
			name: "kubernetes",
			args: args{
				s: "kubernetes",
			},
			want: "k8s",
		},
		{
			name: "cat",
			args: args{
				s: "cat",
			},
			want: "c1t",
		},
		{
			name: "INTERNATIONALIZATION",
			args: args{
				s: "INTERNATIONALIZATION",
			},
			want: "I18N",
		},
		{
			name: "accessibility",
			args: args{
				s: "accessibility",
			},
			want: "a11y",
		},
		{
			name: "localization",
			args: args{
				s: "localization",
			},
			want: "l10n",
		},
		{
			name: "to",
			args: args{
				s: "to",
			},
			want: "to",
		},
		{
			name: "t",
			args: args{
				s: "t",
			},
			want: "t",
		},
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "post-office",
			args: args{
				s: "post-office",
			},
			want: "p9e",
		},
		{
			name: "google.com",
			args: args{
				s: "google.com",
			},
			want: "g8m",
		},
		{
			name: "-kebab-case-",
			args: args{
				s: "-kebab-case-",
			},
			want: "-10-",
		},
		{
			name: "-123",
			args: args{
				s: "-123",
			},
			want: "-23",
		},
		{
			name: "Supercalifragilisticexpialidocious",
			args: args{
				s: "Supercalifragilisticexpialidocious",
			},
			want: "S32s",
		},
		{
			name: "мир",
			args: args{
				s: "мир",
			},
			want: "м1р",
		},
		{
			name: "chinese",
			args: args{
				s: "国际化",
			},
			want: "国1化",
		},
		{
			name: "emoji",
			args: args{
				s: "😀😃😄😁😆😅😂🤣😊😇",
			},
			want: "😀8😇",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Numeronym(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
