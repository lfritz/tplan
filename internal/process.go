package internal

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var spaceRe = regexp.MustCompile(` {2,}`)

func Process(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)

	// read start line
	if !scanner.Scan() {
		return scanner.Err()
	}
	startDate, err := ParseStartLine(scanner.Text())
	if err != nil {
		return err
	}
	PrintStartLine(writer, startDate)

	// read weekend line
	if !scanner.Scan() {
		return scanner.Err()
	}
	weekend, err := ParseWeekendLine(scanner.Text())
	if err != nil {
		return err
	}
	PrintWeekendLine(writer, weekend)

	// read off line
	if !scanner.Scan() {
		return scanner.Err()
	}
	off, err := ParseOffLine(scanner.Text())
	if err != nil {
		return err
	}
	PrintOffLine(writer, off)

	// prepare calendar
	cal := NewCalendar()
	for _, d := range weekend {
		cal.SetWeekend(d)
	}
	for _, d := range off {
		cal.SetOff(d)
	}

	// read and update task lines
	currentEst := startDate // TODO do I need both?
	currentAct := startDate
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "-") || strings.HasPrefix(line, "=") {
			fmt.Fprintln(writer, line)
			continue
		}
		task, err := ParseTaskLine(line)
		if err != nil {
			return err
		}

		currentEst = currentEst.Advance(task.EstDays, cal)
		task.EstDate = currentEst
		if task.HasActualDate() {
			task.ActDays = task.ActDate.Sub(currentAct, cal)
			currentEst = task.ActDate
			currentAct = task.ActDate
		}

		PrintTaskLine(writer, task)
	}
	return scanner.Err()
}

func ParseStartLine(s string) (Date, error) {
	s = strings.TrimPrefix(s, "start: ")
	s = strings.TrimSpace(s)
	return ParseDate(s)
}

func PrintStartLine(w io.Writer, d Date) {
	fmt.Fprintf(w, "start: %s\n", d)
}

func ParseWeekendLine(s string) ([]time.Weekday, error) {
	s = strings.TrimPrefix(s, "weekend: ")
	s = strings.TrimSpace(s)
	return ParseWeekendSpec(s)
}

func PrintWeekendLine(w io.Writer, days []time.Weekday) {
	names := []string{}
	for _, day := range days {
		names = append(names, PrintWeekday(day))
	}
	fmt.Fprintf(w, "weekend: %s\n", strings.Join(names, ", "))
}

func ParseOffLine(s string) ([]Date, error) {
	s = strings.TrimPrefix(s, "off: ")
	s = strings.TrimSpace(s)
	return ParseDateSpec(s)
}

func PrintOffLine(w io.Writer, dates []Date) {
	list := []string{}
	for _, date := range dates {
		list = append(list, date.String())
	}
	fmt.Fprintf(w, "off: %s\n", strings.Join(list, ", "))
}

type Task struct {
	Description string
	EstDays     int
	EstDate     Date
	ActDate     Date
	ActDays     int
}

func (t *Task) HasEstimatedDate() bool {
	return !t.EstDate.IsZero()
}

func (t *Task) HasActualDate() bool {
	return !t.ActDate.IsZero()
}

func ParseTaskLine(s string) (*Task, error) {
	parts := spaceRe.Split(s, 5)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid task line: %q", s)
	}

	task := new(Task)
	task.Description = parts[0]
	estDays, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	task.EstDays = estDays
	if len(parts) > 3 {
		estDate, err := ParseDate(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid date: %q", parts[3])
		}
		task.EstDate = estDate

		actDate, err := ParseDate(parts[3])
		if err != nil {
			return nil, fmt.Errorf("invalid date: %q", parts[3])
		}
		task.ActDate = actDate
	}

	return task, nil
}

func PrintTaskLine(w io.Writer, task *Task) {
	fmt.Fprintf(w, "%-40s  %d  %s", task.Description, task.EstDays, task.EstDate)
	if task.HasActualDate() {
		fmt.Fprintf(w, "  %s  %d", task.ActDate, task.ActDays)
	}
	fmt.Fprintln(w)
}
