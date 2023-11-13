package vcalendar

import (
	"testing"
	"time"
)

func TestB2I(t *testing.T) {
	tc := []struct {
		args     bool
		expected int
	}{
		{
			args:     true,
			expected: 1,
		},
		{
			args:     false,
			expected: 0,
		},
	}

	for _, v := range tc {
		if b2i(v.args) != v.expected {
			t.Errorf("Failed: %v.\nExpected: %d.\nGot: %d", v.args, v.expected, b2i(v.args))
		}
	}
}

func TestLoc2tz(t *testing.T) {
	tc := []struct {
		args     string
		expected int
	}{
		{
			args:     "Asia/Ho_Chi_Minh",
			expected: 7,
		},
		{
			args:     "UTC",
			expected: 0,
		},
		{
			args:     "America/New_York",
			expected: -5,
		},
	}

	for _, v := range tc {
		location, err := time.LoadLocation(v.args)
		if err != nil {
			panic(err)
		}

		if loc2tz(*location) != v.expected {
			t.Errorf("Failed: %v.\nExpected: %d.\nGot: %d", v.args, v.expected, loc2tz(*location))
		}
	}
}

func TestVietnamLocation(t *testing.T) {
	tc := []struct {
		expected string
	}{
		{
			expected: "Asia/Ho_Chi_Minh",
		},
	}

	for _, v := range tc {
		location := VietnamLocation()
		l, err := time.LoadLocation(v.expected)
		if err != nil {
			panic(err)
		}
		if location.String() != l.String() {
			t.Errorf("Failed: %v.\nExpected: %v.\nGot: %v", v.expected, l, location)
		}
	}
}
