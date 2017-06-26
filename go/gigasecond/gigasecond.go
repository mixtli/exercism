package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds 10^9 seconds to the argument
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1E9 * time.Second)
}
