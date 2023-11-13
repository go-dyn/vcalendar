package vcalendar

import (
	"fmt"
)

func (s *Solar) String() string {
	return fmt.Sprintf("%s, ngày %02d, tháng %02d, năm %d", s.Weekday(), s.c.Day, s.c.Month, s.c.Year)
}

func (s *Solar) Weekday() string {
	wd := int(s.c.t.Weekday())
	return WEEKS[wd]
}
