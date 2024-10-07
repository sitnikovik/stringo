package dates

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDynamic(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "now",
			args: args{t: time.Now()},
			want: "now",
		},
		{
			name: "1 minute ago",
			args: args{t: time.Now().Add(-time.Minute)},
			want: "1 minutes ago",
		},
		{
			name: "today",
			args: args{t: time.Now().Add(-time.Hour)},
			want: "today " + time.Now().Add(-time.Hour).Format("15:04"),
		},
		{
			name: "yesterday",
			args: args{t: time.Now().Add(-time.Hour * 24)},
			want: "yesterday " + time.Now().Add(-time.Hour*24).Format("15:04"),
		},
		{
			name: "formatted",
			args: args{t: time.Now().Add(-time.Hour * 24 * 2)},
			want: time.Now().Add(-time.Hour * 24 * 2).Format("02.01.2006 15:04"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, Dynamic(tt.args.t))
		})
	}
}

func TestShortDMY(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{t: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC)},
			want: "2 Jan 2021",
		},
		{
			args: args{t: time.Date(2021, 12, 31, 23, 59, 59, 0, time.UTC)},
			want: "31 Dec 2021",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, ShortDMY(tt.args.t))
		})
	}
}

func TestShortMY(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{t: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC)},
			want: "Jan 2021",
		},
		{
			args: args{t: time.Date(2021, 12, 31, 23, 59, 59, 0, time.UTC)},
			want: "Dec 2021",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, ShortMY(tt.args.t))
		})
	}
}

func TestShortMYHM(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{t: time.Date(2021, 1, 2, 3, 4, 5, 6, time.UTC)},
			want: "Jan 2021 03:04",
		},
		{
			args: args{t: time.Date(2021, 12, 31, 23, 59, 59, 0, time.UTC)},
			want: "Dec 2021 23:59",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.want, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, ShortMYHM(tt.args.t))
		})
	}
}
