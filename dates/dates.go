package dates

import (
	"fmt"
	"time"
)

// Dynamic returns a human-readable time difference string.
func Dynamic(t time.Time) string {
	diff := time.Since(t)

	switch {
	case diff < time.Minute:
		// Less than a minute, return "now".
		return "now"
	case diff < time.Hour:
		// Less than an hour, return the number of minutes.
		return fmt.Sprintf("%d minutes ago", int(diff.Minutes()))
	case diff < time.Hour*24:
		// Today
		return fmt.Sprintf("today %s:%s", zerofy(t.Hour()), zerofy(t.Minute()))
	case diff < time.Hour*24*2:
		// Yesterday
		return fmt.Sprintf("yesterday %s:%s", zerofy(t.Hour()), zerofy(t.Minute()))
	}

	return t.Format("02.01.2006 15:04")
}

// ShortDMY returns a short date string in the format "day month year".
func ShortDMY(t time.Time) string {
	return fmt.Sprintf("%d %s %d", t.Day(), t.Month().String()[:3], t.Year())
}

// ShortMY returns a short date string in the format "month year".
func ShortMY(t time.Time) string {
	return fmt.Sprintf("%s %d", t.Month().String()[:3], t.Year())
}

// ShortMYHM returns a short date string in the format "month year hour:minute".
func ShortMYHM(t time.Time) string {
	return fmt.Sprintf(
		"%s %d %s:%s",
		t.Month().String()[:3], t.Year(), zerofy(t.Hour()), zerofy(t.Minute()),
	)
}

// zerofy adds a leading zero to single-digit numbers.
func zerofy(x int) string {
	if x < 10 {
		return fmt.Sprintf("0%d", x)
	}
	return fmt.Sprintf("%d", x)
}
