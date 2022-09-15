package xtime

import (
	"time"
)

// Today returns a new time truncated by time
func Today(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// Tomorrow returns the time when the next day
func Tomorrow(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location())
}

// ThisWeekday returns the time when the given day of the week most recently started
func ThisWeekday(t time.Time, w time.Weekday) time.Time {
	delta := int(w - t.Weekday())
	if delta > 0 {
		delta -= 7
	}
	return time.Date(t.Year(), t.Month(), t.Day()+delta, 0, 0, 0, 0, t.Location())
}

// NextWeekday returns the time when the given day of the week comes next
func NextWeekday(t time.Time, w time.Weekday) time.Time {
	delta := int(w - t.Weekday())
	if delta <= 0 {
		delta += 7
	}
	return time.Date(t.Year(), t.Month(), t.Day()+delta, 0, 0, 0, 0, t.Location())
}

// BeginningOfMonth returns the time when beginning of this month
func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// NextMonth returns the time when the next month
func NextMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
}
