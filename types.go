package vcalendar

import "time"

type Locale struct {
	Code string
	Name string
}

type Calendar struct {
	Time     time.Time
	Day      int
	Month    int
	Year     int
	Type     int
	Language Locale
}
type Lunar struct {
	Date Calendar
	Leap bool
}

type Solar struct {
	Date Calendar
}

type LuckyHourDetail struct {
	Chi  string
	From int
	To   int
}

type SolarTerm struct {
	Longitude int
	Name      string
}
type YearEvent struct {
	Day   int
	Month int
	Name  string
	Type  int // 1: solar, 2: lunar
}
