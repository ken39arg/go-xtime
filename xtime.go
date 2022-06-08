package xtime

import (
	"context"
	"time"
)

// for test
var nowFunc func() time.Time = time.Now

type ctxKey int

const (
	ctxKeyFix ctxKey = iota + 1
	ctxKeyRelative
)

func fixTime(ctx context.Context) (time.Time, bool) {
	if v := ctx.Value(ctxKeyFix); v != nil {
		return v.(time.Time), true
	}
	return time.Time{}, false
}

func relative(ctx context.Context) (time.Duration, bool) {
	if v := ctx.Value(ctxKeyRelative); v != nil {
		return v.(time.Duration), true
	}
	return 0, false
}

// Now returns the current time via context
//
// if context doesn't have time value then returns time.Now()
func Now(ctx context.Context) time.Time {
	if ft, ok := fixTime(ctx); ok {
		return ft
	}

	if d, ok := relative(ctx); ok {
		return nowFunc().Add(d)
	}

	return nowFunc()
}

// Assume returns a function that takes over settings to another context
func Assume(ctx context.Context) func(context.Context) context.Context {
	if ft, ok := fixTime(ctx); ok {
		return func(ctx context.Context) context.Context {
			return FixTime(ctx, ft)
		}
	}

	if d, ok := relative(ctx); ok {
		return func(ctx context.Context) context.Context {
			return SetRelativeDuration(ctx, d)
		}
	}
	return func(ctx context.Context) context.Context {
		return ctx
	}
}

// FixTime fixes the value returned by Now(ctx)
func FixTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, ctxKeyFix, t)
}

// AddFixTime adds a Duration to the Now() result time and re-fixes the Now() result.
func AddFixTime(ctx context.Context, d time.Duration) context.Context {
	return FixTime(ctx, Now(ctx).Add(d))
}

// SetRelativeTime changes the current time
// FixTime is fixed, while SetRelativeTime moves forward in time from the time it is set.
func SetRelativeTime(ctx context.Context, t time.Time) context.Context {
	return SetRelativeDuration(ctx, t.Sub(nowFunc()))
}

// AddRelativeTime extends RelativeTime
func AddRelativeTime(ctx context.Context, d time.Duration) context.Context {
	return SetRelativeDuration(ctx, GetRelativeDuration(ctx)+d)
}

// GetRelativeDuration returns relative duration
func GetRelativeDuration(ctx context.Context) time.Duration {
	if d, ok := relative(ctx); ok {
		return d
	}
	return 0
}

// SetRelativeDuration sets the relative time difference
func SetRelativeDuration(ctx context.Context, duration time.Duration) context.Context {
	return context.WithValue(ctx, ctxKeyRelative, duration)
}
