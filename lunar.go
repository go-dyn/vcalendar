package vcalendar

import (
	"fmt"
)

func (l *Lunar) String() string {
	return fmt.Sprintf("ngày %s, tháng %s, năm %s", l.DayAlias(), l.MonthAlias(), l.YearAlias())
}

func (l *Lunar) DayAlias() string {
	return fmt.Sprintf("%s %s", l.dayCan(), l.dayChi())
}

func (l *Lunar) dayCan() string {
	jd := date2JuliusDay(l.c.t.Day(), int(l.c.t.Month()), l.c.t.Year())
	return CAN[(jd+9)%10]
}

func (l *Lunar) dayChi() string {
	jd := date2JuliusDay(l.c.t.Day(), int(l.c.t.Month()), l.c.t.Year())
	return CHI[(jd+1)%12]
}

func (l *Lunar) MonthAlias() string {
	if l.Leap {
		return fmt.Sprintf("%s %s Nhuận", l.monthCan(), l.monthChi())
	}
	return fmt.Sprintf("%s %s", l.monthCan(), l.monthChi())
}

func (l *Lunar) monthCan() string {
	i := (l.c.Year*12 + l.c.Month + 3) % 10
	return CAN[i]
}

func (l *Lunar) monthChi() string {
	i := (l.c.Month + 1) % 12
	return CHI[i]
}

func (l *Lunar) YearAlias() string {
	return fmt.Sprintf("%s %s", l.yearCan(), l.yearChi())
}

func (l *Lunar) yearCan() string {
	i := (l.c.Year + 6) % 10
	return CAN[i]
}

func (l *Lunar) yearChi() string {
	i := (l.c.Year + 8) % 12
	return CHI[i]
}

func (l *Lunar) LuckyHour() string {
	jd := date2JuliusDay(l.c.t.Day(), int(l.c.t.Month()), l.c.t.Year())
	chiOfDay := (jd + 1) % 12
	luckyHour := LUCKY_HOURS[chiOfDay%6]

	return luckyHour
}

func (l *Lunar) LuckyHourDetail() []LuckyHourDetail {
	ret := []LuckyHourDetail{}
	luckyHour := l.LuckyHour()

	for i := 0; i < 12; i++ {
		index := luckyHour[i]
		if index == '1' {
			detail := LuckyHourDetail{
				Chi:  CHI[i],
				From: (i*2 + 23) % 24,
				To:   (i*2 + 1) % 24,
			}
			ret = append(ret, detail)
		}
	}
	return ret
}

func (l *Lunar) SolarTerms() SolarTerm {
	//_, offset := l.t.Zone()
	jd := date2JuliusDay(l.c.t.Day(), int(l.c.t.Month()), l.c.t.Year())
	solarTerm := getSolarTerm(jd+1, 7.0)
	return SOLAR_TERMS[solarTerm]
}

func (l *Lunar) Events() []YearEvent {
	events := []YearEvent{}
	dd := l.c.Day
	mm := l.c.Month
	for i := 0; i < len(YEARLY_EVENTS); i++ {
		event := YEARLY_EVENTS[i]
		if event.day == dd && event.month == mm {
			events = append(events, event)
		}
	}

	return events
}

func (l *Lunar) ToSolar() Solar {
	d, m, y := Lunar2Solar(l.c.Day, l.c.Month, l.c.Year, b2i(l.Leap), getTz(l.c.t))
	return Solar{
		c: Calendar{
			Day:   d,
			Month: m,
			Year:  y,
		},
	}
}

func (l *Lunar) Weekday() int {
	jd := date2JuliusDay(l.c.t.Day(), int(l.c.t.Month()), l.c.t.Year())
	return (jd + 1) % 7
}

func (l *Lunar) WeekdayDisplay() string {
	day := l.Weekday()
	return WEEKS[day]
}

// TODO: implement function
func (l *Lunar) Next() *Lunar {
	return nil
}
