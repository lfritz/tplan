package internal_test

import (
	"testing"
	"time"

	"github.com/lfritz/tplan/internal"
)

func TestNewDate(t *testing.T) {
	year, month, day := 2022, 6, 12
	weekday := time.Sunday
	d := internal.NewDate(year, month, day)
	if d.Year() != year {
		t.Errorf("d.Year() == %d, want %d", d.Year(), year)
	}
	if d.Month() != month {
		t.Errorf("d.Month() == %d, want %d", d.Month(), month)
	}
	if d.Day() != day {
		t.Errorf("d.Day() == %d, want %d", d.Day(), day)
	}
	if d.Weekday() != weekday {
		t.Errorf("d.Weekday() == %s, want %s", d.Weekday(), weekday)
	}
}

func TestParseDate(t *testing.T) {
	input := "2022-06-12"
	got, err := internal.ParseDate(input)
	if err != nil {
		t.Fatalf("ParseDate(%q) returned error: %v", input, err)
	}
	want := internal.NewDate(2022, 6, 12)
	if got != want {
		t.Errorf("ParseDate(%q) == %v, want %v", input, got, want)
	}

	input = "2022-06-31"
	_, err = internal.ParseDate(input)
	if err == nil {
		t.Errorf("ParseDate(%q) did not return error", input)
	}
}

var dateTestCases = []struct {
	a, b       internal.Date
	difference int
}{
	{internal.NewDate(2022, 6, 8), internal.NewDate(2022, 6, 8), 0},
	{internal.NewDate(2022, 6, 8), internal.NewDate(2022, 6, 9), +1},
	{internal.NewDate(2022, 6, 8), internal.NewDate(2022, 6, 7), -1},
	{internal.NewDate(2022, 6, 8), internal.NewDate(2022, 6, 13), +3},
	{internal.NewDate(2022, 6, 8), internal.NewDate(2022, 6, 3), -3},
	{internal.NewDate(2022, 6, 12), internal.NewDate(2022, 6, 12), 0},
	{internal.NewDate(2022, 6, 12), internal.NewDate(2022, 6, 13), +1},
	{internal.NewDate(2022, 6, 12), internal.NewDate(2022, 6, 10), -1},
}

func TestDateAdvance(t *testing.T) {
	cal := internal.NewCalendar()
	cal.SetWeekend(time.Saturday)
	cal.SetWeekend(time.Sunday)
	for _, c := range dateTestCases {
		got := c.a.Advance(c.difference, cal)
		if got != c.b {
			t.Errorf("(%v).Advance(%d, cal) == %v, want %v", c.a, c.difference, got, c.b)
		}
	}
}

func TestDateSub(t *testing.T) {
	cal := internal.NewCalendar()
	cal.SetWeekend(time.Saturday)
	cal.SetWeekend(time.Sunday)
	for _, c := range dateTestCases {
		got := c.b.Sub(c.a, cal)
		if got != c.difference {
			t.Errorf("(%v).Sub(%v, cal) == %d, want %d", c.b, c.a, got, c.difference)
		}
	}
}
