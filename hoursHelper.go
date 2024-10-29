package hoursHelper

import (
	"fmt"
	"time"
)

// Hello returns a greeting for the named person.
func Create(start string, end string, interval int, format string) []string {
	var times []string
	theStart, err := time.Parse(time.DateTime, start)
	if err != nil {
		fmt.Println("Could not start time:", err)
	}

	theEnd, err := time.Parse(time.DateTime, end)
	if err != nil {
		fmt.Println("Could not end time:", err)
	}

	isBefore := theStart.After(theEnd)
	if isBefore {
		theEnd = theEnd.Add(time.Hour * 24)
	}

	for theStart.Before(theEnd) || theStart.Equal(theEnd) {
		times = append(times, theStart.Format(format))
		theStart = theStart.Add(time.Minute * time.Duration(interval))
	}

	return times
}