package vcalendar

import (
	"encoding/json"
	"io"
	"math"
	"os"
	"testing"
)

type TestCase struct {
	Solar []int `json:"solar"`
	Lunar []int `json:"lunar"`
}

func TestSolarToLunar(t *testing.T) {
	testCases := readTestData()
	for _, v := range testCases {
		d, m, y, leap := SolarToLunar(v.Solar[0], v.Solar[1], v.Solar[2], v.Solar[3])

		if d != v.Lunar[0] || m != v.Lunar[1] || y != v.Lunar[2] || v.Lunar[3] != leap {
			t.Errorf("Failed test case: %v\nExpect: %v\nGot: %v %v %v %v", v.Solar, v.Lunar, d, m, y, leap)
		}
	}
}

func TestLunarToSolar(t *testing.T) {
	testCases := readTestData()
	for _, v := range testCases {
		d, m, y := LunarToSolar(v.Lunar[0], v.Lunar[1], v.Lunar[2], v.Lunar[3], 7)

		if d != v.Solar[0] || m != v.Solar[1] || y != v.Solar[2] {
			t.Errorf("Failed test case: %v\nExpect: %v\nGot: %v %v %v", v.Lunar, v.Solar, d, m, y)
		}
	}
}

func TestLunarToSolarTimeZone0(t *testing.T) {
	testCases := readTestData0()
	for _, v := range testCases {
		d, m, y := LunarToSolar(v.Lunar[0], v.Lunar[1], v.Lunar[2], v.Lunar[3], 0)

		if d != v.Solar[0] || m != v.Solar[1] || y != v.Solar[2] {
			t.Errorf("Failed test case: %v\nExpect: %v\nGot: %v %v %v", v.Lunar, v.Solar, d, m, y)
		}
	}
}

func readTestData() []TestCase {
	var testCases []TestCase
	jsonFile, err := os.Open("./testdata/testcase.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &testCases)
	if err != nil {
		panic(err)
	}
	return testCases
}

func readTestData0() []TestCase {
	var testCases []TestCase
	jsonFile, err := os.Open("./testdata/testcase0.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &testCases)
	if err != nil {
		panic(err)
	}
	return testCases
}

func TestDate2JuliusDay(t *testing.T) {
	tc := []struct {
		name     string
		expected []int
		jd       int
	}{
		{
			name:     "Test case 1",
			expected: []int{4, 10, 1582},
			jd:       2299160,
		},
		{
			name:     "Test case 2",
			expected: []int{15, 10, 1582},
			jd:       2299161,
		},
		{
			name:     "Test case 3",
			expected: []int{16, 10, 1582},
			jd:       2299162,
		},
		{
			name:     "Test case 4",
			expected: []int{26, 9, 1582},
			jd:       2299152,
		},
	}

	for _, v := range tc {
		day, month, year := JuliusDayToDate(v.jd)
		if day != v.expected[0] || month != v.expected[1] || year != v.expected[2] {
			t.Errorf("Failed: %v.\nExpected: %d.\nGot: %d", v.name, v.expected, struct {
				day   int
				month int
				year  int
			}{day, month, year})
		}
	}
}

func TestJuliusDay2Date(t *testing.T) {
	tc := []struct {
		day      []int
		expected int
	}{
		{
			day:      []int{4, 10, 1582},
			expected: 2299160,
		},
		{
			day:      []int{10, 10, 1999},
			expected: 2451462,
		},
		{
			day:      []int{10, 10, 1965},
			expected: 2439044,
		},
		{
			day:      []int{10, 10, 1900},
			expected: 2415303,
		},
	}

	for _, v := range tc {
		jd := DateToJuliusDay(v.day[0], v.day[1], v.day[2])
		if jd != v.expected {
			t.Errorf("Failed: %v.\nExpected: %d.\nGot: %d", v.day, v.expected, jd)
		}
	}
}

func TestNewMoon(t *testing.T) {
	tc := []struct {
		k        int
		expected float64
	}{
		{
			k:        -14842,
			expected: 1976727.755749,
		},
		{
			k:        0,
			expected: 2415021.076999,
		},
		{
			k:        1,
			expected: 2415050.557026,
		},
		{
			k:        -1,
			expected: 2414991.532265,
		},
		{
			k:        14842,
			expected: 2853313.283231,
		},
	}

	for _, v := range tc {
		jd := NewMoon(v.k)
		if !withTolerane(jd, v.expected) {
			t.Errorf("Failed: %v.\nExpected: %f.\nGot: %f", v.k, v.expected, jd)
		}
	}
}

func withTolerane(a, b float64) bool {
	tolerance := 0.00001
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	} else {
		return false
	}
}
