package util

import "time"

// return in KST timezone, truncated to seconds
func Now() time.Time {
	return time.Now().In(time.FixedZone("KST", 9*60*60)).Truncate(time.Second)
}
