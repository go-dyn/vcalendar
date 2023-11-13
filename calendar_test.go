package vcalendar

import (
	"testing"
	"time"
)

var CalendarTestCases = []struct {
	Name     string
	Day      int
	Month    int
	Year     int
	Type     int
	Language Locale
	Time     time.Time
}{
	{
		Name:     "Test case 1",
		Day:      1,
		Month:    1,
		Year:     2023,
		Language: Locale{Code: VI, Name: "Việt Nam"},
		Time:     time.Now(),
		Type:     SOLAR,
	},
}

func TestCalendar(t *testing.T) {
	for _, v := range CalendarTestCases {
		c := Solar{
			Date: Calendar{
				Day:      v.Day,
				Month:    v.Month,
				Year:     v.Year,
				Language: v.Language,
				Time:     v.Time,
				Type:     v.Type,
			},
		}

		if v.Day != c.Date.GetDay() ||
			v.Month != c.Date.GetMonth() ||
			v.Year != c.Date.GetYear() ||
			v.Language != c.Date.GetLanguage() ||
			v.Type != c.Date.GetType() {
			t.Errorf("Failed test case: %v\nExpect: %v %v %v %v %v\nGot: %v %v %v %v %v", v.Name, v.Day, v.Month, v.Year, v.Language, v.Type, c.Date.GetDay(), c.Date.GetMonth(), c.Date.GetYear(), c.Date.GetLanguage(), c.Date.GetType())
		}

		year, month, day := time.Now().Date()
		language := Locale{Code: EN, Name: "English"}
		c.Date.SetDay(day).SetMonth(int(month)).SetYear(year).SetLanguage(language)

		if day != c.Date.GetDay() ||
			int(month) != c.Date.GetMonth() ||
			year != c.Date.GetYear() ||
			language != c.Date.GetLanguage() {
			t.Errorf("Failed test case: %v\nExpect: %v %v %v %v\nGot: %v %v %v %v", v.Name, v.Day, v.Month, v.Year, v.Language, c.Date.GetDay(), c.Date.GetMonth(), c.Date.GetYear(), c.Date.GetLanguage())
		}
	}
}
