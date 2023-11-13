package vcalendar

import "time"

func New(t time.Time) Lunar {
	// Compute offset of timezone
	tzoff := getTz(t)

	// Compute lunar date
	d, m, y, leap := Solar2Lunar(t.Day(), int(t.Month()), t.Year(), tzoff)
	lunar := Lunar{
		c: Calendar{
			t:     t,
			Day:   d,
			Month: m,
			Year:  y,
		},
		Leap: leap == 1,
	}

	return lunar
}

func NewWithLunar(d, m, y int, l bool, loc time.Location) Lunar {
	lunar := Lunar{
		c: Calendar{
			Day:   d,
			Month: m,
			Year:  y,
		},
		Leap: l,
	}
	tzoff := loc2tz(loc)

	sd, sm, sy := Lunar2Solar(d, m, y, b2i(l), tzoff)
	t := time.Date(sy, time.Month(sm), sd, 0, 0, 0, 0, &loc)
	lunar.c.t = t

	return lunar
}
