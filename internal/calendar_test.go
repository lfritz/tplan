package internal_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/lfritz/tplan/internal"
)

func TestNewCalendar(t *testing.T) {
	cal := internal.NewCalendar()
	if cal.IsOff(internal.NewDate(2022, 6, 12)) {
		t.Errorf("NewCalendar() returned calendar with date marked as off")
	}
}

func TestCalendarSetWeekend(t *testing.T) {
	cal := internal.NewCalendar()
	cal.SetWeekend(time.Friday)
	cal.SetWeekend(time.Saturday)
	cases := []struct {
		date internal.Date
		want bool
	}{
		{internal.NewDate(2022, 6, 9), false},
		{internal.NewDate(2022, 6, 10), true},
		{internal.NewDate(2022, 6, 11), true},
		{internal.NewDate(2022, 6, 12), false},
	}
	for _, c := range cases {
		got := cal.IsOff(c.date)
		if got != c.want {
			t.Errorf("cal.IsOff(%s) == %v, want %v", c.date, got, c.want)
		}
	}
}

func TestParseWeekendSpec(t *testing.T) {
	cases := []struct {
		spec string
		want []time.Weekday
	}{
		{"", []time.Weekday{}},
		{"Sun", []time.Weekday{time.Sunday}},
		{"Fri, Sat, Sun", []time.Weekday{time.Friday, time.Saturday, time.Sunday}},
	}
	for _, c := range cases {
		got, err := internal.ParseWeekendSpec(c.spec)
		if err != nil {
			t.Fatalf("ParseWeekendSpec(%q) returned error: %v", c.spec, err)
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ParseWeekendSpec(%q) == %v, want %v", c.spec, got, c.want)
		}
	}

	spec := "Xyz"
	_, err := internal.ParseWeekendSpec(spec)
	if err == nil {
		t.Errorf("ParseWeekendSpec(%q) did not return error", spec)
	}
}

func TestParseDateSpec(t *testing.T) {
	june12 := internal.NewDate(2022, 6, 12)
	june13 := internal.NewDate(2022, 6, 13)
	march1 := internal.NewDate(2022, 3, 1)
	cases := []struct {
		spec string
		want []internal.Date
	}{
		{"", []internal.Date{}},
		{"2022-06-12", []internal.Date{june12}},
		{"2022-06-12, 2022-06-13, 2022-03-01", []internal.Date{june12, june13, march1}},
	}
	for _, c := range cases {
		got, err := internal.ParseDateSpec(c.spec)
		if err != nil {
			t.Fatalf("ParseDateSpec(%q) returned error: %v", c.spec, err)
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ParseDateSpec(%q) == %v, want %v", c.spec, got, c.want)
		}
	}

	spec := "2022-04-31"
	_, err := internal.ParseDateSpec(spec)
	if err == nil {
		t.Errorf("ParseDateSpec(%q) did not return error", spec)
	}
}

func TestCalendarSetOff(t *testing.T) {
	cal := internal.NewCalendar()
	cal.SetOff(internal.NewDate(2022, 6, 10))
	cal.SetOff(internal.NewDate(2022, 6, 11))
	cases := []struct {
		date internal.Date
		want bool
	}{
		{internal.NewDate(2022, 6, 9), false},
		{internal.NewDate(2022, 6, 10), true},
		{internal.NewDate(2022, 6, 11), true},
		{internal.NewDate(2022, 6, 12), false},
	}
	for _, c := range cases {
		got := cal.IsOff(c.date)
		if got != c.want {
			t.Errorf("cal.IsOff(%s) == %v, want %v", c.date, got, c.want)
		}
	}
}
