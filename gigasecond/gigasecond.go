package gigasecond

import (
	"time"
)

// AddGigasecond adds gigasecond and returns new time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
