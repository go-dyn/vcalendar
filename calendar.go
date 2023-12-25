package vcalendar

import "time"

func (c *Calendar) String() string {
	return c.String()
}

func (c *Calendar) Format(format string) string {
	return c.Format(format)
}

func (c *Calendar) GetTime() time.Time {
	return time.Date(c.Year(), c.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
}

func (c *Calendar) GetDay() int {
	return c.Day()
}

func (c *Calendar) GetMonth() time.Month {
	return c.Month()
}

func (c *Calendar) GetIntMonth() int {
	return int(c.Month())
}

func (c *Calendar) GetYear() int {
	return c.Year()
}

func (c *Calendar) GetWeekDay() int {
	return int(c.Weekday())
}

func (c *Calendar) GetLocation() *time.Location {
	return c.Location()
}

func (c *Calendar) GetTimeZone() (name string, offset int) {
	return c.Zone()
}

func (c *Calendar) SetDay(Day int) *Calendar {
	c.AddDate(0, 0, Day)
	return c
}

func (c *Calendar) SetIntMonth(Month int) *Calendar {
	c.AddDate(0, Month, 0)
	return c
}

func (c *Calendar) SetMonth(Month time.Month) *Calendar {
	c.AddDate(0, int(Month), 0)
	return c
}

func (c *Calendar) SetYear(Year int) *Calendar {
	c.AddDate(Year, 0, 0)
	return c
}

func (c *Calendar) SetLocation(location *time.Location) *Calendar {
	c.In(location)
	return c
}

func (c *Calendar) SetTimeZone(name string, offset int) *Calendar {
	c.In(time.FixedZone(name, offset))
	return c
}

func (c *Calendar) IsLeap() bool {
	year := c.GetYear()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func (c *Calendar) GetJd() int {
	jd := DateToJuliusDay(c.Day(), c.GetIntMonth(), c.Year())
	return jd
}

func (c *Calendar) MoveVnLunar() *Calendar {
	_, offset := c.GetTimeZone()
	lD, lM, lY, _ := SolarToLunar(c.GetDay(), c.GetIntMonth(), c.GetYear(), offset)
	c.AddDate(lY, lM, lD)
	return c
}

func (c *Calendar) ToVnLunar() time.Time {
	_, offset := c.GetTimeZone()
	lD, lM, lY, _ := SolarToLunar(c.GetDay(), c.GetIntMonth(), c.GetYear(), offset)
	return time.Date(lY, time.Month(lM), lD, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.GetLocation())
}

func (c *Calendar) FromSolar(time time.Time) *Calendar {
	c.AddDate(time.Day(), int(time.Month()), time.Year())

	return c
}

func (c *Calendar) FromVnLunar(time time.Time, leap int) *Calendar {
	_, offset := time.Zone()

	sD, sM, sY := LunarToSolar(time.Day(), int(time.Month()), time.Year(), leap, offset)
	c.AddDate(sY, sM, sD)
	return c
}
