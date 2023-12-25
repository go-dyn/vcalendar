package vcalendar

import (
	"reflect"
	"testing"
	"time"
)

func TestCalendar_String(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("Calendar.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_Format(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.Format(tt.args.format); got != tt.want {
				t.Errorf("Calendar.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetTime(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.GetTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.GetTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetDay(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.GetDay(); got != tt.want {
				t.Errorf("Calendar.GetDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetMonth(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Month
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.GetMonth(); got != tt.want {
				t.Errorf("Calendar.GetMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetIntMonth(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.GetIntMonth(); got != tt.want {
				t.Errorf("Calendar.GetIntMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetYear(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.GetYear(); got != tt.want {
				t.Errorf("Calendar.GetYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetWeekDay(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.GetWeekDay(); got != tt.want {
				t.Errorf("Calendar.GetWeekDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetLocation(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *time.Location
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.GetLocation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.GetLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetTimeZone(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name       string
		fields     fields
		wantName   string
		wantOffset int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			gotName, gotOffset := c.GetTimeZone()
			if gotName != tt.wantName {
				t.Errorf("Calendar.GetTimeZone() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotOffset != tt.wantOffset {
				t.Errorf("Calendar.GetTimeZone() gotOffset = %v, want %v", gotOffset, tt.wantOffset)
			}
		})
	}
}

func TestCalendar_SetDay(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		Day int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.SetDay(tt.args.Day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.SetDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_SetIntMonth(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		Month int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.SetIntMonth(tt.args.Month); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.SetIntMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_SetMonth(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		Month time.Month
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.SetMonth(tt.args.Month); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.SetMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_SetYear(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		Year int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.SetYear(tt.args.Year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.SetYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_SetLocation(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		location *time.Location
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.SetLocation(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.SetLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_SetTimeZone(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		name   string
		offset int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.SetTimeZone(tt.args.name, tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.SetTimeZone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_IsLeap(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.IsLeap(); got != tt.want {
				t.Errorf("Calendar.IsLeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_GetJd(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.GetJd(); got != tt.want {
				t.Errorf("Calendar.GetJd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_ToVnLunar(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.ToVnLunar(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.ToVnLunar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_FromSolar(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.FromSolar(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.FromSolar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalendar_FromVnLunar(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		time time.Time
		leap int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calendar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calendar{
				Time: tt.fields.Time,
			}
			if got := c.FromVnLunar(tt.args.time, tt.args.leap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calendar.FromVnLunar() = %v, want %v", got, tt.want)
			}
		})
	}
}
