package vcalendar

import "time"

type Locale struct {
	code string
	name string
}

type Calendar struct {
	time.Time
}

type YearEvent struct {
	day   int
	month int
	name  string
	kind  int // 1: solar, 2: lunar
}

// only lunar
type LuckyHourDetail struct {
	chi  string
	from int
	to   int
}
type SolarTerm struct {
	longitude int
	name      string
}
