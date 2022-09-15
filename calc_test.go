package xtime

import (
	"testing"
	"time"
)

var (
	jst = time.FixedZone("JST", 9*60*60)
	pst = time.FixedZone("PST", -8*60*60)
)

func TestToday(t *testing.T) {
	for _, tc := range []struct {
		in       time.Time
		expected time.Time
	}{
		{
			in:       time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
			expected: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, jst),
			expected: time.Date(2022, 2, 3, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, pst),
			expected: time.Date(2022, 2, 3, 0, 0, 0, 0, pst),
		},
	} {
		got := Today(tc.in)
		if !got.Equal(tc.expected) {
			t.Errorf("%s returns different time %s [want: %s]", tc.in, got, tc.expected)
		}
	}
}

func TestTomorrow(t *testing.T) {
	for _, tc := range []struct {
		in       time.Time
		expected time.Time
	}{
		{
			in:       time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
			expected: time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, jst),
			expected: time.Date(2022, 2, 4, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2024, 2, 28, 4, 5, 6, 7, pst), // leapday
			expected: time.Date(2024, 2, 29, 0, 0, 0, 0, pst),
		},
		{
			in:       time.Date(2024, 2, 29, 4, 5, 6, 7, jst),
			expected: time.Date(2024, 3, 1, 0, 0, 0, 0, jst),
		},
	} {
		got := Tomorrow(tc.in)
		if !got.Equal(tc.expected) {
			t.Errorf("%s returns different time %s [want: %s]", tc.in, got, tc.expected)
		}
	}
}

func TestThisWeekday(t *testing.T) {
	for _, tc := range []struct {
		in       time.Time
		wd       time.Weekday
		expected time.Time
	}{
		{
			in:       time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC), // sunday
			wd:       time.Sunday,
			expected: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			in:       time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC), // sunday
			wd:       time.Monday,
			expected: time.Date(2021, 12, 27, 0, 0, 0, 0, time.UTC),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, jst), // thursday
			wd:       time.Sunday,
			expected: time.Date(2022, 1, 30, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, jst), // thursday
			wd:       time.Monday,
			expected: time.Date(2022, 1, 31, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2024, 3, 2, 4, 5, 6, 7, pst), // leapday
			wd:       time.Friday,
			expected: time.Date(2024, 3, 1, 0, 0, 0, 0, pst),
		},
		{
			in:       time.Date(2024, 3, 2, 4, 5, 6, 7, pst), // leapday
			wd:       time.Sunday,
			expected: time.Date(2024, 2, 25, 0, 0, 0, 0, pst),
		},
	} {
		got := ThisWeekday(tc.in, tc.wd)
		if !got.Equal(tc.expected) {
			t.Errorf("%s returns different time %s by %s [want: %s]", tc.in, got, tc.wd, tc.expected)
		}
	}
}
func TestNextWeekday(t *testing.T) {
	for _, tc := range []struct {
		in       time.Time
		wd       time.Weekday
		expected time.Time
	}{
		{
			in:       time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC), // sunday
			wd:       time.Sunday,
			expected: time.Date(2022, 1, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			in:       time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC), // sunday
			wd:       time.Saturday,
			expected: time.Date(2022, 1, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, jst), // thursday
			wd:       time.Sunday,
			expected: time.Date(2022, 2, 6, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, jst), // thursday
			wd:       time.Monday,
			expected: time.Date(2022, 2, 7, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2022, 12, 28, 4, 5, 6, 7, jst), // wednesday
			wd:       time.Sunday,
			expected: time.Date(2023, 1, 1, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2022, 12, 28, 4, 5, 6, 7, jst), // wednesday
			wd:       time.Monday,
			expected: time.Date(2023, 1, 2, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2024, 2, 27, 4, 5, 6, 7, pst), // leapday
			wd:       time.Friday,
			expected: time.Date(2024, 3, 1, 0, 0, 0, 0, pst),
		},
		{
			in:       time.Date(2024, 2, 27, 4, 5, 6, 7, pst), // leapday
			wd:       time.Sunday,
			expected: time.Date(2024, 3, 3, 0, 0, 0, 0, pst),
		},
	} {
		got := NextWeekday(tc.in, tc.wd)
		if !got.Equal(tc.expected) {
			t.Errorf("%s returns different time %s by %s [want: %s]", tc.in, got, tc.wd, tc.expected)
		}
	}
}

func TestBeginningOfMonth(t *testing.T) {
	for _, tc := range []struct {
		in       time.Time
		expected time.Time
	}{
		{
			in:       time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
			expected: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, jst),
			expected: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2024, 2, 29, 4, 5, 6, 7, pst), // leapday
			expected: time.Date(2024, 2, 1, 0, 0, 0, 0, pst),
		},
	} {
		got := BeginningOfMonth(tc.in)
		if !got.Equal(tc.expected) {
			t.Errorf("%s returns different time %s [want: %s]", tc.in, got, tc.expected)
		}
	}
}

func TestNextMonth(t *testing.T) {
	for _, tc := range []struct {
		in       time.Time
		expected time.Time
	}{
		{
			in:       time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
			expected: time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			in:       time.Date(2022, 2, 3, 4, 5, 6, 7, jst),
			expected: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2022, 12, 3, 4, 5, 6, 7, jst),
			expected: time.Date(2023, 1, 1, 0, 0, 0, 0, jst),
		},
		{
			in:       time.Date(2024, 2, 28, 4, 5, 6, 7, pst), // leapday
			expected: time.Date(2024, 3, 1, 0, 0, 0, 0, pst),
		},
	} {
		got := NextMonth(tc.in)
		if !got.Equal(tc.expected) {
			t.Errorf("%s returns different time %s [want: %s]", tc.in, got, tc.expected)
		}
	}
}
