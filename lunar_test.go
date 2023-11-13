package vcalendar

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

var lunarTestCases = []struct {
	Name                    string
	date                    time.Time
	expectedDayAlias        string
	expectedMonthAlias      string
	expectedYearAlias       string
	expectedWeekday         int
	expectedWeekdayDisplay  string
	expectedLuckyHour       string
	expectedLuckyHourDetail []LuckyHourDetail
	expectedEvent           []YearEvent
	expectedSolarTerm       SolarTerm
}{
	{
		Name:                   "Test Case 1",
		date:                   time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Kỷ Dậu",
		expectedMonthAlias:     "Mậu Ngọ",
		expectedYearAlias:      "Quý Mão",
		expectedWeekday:        2,
		expectedWeekdayDisplay: "Thứ Ba",
		expectedLuckyHour:      "101100110100",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			Longitude: 75,
			Name:      "Mang chủng",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		},
	},
	{
		Name:                   "Test Case 2",
		date:                   time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Canh Thìn",
		expectedMonthAlias:     "Ất Mão Nhuận",
		expectedYearAlias:      "Quý Mão",
		expectedWeekday:        4,
		expectedWeekdayDisplay: "Thứ Năm",
		expectedLuckyHour:      "001011001101",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			Longitude: 0,
			Name:      "Xuân phân",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		},
	},
	{
		Name:                   "Test Case 3",
		date:                   time.Date(2030, time.February, 12, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Mậu Dần",
		expectedMonthAlias:     "Mậu Dần",
		expectedYearAlias:      "Canh Tuất",
		expectedWeekday:        2,
		expectedWeekdayDisplay: "Thứ Ba",
		expectedLuckyHour:      "110011010010",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			Longitude: 315,
			Name:      "Lập xuân",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
		},
	},
	{
		Name:                   "Test Case 4",
		date:                   time.Date(2004, time.November, 30, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Quý Sửu",
		expectedMonthAlias:     "Ất Hợi",
		expectedYearAlias:      "Giáp Thân",
		expectedWeekday:        2,
		expectedWeekdayDisplay: "Thứ Ba",
		expectedLuckyHour:      "001101001011",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			Longitude: 240,
			Name:      "Tiểu tuyết",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		},
	},
	{
		Name:                   "Test Case 5",
		date:                   time.Date(1997, time.January, 19, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Tân Dậu",
		expectedMonthAlias:     "Tân Sửu",
		expectedYearAlias:      "Bính Tý",
		expectedWeekday:        0,
		expectedWeekdayDisplay: "Chủ Nhật",
		expectedLuckyHour:      "101100110100",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			Longitude: 285,
			Name:      "Tiểu hàn",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		},
	},
	{
		Name:                   "Test Case 6",
		date:                   time.Date(2024, time.February, 10, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Giáp Thìn",
		expectedMonthAlias:     "Bính Dần",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        6,
		expectedWeekdayDisplay: "Thứ Bảy",
		expectedLuckyHour:      "001011001101",
		expectedEvent: []YearEvent{{
			Day:   1,
			Month: 1,
			Name:  "Tết Nguyên Đán",
			Type:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			Longitude: 315,
			Name:      "Lập xuân",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		},
	},
	{
		Name:                   "Test Case 7",
		date:                   time.Date(2024, time.February, 24, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Mậu Ngọ",
		expectedMonthAlias:     "Bính Dần",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        6,
		expectedWeekdayDisplay: "Thứ Bảy",
		expectedLuckyHour:      "110100101100",
		expectedEvent: []YearEvent{{
			Day:   15,
			Month: 1,
			Name:  "Rằm tháng Giêng",
			Type:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			Longitude: 330,
			Name:      "Vũ Thủy",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		},
	},
	{
		Name:                   "Test Case 8",
		date:                   time.Date(2024, time.April, 18, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Nhâm Tý",
		expectedMonthAlias:     "Mậu Thìn",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        4,
		expectedWeekdayDisplay: "Thứ Năm",
		expectedLuckyHour:      "110100101100",
		expectedEvent: []YearEvent{{
			Day:   10,
			Month: 3,
			Name:  "Giỗ Tổ Hùng Vương",
			Type:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			Longitude: 15,
			Name:      "Thanh minh",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		},
	},
	{
		Name:                   "Test Case 9",
		date:                   time.Date(2024, time.May, 22, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Bính Tuất",
		expectedMonthAlias:     "Kỷ Tỵ",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        3,
		expectedWeekdayDisplay: "Thứ Tư",
		expectedLuckyHour:      "001011001101",
		expectedEvent: []YearEvent{{
			Day:   15,
			Month: 4,
			Name:  "Phật Đản",
			Type:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			Longitude: 60,
			Name:      "Tiểu mãn",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Thân",
				From: 15,
				To:   17,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		},
	},
	{
		Name:                   "Test Case 10",
		date:                   time.Date(2024, time.June, 10, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Ất Tỵ",
		expectedMonthAlias:     "Canh Ngọ",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        1,
		expectedWeekdayDisplay: "Thứ Hai",
		expectedLuckyHour:      "010010110011",
		expectedEvent: []YearEvent{{
			Day:   5,
			Month: 5,
			Name:  "Lễ Đoan Ngọ",
			Type:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			Longitude: 75,
			Name:      "Mang chủng",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
			{
				Chi:  "Hợi",
				From: 21,
				To:   23,
			},
		},
	},
	{
		Name:                   "Test Case 11",
		date:                   time.Date(2024, time.August, 18, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Giáp Dần",
		expectedMonthAlias:     "Nhâm Thân",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        0,
		expectedWeekdayDisplay: "Chủ Nhật",
		expectedLuckyHour:      "110011010010",
		expectedEvent: []YearEvent{{
			Day:   15,
			Month: 7,
			Name:  "Vu Lan",
			Type:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			Longitude: 135,
			Name:      "Lập thu",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
		},
	},
	{
		Name:                   "Test Case 12",
		date:                   time.Date(2024, time.September, 17, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Giáp Thân",
		expectedMonthAlias:     "Quý Dậu",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        2,
		expectedWeekdayDisplay: "Thứ Ba",
		expectedLuckyHour:      "110011010010",
		expectedEvent: []YearEvent{{
			Day:   15,
			Month: 8,
			Name:  "Tết Trung Thu",
			Type:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			Longitude: 165,
			Name:      "Bạch lộ",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Sửu",
				From: 1,
				To:   3,
			},
			{
				Chi:  "Thìn",
				From: 7,
				To:   9,
			},
			{
				Chi:  "Tỵ",
				From: 9,
				To:   11,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Tuất",
				From: 19,
				To:   21,
			},
		},
	},
	{
		Name:                   "Test Case 13",
		date:                   time.Date(2025, time.January, 22, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Tân Mão",
		expectedMonthAlias:     "Đinh Sửu",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        3,
		expectedWeekdayDisplay: "Thứ Tư",
		expectedLuckyHour:      "101100110100",
		expectedEvent: []YearEvent{{
			Day:   23,
			Month: 12,
			Name:  "Ông Táo chầu trời",
			Type:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			Longitude: 300,
			Name:      "Đại hàn",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				Chi:  "Tý",
				From: 23,
				To:   1,
			},
			{
				Chi:  "Dần",
				From: 3,
				To:   5,
			},
			{
				Chi:  "Mão",
				From: 5,
				To:   7,
			},
			{
				Chi:  "Ngọ",
				From: 11,
				To:   13,
			},
			{
				Chi:  "Mùi",
				From: 13,
				To:   15,
			},
			{
				Chi:  "Dậu",
				From: 17,
				To:   19,
			},
		},
	},
}

func TestToString(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.String()
			expect := fmt.Sprintf("ngày %s, tháng %s, năm %s", tc.expectedDayAlias, tc.expectedMonthAlias, tc.expectedYearAlias)
			if result != expect {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.Name, tc.expectedDayAlias, result)
			}
		})
	}
}

func TestDayAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.DayAlias()

			if result != tc.expectedDayAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.Name, tc.expectedDayAlias, result)
			}
		})
	}
}

func TestMonthAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.MonthAlias()

			if result != tc.expectedMonthAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.Name, tc.expectedMonthAlias, result)
			}
		})
	}
}

func TestYearAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.YearAlias()

			if result != tc.expectedYearAlias {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.Name, tc.expectedYearAlias, result)
			}
		})
	}
}

func TestToSolar(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			s := l.ToSolar()
			result := s.Date.Day

			if result != tc.date.Day() {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %d", tc.Name, tc.expectedYearAlias, result)
			}
		})
	}
}

func TestWeekday(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			// this is wrong for test case
			result := l.Weekday()

			if result != tc.expectedWeekday {
				t.Errorf("Unexpected result for %s. Expected: %d, got: %d", tc.Name, tc.expectedWeekday, result)
			}
		})
	}
}
func TestWeekdayDisplay(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			// this is wrong for test case
			result := l.WeekdayDisplay()

			if result != tc.expectedWeekdayDisplay {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.Name, tc.expectedWeekdayDisplay, result)
			}
		})
	}
}

func TestNext(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.Next()

			if result != nil {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.Name, nil, result)
			}
		})
	}
}

func TestLuckyHour(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.LuckyHour()

			if result != tc.expectedLuckyHour {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.Name, tc.expectedLuckyHour, result)
			}
		})
	}
}

func TestLuckyHourDetail(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.LuckyHourDetail()
			if !reflect.DeepEqual(result, tc.expectedLuckyHourDetail) {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.Name, tc.expectedLuckyHourDetail, result)
			}

		})
	}
}

func TestEvents(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.Events()
			if !reflect.DeepEqual(result, tc.expectedEvent) {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.Name, tc.expectedEvent, result)
			}

		})
	}
}

func TestSolarTerms(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			l := New(tc.date)
			result := l.SolarTerms()
			if !reflect.DeepEqual(result, tc.expectedSolarTerm) {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.Name, tc.expectedSolarTerm, result)
			}

		})
	}
}
