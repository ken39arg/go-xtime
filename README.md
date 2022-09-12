# go-xtime

this package is time.Now with context.Context

and some useful functions for time.Time

## Usage

```go
func SomeFunc(ctx context.Context) {
    // now := time.Now()
    now := xtime.Now(ctx) // modify time.Now

    return now.Add(time.Second)
}

// in test

func TestSomeFunc(t *testing.T) {
    ctx := xtime.FixTime(context.Background(),  time.Date(2021, 2, 3, 4, 5, 6, 0, time.UTC))

    if SomeFunc(ctx) !=  time.Date(2021, 2, 3, 4, 5, 7, 0, time.UTC) { // always the same result
        t.Error("not match")
    }
}
```
