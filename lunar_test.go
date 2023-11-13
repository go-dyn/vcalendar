package vcalendar

import (
	"fmt"
	"testing"
	"time"
)

var lunarTestCases = []struct {
	name               string
	date               time.Time
	expectedDayAlias   string
	expectedMonthAlias string
	expectedYearAlias  string
	expectedWeekday    string
}{
	{
		name:               "Test Case 1",
		date:               time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Kỷ Dậu",
		expectedMonthAlias: "Mậu Ngọ",
		expectedYearAlias:  "Quý Mão",
		expectedWeekday:    "Thứ Ba",
	},
	{
		name:               "Test Case 2",
		date:               time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Canh Thìn",
		expectedMonthAlias: "Ất Mão Nhuận",
		expectedYearAlias:  "Quý Mão",
		expectedWeekday:    "Thứ Năm",
	},
	{
		name:               "Test Case 3",
		date:               time.Date(2030, time.February, 12, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Mậu Dần",
		expectedMonthAlias: "Mậu Dần",
		expectedYearAlias:  "Canh Tuất",
		expectedWeekday:    "Thứ Ba",
	},
	{
		name:               "Test Case 4",
		date:               time.Date(2004, time.November, 30, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Quý Sửu",
		expectedMonthAlias: "Ất Hợi",
		expectedYearAlias:  "Giáp Thân",
		expectedWeekday:    "Thứ Ba",
	},
	{
		name:               "Test Case 5",
		date:               time.Date(1997, time.January, 19, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:   "Tân Dậu",
		expectedMonthAlias: "Tân Sửu",
		expectedYearAlias:  "Bính Tý",
		expectedWeekday:    "Chủ Nhật",
	},
}

func TestToString(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.String()
			expect := fmt.Sprintf("ngày %s, tháng %s, năm %s", tc.expectedDayAlias, tc.expectedMonthAlias, tc.expectedYearAlias)
			if result != expect {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedDayAlias, result)
			}
		})
	}
}

func TestDayAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.DayAlias()

			if result != tc.expectedDayAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedDayAlias, result)
			}
		})
	}
}

func TestMonthAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.MonthAlias()

			if result != tc.expectedMonthAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedMonthAlias, result)
			}
		})
	}
}

func TestYearAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.YearAlias()

			if result != tc.expectedYearAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedYearAlias, result)
			}
		})
	}
}

func TestToSolar(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			s := l.ToSolar()
			result := s.Day

			if result != tc.date.Day() {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %d", tc.name, tc.expectedYearAlias, result)
			}
		})
	}
}

func TestWeekday(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			// this is wrong for test case
			result := l.Weekday()

			if result != tc.expectedWeekday {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedWeekday, result)
			}
		})
	}
}

func TestNext(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.date)
			result := l.Next()

			if result != nil {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.name, nil, result)
			}
		})
	}
}
