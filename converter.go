package vcalendar

import (
	"math"
	"time"
)

// BooleanToInt converts a boolean value to an integer value.
func BooleanToInt(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue #6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func IntToFloat(i int) float64 {
	return float64(i)
}

func MathFloor(f float64) float64 {
	return math.Floor(f)
}

func Floor(f float64) int {
	return int(MathFloor(f))
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

// GetTimeZone computes offset in hours east of UTC
// e.g. UTC+7 returns 7, UTC-7 returns -7
func GetTimeZone(t time.Time) int {
	_, offset := t.Zone()
	tz := offset / 3600

	return tz
}

// LocationToTimeZone computes offset in hours east of UTC from time.Location
// e.g. UTC+7 returns 7, UTC-7 returns -7
func LocationToTimeZone(l time.Location) int {
	t := time.Now().In(&l)
	return GetTimeZone(t)
}

// VietnamLocation returns the *time.Location for Vietnam
func VietnamLocation() *time.Location {
	loc, _ := time.LoadLocation(VIETNAM_TIME_ZONE)
	return loc
}

func DateToJuliusDay(dd, mm, yy int) int {
	a := Floor((14 - IntToFloat(mm)) / 12)
	y := IntToFloat(yy) + 4800 - IntToFloat(a)
	m := IntToFloat(mm) + 12*IntToFloat(a) - 3
	jd := IntToFloat(dd) + MathFloor((153*m+2)/5) + 365*y + MathFloor(y/4) - MathFloor(y/100) + MathFloor(y/400) - 32045
	if jd < 2299161 {
		jd = IntToFloat(dd) + MathFloor((153*m+2)/5) + 365*y + MathFloor(y/4) - 32083
	}

	return Floor(jd)
}

func JuliusDayToDate(jd int) (day, month, year int) {
	var a, b, c, d, e, m int
	if jd > 2299160 {
		a = jd + 32044
		b = (4*a + 3) / 146097
		c = a - (b*146097)/4
	} else {
		b = 0
		c = jd + 32082
	}

	d = (4*c + 3) / 1461
	e = c - (1461*d)/4
	m = (5*e + 2) / 153

	day = e - (153*m+2)/5 + 1
	month = m + 3 - 12*(m/10)
	year = b*100 + d - 4800 + (m / 10)

	return
}

func NewMoon(k int) float64 {
	var kf, t, t2, t3, dr, jd1, m, mpr, f, c1, deltat, jdNew float64
	kf = IntToFloat(k)
	t = kf / 1236.85 // Time in Julian centuries from 1900 January 0.5
	t2 = t * t
	t3 = t2 * t
	dr = PI / 180
	jd1 = 2415020.75933 + 29.53058868*kf + 0.0001178*t2 - 0.000000155*t3
	jd1 = jd1 + 0.00033*Sin((166.56+132.87*t-0.009173*t2)*dr)       // Mean new moon
	m = 359.2242 + 29.10535608*kf - 0.0000333*t2 - 0.00000347*t3    // Sun's mean anomaly
	mpr = 306.0253 + 385.81691806*kf + 0.0107306*t2 + 0.00001236*t3 // Moon's mean anomaly
	f = 21.2964 + 390.67050646*kf - 0.0016528*t2 - 0.00000239*t3    // Moon's argument of latitude
	c1 = (0.1734-0.000393*t)*Sin(m*dr) + 0.0021*Sin(2*dr*m)
	c1 = c1 - 0.4068*Sin(mpr*dr) + 0.0161*Sin(dr*2*mpr)
	c1 = c1 - 0.0004*Sin(dr*3*mpr)
	c1 = c1 + 0.0104*Sin(dr*2*f) - 0.0051*Sin(dr*(m+mpr))
	c1 = c1 - 0.0074*Sin(dr*(m-mpr)) + 0.0004*Sin(dr*(2*f+m))
	c1 = c1 - 0.0004*Sin(dr*(2*f-m)) - 0.0006*Sin(dr*(2*f+mpr))
	c1 = c1 + 0.0010*Sin(dr*(2*f-mpr)) + 0.0005*Sin(dr*(2*mpr+m))
	if t < -11 {
		deltat = 0.001 + 0.000839*t + 0.0002261*t2 - 0.00000845*t3 - 0.000000081*t*t3
	} else {
		deltat = -0.000278 + 0.000265*t + 0.000262*t2
	}
	jdNew = jd1 + c1 - deltat

	return jdNew
}

func GetNewMoonDay(k, tz int) int {
	return Floor(NewMoon(k) + 0.5 + IntToFloat(tz)/24)
}

func SunLongitude(jdn float64) float64 {
	var t, t2, dr, m, l0, dl, l float64
	t = (jdn - 2451545.0) / 36525 // Time in Julian centuries from 2000-01-01 12:00:00 GMT
	t2 = t * t
	dr = PI / 180                                                  // degree to radian
	m = 357.52910 + 35999.05030*t - 0.0001559*t2 - 0.00000048*t*t2 // mean anomaly, degree
	l0 = 280.46645 + 36000.76983*t + 0.0003032*t2                  // mean longitude, degree
	dl = (1.914600 - 0.004817*t - 0.000014*t2) * Sin(dr*m)
	dl = dl + (0.019993-0.000101*t)*Sin(dr*2*m) + 0.00029*Sin(dr*3*m)
	l = l0 + dl // true longitude, degree
	l = l * dr
	l = l - PI*2*(MathFloor(l/(PI*2))) // Normalize to (0, 2*PI)
	return l
}

func GetSunLongitude(d, tz int) int {
	return Floor((SunLongitude(IntToFloat(d)-0.5-IntToFloat(tz)/24) / PI) * 6)
}

func GetSolarTerm(dayNumber int, timeZone int) int {
	return Floor(SunLongitude(IntToFloat(dayNumber)-0.5-IntToFloat(timeZone)/24.0) / PI * 12)
}

/* Find the day that starts the lunar month 11 of the given year for the given time zone */
func GetLunarMonthStartDate(yy, tz int) int {
	var off, k, nm int

	off = DateToJuliusDay(31, 12, yy) - 2415021
	k = Floor(IntToFloat(off) / 29.530588853)
	nm = GetNewMoonDay(k, tz)
	sunLong := GetSunLongitude(nm, tz) // sun longitude at local midnight
	if sunLong >= 9 {
		nm = GetNewMoonDay(k-1, tz)
	}
	return nm
}

func GetLeapMonthOffset(a11 float64, tz int) int {
	k := Floor((a11-2415021.076998695)/29.530588853 + 0.5)
	last := 0
	i := 1 // We start with the month following lunar month 11
	arc := GetSunLongitude(GetNewMoonDay(k+i, tz), tz)
	for {
		last = arc
		i++
		newmoon := GetNewMoonDay(k+i, tz)
		arc = GetSunLongitude(newmoon, tz)

		if arc == last || i >= 14 {
			break
		}
	}
	return i - 1
}

// The function SolarToLunar converts a given Solar date (dd, mm, yy) into
// the corresponding Lunar date (lunarD, lunarM, lunarY) in the Vietnamese Lunar calendar.
// It also determines if the given Lunar month is a leap month or not (lunarLeap).
// The time zone (tz) is used to calculate the date in the Lunar calendar.
func SolarToLunar(dd, mm, yy, tz int) (lunarD, lunarM, lunarY, lunarLeap int) {
	dayNumber := DateToJuliusDay(dd, mm, yy)
	k := Floor((IntToFloat(dayNumber) - 2415021.076998695) / 29.530588853)
	monthStart := GetNewMoonDay(k+1, tz)
	if monthStart > dayNumber {
		monthStart = GetNewMoonDay(k, tz)
	}
	a11 := GetLunarMonthStartDate(yy, tz)
	b11 := a11
	if a11 >= monthStart {
		lunarY = yy
		a11 = GetLunarMonthStartDate(yy-1, tz)
	} else {
		lunarY = yy + 1
		b11 = GetLunarMonthStartDate(yy+1, tz)
	}
	lunarD = dayNumber - monthStart + 1
	diff := Floor((IntToFloat(monthStart) - IntToFloat(a11)) / 29)

	lunarLeap = 0
	lunarM = diff + 11
	if b11-a11 > 365 {
		leapMonthDiff := GetLeapMonthOffset(IntToFloat(a11), tz)
		if diff >= leapMonthDiff {
			lunarM = diff + 10
			if diff == leapMonthDiff {
				lunarLeap = 1
			}
		}
	}
	if lunarM > 12 {
		lunarM = lunarM - 12
	}
	if lunarM >= 11 && diff < 4 {
		lunarY -= 1
	}

	return
}

// The function LunarToSolar converts a given Lunar date (lunarDay, lunarMonth, lunarYear, lunarLeap)
// in the Vietnamese Lunar calendar into the corresponding Solar date (d, m, y).
// The time zone (tz) is used to calculate the date in the Solar calendar.
func LunarToSolar(lunarDay, lunarMonth, lunarYear, lunarLeap, tz int) (d, m, y int) {
	var k, a11, b11, off, leapOff, leapMonth, monthStart int
	if lunarMonth < 11 {
		a11 = GetLunarMonthStartDate(lunarYear-1, tz)
		b11 = GetLunarMonthStartDate(lunarYear, tz)
	} else {
		a11 = GetLunarMonthStartDate(lunarYear, tz)
		b11 = GetLunarMonthStartDate(lunarYear+1, tz)
	}
	k = Floor(0.5 + (IntToFloat(a11)-2415021.076998695)/29.530588853)
	off = lunarMonth - 11
	if off < 0 {
		off += 12
	}
	if b11-a11 > 365 {
		leapOff = GetLeapMonthOffset(IntToFloat(a11), tz)
		leapMonth = leapOff - 2
		if leapMonth < 0 {
			leapMonth += 12
		}
		if lunarLeap != 0 && lunarMonth != leapMonth {
			return
		} else if lunarLeap != 0 || off >= leapOff {
			off += 1
		}
	}
	monthStart = GetNewMoonDay(k+off, tz)

	return JuliusDayToDate(monthStart + lunarDay - 1)
}
