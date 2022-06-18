package internal

import "time"

type Date struct {
	t time.Time
}

func NewDate(year, month, day int) Date {
	return Date{
		t: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
	}
}

func ParseDate(s string) (Date, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return Date{}, err
	}
	return Date{t: t}, nil
}

func (d Date) String() string {
	return d.t.Format("2006-01-02")
}

func (d Date) Year() int {
	return d.t.Year()
}

func (d Date) Month() int {
	return int(d.t.Month())
}

func (d Date) Day() int {
	return d.t.Day()
}

func (d Date) Weekday() time.Weekday {
	return d.t.Weekday()
}

func (d Date) IsZero() bool {
	return d.t.IsZero()
}

func (d Date) After(e Date) bool {
	return d.t.After(e.t)
}

func (d Date) Advance(n int, cal *Calendar) Date {
	direction := 1
	count := n
	if n < 0 {
		direction = -1
		count = -n
	}

	for i := 0; i < count; i++ {
		d.t = d.t.AddDate(0, 0, direction)
		for cal.IsOff(d) {
			d.t = d.t.AddDate(0, 0, direction)
		}
	}
	return d
}

func (d Date) Sub(e Date, cal *Calendar) int {
	if e.After(d) {
		return -e.Sub(d, cal)
	}
	count := 0
	for d.After(e) {
		e = e.Advance(1, cal)
		count++
	}
	return count
}
