package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	d, _ := time.Parse(layout, date)
	time.Parse(layout, date)
	return d
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"
	d, _ := time.Parse(layout, date)
	return time.Now().After(d)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	//"Thursday, July 25, 2019 13:45:00"
	const layout = "Monday, January 2, 2006 15:04:05"
	d, _ := time.Parse(layout, date)
	h, _, _ := d.Clock()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const layout = "Monday, January 2, 2006, at 15:04"
	d := Schedule(date)
	fd := d.Format(layout)
	return fmt.Sprintf("You have an appointment on %s.", fd)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	y := time.Now().Year()
	a := time.Date(y, time.September, 15, 0, 0, 0, 0, time.UTC)
	return a
}
