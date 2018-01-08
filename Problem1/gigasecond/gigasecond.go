// Package gigasecond will add 1,000,000,000 seconds to a time value
// and return the result back to the caller
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond adds 1,000,000,000 secs to the referenced time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1000000000) * time.Second)
}
