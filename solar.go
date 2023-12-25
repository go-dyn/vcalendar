package vcalendar

func (s *Calendar) SolarWeekDay() string {
	wd := int(s.GetWeekDay())
	return WEEKS[wd]
}
