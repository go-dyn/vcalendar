package vcalendar

import (
	"fmt"
	"time"
)

func (l *Calendar) DayAlias() string {
	return fmt.Sprintf("%s %s", l.DayCan(), l.DayChi())
}

func (l *Calendar) DayCan() string {
	jd := DateToJuliusDay(l.GetDay(), l.GetIntMonth(), l.GetYear())
	return CAN[(jd+9)%10]
}

func (l *Calendar) DayChi() string {
	jd := DateToJuliusDay(l.GetDay(), l.GetIntMonth(), l.GetYear())
	return CHI[(jd+1)%12]
}

func (l *Calendar) MonthAlias() string {
	if l.IsLeap() {
		return fmt.Sprintf("%s %s Nhuận", l.MonthCan(), l.MonthChi())
	}
	return fmt.Sprintf("%s %s", l.MonthCan(), l.MonthChi())
}

func (l *Calendar) MonthCan() string {
	i := (l.GetYear()*12 + l.GetIntMonth() + 3) % 10
	return CAN[i]
}

func (l *Calendar) MonthChi() string {
	i := (l.GetIntMonth() + 1) % 12
	return CHI[i]
}

func (l *Calendar) YearAlias() string {
	return fmt.Sprintf("%s %s", l.YearCan(), l.YearChi())
}

func (l *Calendar) YearCan() string {
	i := (l.GetYear() + 6) % 10
	return CAN[i]
}

func (l *Calendar) YearChi() string {
	i := (l.GetYear() + 8) % 12
	return CHI[i]
}

func (l *Calendar) LuckyHour() string {
	jd := DateToJuliusDay(l.GetDay(), int(l.GetIntMonth()), l.GetYear())
	chiOfDay := (jd + 1) % 12
	luckyHour := LUCKY_HOURS[chiOfDay%6]

	return luckyHour
}

func (l *Calendar) LuckyHourDetail() []LuckyHourDetail {
	ret := []LuckyHourDetail{}
	luckyHour := l.LuckyHour()

	for i := 0; i < 12; i++ {
		index := luckyHour[i]
		if index == '1' {
			detail := LuckyHourDetail{
				chi:  CHI[i],
				from: (i*2 + 23) % 24,
				to:   (i*2 + 1) % 24,
			}
			ret = append(ret, detail)
		}
	}
	return ret
}

func (l *Calendar) SolarTerms() SolarTerm {
	//_, offset := l.Zone()
	jd := DateToJuliusDay(l.GetDay(), int(l.GetIntMonth()), l.GetYear())
	solarTerm := GetSolarTerm(jd+1, 7.0)
	return SOLAR_TERMS[solarTerm]
}

func (l *Calendar) Events() []YearEvent {
	events := []YearEvent{}
	dd := l.GetDay()
	mm := l.GetIntMonth()
	for i := 0; i < len(YEARLY_EVENTS); i++ {
		event := YEARLY_EVENTS[i]
		if event.day == dd && event.month == mm && event.kind == LUNAR {
			events = append(events, event)
		}
	}

	return events
}

func (l *Calendar) ToSolar() Calendar {
	_, offset := l.Zone()
	d, m, y := LunarToSolar(l.GetDay(), l.GetIntMonth(), l.GetYear(), BooleanToInt(l.IsLeap()), offset)
	calendar := &Calendar{}
	calendar.FromVnLunar(time.Date(y, time.Month(m), d, 0, 0, 0, 0, l.GetLocation()), BooleanToInt(l.IsLeap()))
	return *calendar
}

func (l *Calendar) WeekDay() int {
	jd := DateToJuliusDay(l.GetDay(), int(l.GetIntMonth()), l.GetYear())
	return (jd + 1) % 7
}

func (l *Calendar) WeekdayDisplay() string {
	day := l.WeekDay()
	return WEEKS[day]
}

// TODO: implement function
func (l *Calendar) Next() *Calendar {
	return nil
}
