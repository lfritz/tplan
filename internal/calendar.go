package internal

import (
	"fmt"
	"strings"
	"time"
)

type Calendar struct {
	weekend map[time.Weekday]bool
	off     map[Date]bool
}

func NewCalendar() *Calendar {
	return &Calendar{
		weekend: make(map[time.Weekday]bool),
		off:     make(map[Date]bool),
	}
}

func (c *Calendar) IsOff(d Date) bool {
	return c.weekend[d.Weekday()] || c.off[d]
}

func (c *Calendar) SetWeekend(day time.Weekday) {
	c.weekend[day] = true
}

func (c *Calendar) SetOff(date Date) {
	c.off[date] = true
}

var weekdayNames = map[time.Weekday]string{
	time.Monday:    "Mon",
	time.Tuesday:   "Tue",
	time.Wednesday: "Wed",
	time.Thursday:  "Thu",
	time.Friday:    "Fri",
	time.Saturday:  "Sat",
	time.Sunday:    "Sun",
}

func PrintWeekday(d time.Weekday) string {
	return weekdayNames[d]
}

var weekdays = map[string]time.Weekday{
	"Mon": time.Monday,
	"Tue": time.Tuesday,
	"Wed": time.Wednesday,
	"Thu": time.Thursday,
	"Fri": time.Friday,
	"Sat": time.Saturday,
	"Sun": time.Sunday,
}

func ParseWeekendSpec(s string) (days []time.Weekday, err error) {
	days = []time.Weekday{}
	if s == "" {
		return
	}
	parts := strings.Split(s, ", ")
	for _, part := range parts {
		day, ok := weekdays[part]
		if !ok {
			return nil, fmt.Errorf("invalid weekday: %q", part)
		}
		days = append(days, day)
	}
	return
}

func ParseDateSpec(s string) (days []Date, err error) {
	days = []Date{}
	if s == "" {
		return
	}
	parts := strings.Split(s, ", ")
	for _, part := range parts {
		day, err := ParseDate(part)
		if err != nil {
			return nil, fmt.Errorf("invalid date: %q", part)
		}
		days = append(days, day)
	}
	return
}
