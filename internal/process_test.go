package internal_test

import (
	"reflect"
	"testing"

	"github.com/lfritz/tplan/internal"
)

func TestParseValidTaskLine(t *testing.T) {
	cases := []struct {
		line string
		want *internal.Task
	}{
		{
			"foo  4",
			&internal.Task{"foo", 4, internal.Date{}, internal.Date{}, 0},
		},
		{
			"foo bar     123     ",
			&internal.Task{"foo bar", 123, internal.Date{}, internal.Date{}, 0},
		},
		{
			"foo  4  2022-06-13",
			&internal.Task{"foo", 4, internal.Date{}, internal.Date{}, 0},
		},
		{
			"foo  4  2022-06-13  2022-06-14",
			&internal.Task{"foo", 4, internal.NewDate(2022, 6, 13), internal.NewDate(2022, 6, 14), 0},
		},
		{
			"foo  4  2022-06-13  2022-06-14  5 (+1)",
			&internal.Task{"foo", 4, internal.NewDate(2022, 6, 13), internal.NewDate(2022, 6, 14), 0},
		},
	}
	for _, c := range cases {
		got, err := internal.ParseTaskLine(c.line)
		if err != nil {
			t.Errorf("ParseTaskLine(%q) returned error: %v", c.line, err)
		} else if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ParseTaskLine(%q) == %v, want %v", c.line, got, c.want)
		}
	}
}

func TestParseInvalidTaskLine(t *testing.T) {
	cases := []string{
		"foo bar baz",
		"foo  abc",
		"foo  4  2022-06-31  2022-07-01",
		"foo  4  2022-07-01  2022-06-31",
	}
	for _, c := range cases {
		_, err := internal.ParseTaskLine(c)
		if err == nil {
			t.Errorf("ParseTaskLine(%q) did not return error", c)
		}
	}
}
