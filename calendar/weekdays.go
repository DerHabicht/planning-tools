package calendar

import (
	"time"
)

var WeekdayLetters = map[time.Weekday]string{
	time.Monday:    "M",
	time.Tuesday:   "T",
	time.Wednesday: "W",
	time.Thursday:  "H",
	time.Friday:    "F",
	time.Saturday:  "S",
	time.Sunday:    "U",
}
