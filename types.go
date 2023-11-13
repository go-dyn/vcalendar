package vcalendar

import "time"

type Calendar struct {
	t     time.Time
	Day   int
	Month int
	Year  int
}
type Lunar struct {
	c    Calendar
	Leap bool
}

type Solar struct {
	c Calendar
}

type LuckyHourDetail struct {
	Chi  string
	From int
	To   int
}

type SolarTerm struct {
	longitude int
	name      string
}

type YearEvent struct {
	day   int
	month int
	name  string
}
