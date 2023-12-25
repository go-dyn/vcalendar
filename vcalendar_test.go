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

func TestNewVvCalendar(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			result := l.DayAlias()

			if result != tc.expectedDayAlias {
				t.Errorf("Unexpected result for %v-%v: %s. Expected: %s, got: %s", l.DayCan(), tc.date, tc.name, tc.expectedDayAlias, result)
			}

			c := NewLunarFromLunar(l.GetDay(), l.GetIntMonth(), l.GetYear(), l.IsLeap(), *l.GetLocation())

			if c.DayAlias() != tc.expectedDayAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedDayAlias, c.DayAlias())
			}
		})
	}
}
