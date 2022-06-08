package xtime

import (
	"context"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	def := nowFunc
	defer func() {
		nowFunc = def
	}()

	testTime := func() time.Time {
		return time.Date(2021, 2, 3, 4, 5, 6, 0, time.UTC)
	}
	nowFunc = testTime

	for _, tc := range []struct {
		title string
		ctx   context.Context
		now   time.Time
	}{
		{
			title: "default Now",
			ctx:   context.Background(),
			now:   testTime(),
		},
		{
			title: "FixTime",
			ctx:   FixTime(context.Background(), time.Date(2021, 3, 4, 5, 6, 7, 0, time.UTC)),
			now:   time.Date(2021, 3, 4, 5, 6, 7, 0, time.UTC),
		},
		{
			title: "AddFixTime",
			ctx:   AddFixTime(context.Background(), time.Minute),
			now:   time.Date(2021, 2, 3, 4, 6, 6, 0, time.UTC),
		},
		{
			title: "AddFixTime To FixTime",
			ctx:   AddFixTime(FixTime(context.Background(), time.Date(2021, 3, 4, 5, 6, 7, 0, time.UTC)), time.Minute),
			now:   time.Date(2021, 3, 4, 5, 7, 7, 0, time.UTC),
		},
		{
			title: "SetRelativeDuration",
			ctx:   SetRelativeDuration(context.Background(), 15*time.Second),
			now:   time.Date(2021, 2, 3, 4, 5, 21, 0, time.UTC),
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			now := Now(tc.ctx)
			if !now.Equal(tc.now) {
				t.Errorf("mismatch. got: %s, want: %s", now, tc.now)
			}
		})
	}
}
