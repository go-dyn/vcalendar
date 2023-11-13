package vcalendar

func (c *Calendar) GetType() int {
	return c.Type
}

func (c *Calendar) GetDay() int {
	return c.Day
}

func (c *Calendar) GetMonth() int {
	return c.Month
}

func (c *Calendar) GetYear() int {
	return c.Year
}

func (c *Calendar) GetLanguage() Locale {
	return c.Language
}

func (c *Calendar) SetDay(Day int) *Calendar {
	c.Day = Day
	return c
}

func (c *Calendar) SetMonth(Month int) *Calendar {
	c.Month = Month
	return c
}

func (c *Calendar) SetYear(Year int) *Calendar {
	c.Year = Year
	return c
}

func (c *Calendar) SetLanguage(Language Locale) *Calendar {
	c.Language = Language
	return c
}
