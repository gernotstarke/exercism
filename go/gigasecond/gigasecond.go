// solving the exercism "gigasecond" exercise
package gigasecond

import "time"

const gigasecondstring = "1000000000s"

// AddGigasecond adds a gigasecond.
// It determines the moment that would be after a gigasecond has passed.
//A gigasecond is 10^9 (1,000,000,000) seconds.
func AddGigasecond(t time.Time) time.Time {
	gigasecond, _ := time.ParseDuration(gigasecondstring)

	// no error handling... thus the "_" as second param

	return t.Add(gigasecond)
}
