package vcalendar

import (
	"testing"
	"time"
)

var testCases = []struct {
	name               string
	date               time.Time
	expectedDayAlias   string
	expectedMonthAlias string
	expectedYearAlias  string
}{
	{
		name:               "Test Case 1",
		date:               time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Kỷ Dậu",
		expectedMonthAlias: "Mậu Ngọ",
		expectedYearAlias:  "Quý Mão",
	},
}

func TestNewVCalendar(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.DayAlias()

			if result != tc.expectedDayAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedDayAlias, result)
			}

			c := NewWithLunar(l.c.Day, l.c.Month, l.c.Year, l.Leap, *l.c.t.Location())

			if c.DayAlias() != tc.expectedDayAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedDayAlias, c.DayAlias())
			}
		})
	}
}
