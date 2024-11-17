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
