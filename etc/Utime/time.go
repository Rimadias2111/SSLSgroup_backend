package Utime

import "time"

func Now() time.Time {
	loc, _ := time.LoadLocation("America/New_York")
	return time.Now().In(loc)
}
