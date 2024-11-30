package Utime

import (
	"log"
	"time"
)

func Now() time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Printf("Error loading location: %v, using UTC as fallback", err)
	}
	return time.Now().In(loc)
}

func Parse(t time.Time) time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Printf("Error loading location: %v, using UTC as fallback", err)
		return time.Time{}
	}

	easternTime := time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
		t.Nanosecond(),
		loc)

	return easternTime
}
