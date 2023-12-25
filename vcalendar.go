package vcalendar

import "time"

func NewLuna(t time.Time) *Calendar {
	calendar := &Calendar{}
	calendar.FromSolar(t)
	return calendar
}

func NewLunarFromLunar(d, m, y int, l bool, loc time.Location) *Calendar {

	tzoff := LocationToTimeZone(loc)

	sd, sm, sy := LunarToSolar(d, m, y, BooleanToInt(l), tzoff)
	calendar := &Calendar{}
	time := time.Date(sy, time.Month(sm), sd, 0, 0, 0, 0, &loc)
	calendar.FromVnLunar(time, BooleanToInt(l))

	return calendar
}

func NewSolar(t time.Time) *Calendar {
	calendar := &Calendar{}
	calendar.FromSolar(t)

	return calendar
}

func NewSolarFromLunar(d, m, y int, l bool, loc time.Location) *Calendar {

	tzoff := LocationToTimeZone(loc)

	sd, sm, sy := LunarToSolar(d, m, y, BooleanToInt(l), tzoff)

	calendar := &Calendar{}
	time := time.Date(sy, time.Month(sm), sd, 0, 0, 0, 0, &loc)
	calendar.FromSolar(time)

	return calendar
}
