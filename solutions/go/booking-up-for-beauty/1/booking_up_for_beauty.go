package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	dt, err := time.Parse("1/2/2006 15:04:05", date)
    if err == nil{
        return dt
    }
	panic("An error occurred")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    dt, err := time.Parse("January 2, 2006 15:04:05", date)
    if err == nil {
		return time.Now().After(dt)
    }
    panic("An error occurred")
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dt, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if err == nil {
    	return dt.Hour() >= 12 && dt.Hour() < 18
	}
	panic("An error occurred")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    dt, err := time.Parse("1/2/2006 15:04:05", date)
    if err == nil{
	return fmt.Sprintf("You have an appointment on %s", dt.Format("Monday, January 2, 2006, at 15:04."))
        }
    panic("An error occurred")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    now := time.Now()
	return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, now.Location())
}
