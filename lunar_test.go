package vcalendar

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

var lunarTestCases = []struct {
	name                    string
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
		name:                   "Test Case 1",
		date:                   time.Date(2023, time.June, 20, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Kỷ Dậu",
		expectedMonthAlias:     "Mậu Ngọ",
		expectedYearAlias:      "Quý Mão",
		expectedWeekday:        2,
		expectedWeekdayDisplay: "Thứ Ba",
		expectedLuckyHour:      "101100110100",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			longitude: 75,
			name:      "Mang chủng",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Tý",
				from: 23,
				to:   1,
			},
			{
				chi:  "Dần",
				from: 3,
				to:   5,
			},
			{
				chi:  "Mão",
				from: 5,
				to:   7,
			},
			{
				chi:  "Ngọ",
				from: 11,
				to:   13,
			},
			{
				chi:  "Mùi",
				from: 13,
				to:   15,
			},
			{
				chi:  "Dậu",
				from: 17,
				to:   19,
			},
		},
	},
	{
		name:                   "Test Case 2",
		date:                   time.Date(2023, time.March, 23, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Canh Thìn",
		expectedMonthAlias:     "Ất Mão Nhuận",
		expectedYearAlias:      "Quý Mão",
		expectedWeekday:        4,
		expectedWeekdayDisplay: "Thứ Năm",
		expectedLuckyHour:      "001011001101",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			longitude: 0,
			name:      "Xuân phân",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Dần",
				from: 3,
				to:   5,
			},
			{
				chi:  "Thìn",
				from: 7,
				to:   9,
			},
			{
				chi:  "Tỵ",
				from: 9,
				to:   11,
			},
			{
				chi:  "Thân",
				from: 15,
				to:   17,
			},
			{
				chi:  "Dậu",
				from: 17,
				to:   19,
			},
			{
				chi:  "Hợi",
				from: 21,
				to:   23,
			},
		},
	},
	{
		name:                   "Test Case 3",
		date:                   time.Date(2030, time.February, 12, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Mậu Dần",
		expectedMonthAlias:     "Mậu Dần",
		expectedYearAlias:      "Canh Tuất",
		expectedWeekday:        2,
		expectedWeekdayDisplay: "Thứ Ba",
		expectedLuckyHour:      "110011010010",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			longitude: 315,
			name:      "Lập xuân",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Tý",
				from: 23,
				to:   1,
			},
			{
				chi:  "Sửu",
				from: 1,
				to:   3,
			},
			{
				chi:  "Thìn",
				from: 7,
				to:   9,
			},
			{
				chi:  "Tỵ",
				from: 9,
				to:   11,
			},
			{
				chi:  "Mùi",
				from: 13,
				to:   15,
			},
			{
				chi:  "Tuất",
				from: 19,
				to:   21,
			},
		},
	},
	{
		name:                   "Test Case 4",
		date:                   time.Date(2004, time.November, 30, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Quý Sửu",
		expectedMonthAlias:     "Ất Hợi",
		expectedYearAlias:      "Giáp Thân",
		expectedWeekday:        2,
		expectedWeekdayDisplay: "Thứ Ba",
		expectedLuckyHour:      "001101001011",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			longitude: 240,
			name:      "Tiểu tuyết",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Dần",
				from: 3,
				to:   5,
			},
			{
				chi:  "Mão",
				from: 5,
				to:   7,
			},
			{
				chi:  "Tỵ",
				from: 9,
				to:   11,
			},
			{
				chi:  "Thân",
				from: 15,
				to:   17,
			},
			{
				chi:  "Tuất",
				from: 19,
				to:   21,
			},
			{
				chi:  "Hợi",
				from: 21,
				to:   23,
			},
		},
	},
	{
		name:                   "Test Case 5",
		date:                   time.Date(1997, time.January, 19, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Tân Dậu",
		expectedMonthAlias:     "Tân Sửu",
		expectedYearAlias:      "Bính Tý",
		expectedWeekday:        0,
		expectedWeekdayDisplay: "Chủ Nhật",
		expectedLuckyHour:      "101100110100",
		expectedEvent:          []YearEvent{},
		expectedSolarTerm: SolarTerm{
			longitude: 285,
			name:      "Tiểu hàn",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Tý",
				from: 23,
				to:   1,
			},
			{
				chi:  "Dần",
				from: 3,
				to:   5,
			},
			{
				chi:  "Mão",
				from: 5,
				to:   7,
			},
			{
				chi:  "Ngọ",
				from: 11,
				to:   13,
			},
			{
				chi:  "Mùi",
				from: 13,
				to:   15,
			},
			{
				chi:  "Dậu",
				from: 17,
				to:   19,
			},
		},
	},
	{
		name:                   "Test Case 6",
		date:                   time.Date(2024, time.February, 10, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Giáp Thìn",
		expectedMonthAlias:     "Bính Dần",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        6,
		expectedWeekdayDisplay: "Thứ Bảy",
		expectedLuckyHour:      "001011001101",
		expectedEvent: []YearEvent{{
			day:   1,
			month: 1,
			name:  "Tết Nguyên Đán",
			kind:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			longitude: 315,
			name:      "Lập xuân",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Dần",
				from: 3,
				to:   5,
			},
			{
				chi:  "Thìn",
				from: 7,
				to:   9,
			},
			{
				chi:  "Tỵ",
				from: 9,
				to:   11,
			},
			{
				chi:  "Thân",
				from: 15,
				to:   17,
			},
			{
				chi:  "Dậu",
				from: 17,
				to:   19,
			},
			{
				chi:  "Hợi",
				from: 21,
				to:   23,
			},
		},
	},
	{
		name:                   "Test Case 7",
		date:                   time.Date(2024, time.February, 24, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Mậu Ngọ",
		expectedMonthAlias:     "Bính Dần",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        6,
		expectedWeekdayDisplay: "Thứ Bảy",
		expectedLuckyHour:      "110100101100",
		expectedEvent: []YearEvent{{
			day:   15,
			month: 1,
			name:  "Rằm tháng Giêng",
			kind:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			longitude: 330,
			name:      "Vũ Thủy",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Tý",
				from: 23,
				to:   1,
			},
			{
				chi:  "Sửu",
				from: 1,
				to:   3,
			},
			{
				chi:  "Mão",
				from: 5,
				to:   7,
			},
			{
				chi:  "Ngọ",
				from: 11,
				to:   13,
			},
			{
				chi:  "Thân",
				from: 15,
				to:   17,
			},
			{
				chi:  "Dậu",
				from: 17,
				to:   19,
			},
		},
	},
	{
		name:                   "Test Case 8",
		date:                   time.Date(2024, time.April, 18, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Nhâm Tý",
		expectedMonthAlias:     "Mậu Thìn",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        4,
		expectedWeekdayDisplay: "Thứ Năm",
		expectedLuckyHour:      "110100101100",
		expectedEvent: []YearEvent{{
			day:   10,
			month: 3,
			name:  "Giỗ Tổ Hùng Vương",
			kind:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			longitude: 15,
			name:      "Thanh minh",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Tý",
				from: 23,
				to:   1,
			},
			{
				chi:  "Sửu",
				from: 1,
				to:   3,
			},
			{
				chi:  "Mão",
				from: 5,
				to:   7,
			},
			{
				chi:  "Ngọ",
				from: 11,
				to:   13,
			},
			{
				chi:  "Thân",
				from: 15,
				to:   17,
			},
			{
				chi:  "Dậu",
				from: 17,
				to:   19,
			},
		},
	},
	{
		name:                   "Test Case 9",
		date:                   time.Date(2024, time.May, 22, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Bính Tuất",
		expectedMonthAlias:     "Kỷ Tỵ",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        3,
		expectedWeekdayDisplay: "Thứ Tư",
		expectedLuckyHour:      "001011001101",
		expectedEvent: []YearEvent{{
			day:   15,
			month: 4,
			name:  "Phật Đản",
			kind:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			longitude: 60,
			name:      "Tiểu mãn",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Dần",
				from: 3,
				to:   5,
			},
			{
				chi:  "Thìn",
				from: 7,
				to:   9,
			},
			{
				chi:  "Tỵ",
				from: 9,
				to:   11,
			},
			{
				chi:  "Thân",
				from: 15,
				to:   17,
			},
			{
				chi:  "Dậu",
				from: 17,
				to:   19,
			},
			{
				chi:  "Hợi",
				from: 21,
				to:   23,
			},
		},
	},
	{
		name:                   "Test Case 10",
		date:                   time.Date(2024, time.June, 10, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Ất Tỵ",
		expectedMonthAlias:     "Canh Ngọ",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        1,
		expectedWeekdayDisplay: "Thứ Hai",
		expectedLuckyHour:      "010010110011",
		expectedEvent: []YearEvent{{
			day:   5,
			month: 5,
			name:  "Lễ Đoan Ngọ",
			kind:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			longitude: 75,
			name:      "Mang chủng",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Sửu",
				from: 1,
				to:   3,
			},
			{
				chi:  "Thìn",
				from: 7,
				to:   9,
			},
			{
				chi:  "Ngọ",
				from: 11,
				to:   13,
			},
			{
				chi:  "Mùi",
				from: 13,
				to:   15,
			},
			{
				chi:  "Tuất",
				from: 19,
				to:   21,
			},
			{
				chi:  "Hợi",
				from: 21,
				to:   23,
			},
		},
	},
	{
		name:                   "Test Case 11",
		date:                   time.Date(2024, time.August, 18, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Giáp Dần",
		expectedMonthAlias:     "Nhâm Thân",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        0,
		expectedWeekdayDisplay: "Chủ Nhật",
		expectedLuckyHour:      "110011010010",
		expectedEvent: []YearEvent{{
			day:   15,
			month: 7,
			name:  "Vu Lan",
			kind:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			longitude: 135,
			name:      "Lập thu",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Tý",
				from: 23,
				to:   1,
			},
			{
				chi:  "Sửu",
				from: 1,
				to:   3,
			},
			{
				chi:  "Thìn",
				from: 7,
				to:   9,
			},
			{
				chi:  "Tỵ",
				from: 9,
				to:   11,
			},
			{
				chi:  "Mùi",
				from: 13,
				to:   15,
			},
			{
				chi:  "Tuất",
				from: 19,
				to:   21,
			},
		},
	},
	{
		name:                   "Test Case 12",
		date:                   time.Date(2024, time.September, 17, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Giáp Thân",
		expectedMonthAlias:     "Quý Dậu",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        2,
		expectedWeekdayDisplay: "Thứ Ba",
		expectedLuckyHour:      "110011010010",
		expectedEvent: []YearEvent{{
			day:   15,
			month: 8,
			name:  "Tết Trung Thu",
			kind:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			longitude: 165,
			name:      "Bạch lộ",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Tý",
				from: 23,
				to:   1,
			},
			{
				chi:  "Sửu",
				from: 1,
				to:   3,
			},
			{
				chi:  "Thìn",
				from: 7,
				to:   9,
			},
			{
				chi:  "Tỵ",
				from: 9,
				to:   11,
			},
			{
				chi:  "Mùi",
				from: 13,
				to:   15,
			},
			{
				chi:  "Tuất",
				from: 19,
				to:   21,
			},
		},
	},
	{
		name:                   "Test Case 13",
		date:                   time.Date(2025, time.January, 22, 0, 0, 0, 0, VietnamLocation()),
		expectedDayAlias:       "Tân Mão",
		expectedMonthAlias:     "Đinh Sửu",
		expectedYearAlias:      "Giáp Thìn",
		expectedWeekday:        3,
		expectedWeekdayDisplay: "Thứ Tư",
		expectedLuckyHour:      "101100110100",
		expectedEvent: []YearEvent{{
			day:   23,
			month: 12,
			name:  "Ông Táo chầu trời",
			kind:  LUNAR,
		}},
		expectedSolarTerm: SolarTerm{
			longitude: 300,
			name:      "Đại hàn",
		},
		expectedLuckyHourDetail: []LuckyHourDetail{
			{
				chi:  "Tý",
				from: 23,
				to:   1,
			},
			{
				chi:  "Dần",
				from: 3,
				to:   5,
			},
			{
				chi:  "Mão",
				from: 5,
				to:   7,
			},
			{
				chi:  "Ngọ",
				from: 11,
				to:   13,
			},
			{
				chi:  "Mùi",
				from: 13,
				to:   15,
			},
			{
				chi:  "Dậu",
				from: 17,
				to:   19,
			},
		},
	},
}

func TestToString(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			result := l.String()
			expect := fmt.Sprintf("%s-%s-%s", tc.expectedDayAlias, tc.expectedMonthAlias, tc.expectedYearAlias)
			if result != expect {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, expect, result)
			}
		})
	}
}

func TestDayAlias(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
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
			l := NewLuna(tc.date)
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
			l := NewLuna(tc.date)
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
			l := NewLuna(tc.date)
			s := l.ToSolar()
			result := s.GetDay()

			if result != tc.date.Day() {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %d", tc.name, tc.expectedYearAlias, result)
			}
		})
	}
}

func TestWeekday(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			// this is wrong for test case
			result := l.WeekDay()

			if result != tc.expectedWeekday {
				t.Errorf("Unexpected result for %s. Expected: %d, got: %d", tc.name, tc.expectedWeekday, result)
			}
		})
	}
}
func TestWeekdayDisplay(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			// this is wrong for test case
			result := l.WeekdayDisplay()

			if result != tc.expectedWeekdayDisplay {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedWeekdayDisplay, result)
			}
		})
	}
}

func TestNext(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			result := l.Next()

			if result != nil {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.name, nil, result)
			}
		})
	}
}

func TestLuckyHour(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			result := l.LuckyHour()

			if result != tc.expectedLuckyHour {
				t.Errorf("Unexpected result for %s. Expected: %s, got: %s", tc.name, tc.expectedLuckyHour, result)
			}
		})
	}
}

func TestLuckyHourDetail(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			result := l.LuckyHourDetail()
			if !reflect.DeepEqual(result, tc.expectedLuckyHourDetail) {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.name, tc.expectedLuckyHourDetail, result)
			}

		})
	}
}

func TestEvents(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			result := l.Events()
			if !reflect.DeepEqual(result, tc.expectedEvent) {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.name, tc.expectedEvent, result)
			}

		})
	}
}

func TestSolarTerms(t *testing.T) {
	for _, tc := range lunarTestCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLuna(tc.date)
			result := l.SolarTerms()
			if !reflect.DeepEqual(result, tc.expectedSolarTerm) {
				t.Errorf("Unexpected result for %s. Expected: %v, got: %v", tc.name, tc.expectedSolarTerm, result)
			}

		})
	}
}
