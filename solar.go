package vcalendar

import (
	"fmt"
)

func (s *Solar) String() string {
	return fmt.Sprintf("%s, ngày %02d, tháng %02d, năm %d", s.Weekday(), s.Date.Day, s.Date.Month, s.Date.Year)
}

func (s *Solar) Weekday() string {
	wd := int(s.Date.Time.Weekday())
	return WEEKS[wd]
}
