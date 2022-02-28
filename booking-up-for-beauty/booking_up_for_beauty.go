package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)

	return parsedDate
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)

	return time.Now().After(parsedDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)

	return parsedDate.Hour() < 18 && parsedDate.Hour() > 11
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		parsedDate.Weekday(),
		parsedDate.Month(),
		parsedDate.Day(),
		parsedDate.Year(),
		parsedDate.Hour(),
		parsedDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
