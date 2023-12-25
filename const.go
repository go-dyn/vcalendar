package vcalendar

import "math"

var (
	UNKNOWN   = 0
	SOLAR     = 1
	LUNAR     = 2
	PI        = math.Pi
	VI        = "vi"
	EN        = "en"
	CAN       = [10]string{"Giáp", "Ất", "Bính", "Đinh", "Mậu", "Kỷ", "Canh", "Tân", "Nhâm", "Quý"}
	CHI       = [12]string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}
	WEEKS     = [7]string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"}
	MONTHS    = [12]string{"Giêng", "Hai", "Ba", "Tư", "Năm", "Sáu", "Bảy", "Tám", "Chín", "Mười", "Một", "Chạp"}
	LANGUAGES = []Locale{{
		code: VI,
		name: "Việt Nam",
	}, {
		code: EN,
		name: "English",
	}}
	SOLAR_TERMS = []SolarTerm{
		{
			longitude: 0,
			name:      "Xuân phân",
		},
		{
			longitude: 15,
			name:      "Thanh minh",
		},
		{
			longitude: 30,
			name:      "Cốc vũ",
		},
		{
			longitude: 45,
			name:      "Lập hạ",
		},
		{
			longitude: 60,
			name:      "Tiểu mãn",
		},
		{
			longitude: 75,
			name:      "Mang chủng",
		},
		{
			longitude: 90,
			name:      "Hạ chí",
		},
		{
			longitude: 105,
			name:      "Tiểu thử",
		},
		{
			longitude: 120,
			name:      "Đại thử",
		},
		{
			longitude: 135,
			name:      "Lập thu",
		},
		{
			longitude: 150,
			name:      "Xử thử",
		},
		{
			longitude: 165,
			name:      "Bạch lộ",
		},
		{
			longitude: 180,
			name:      "Thu phân",
		},
		{
			longitude: 195,
			name:      "Hàn lộ",
		},
		{
			longitude: 210,
			name:      "Sương giáng",
		},
		{
			longitude: 225,
			name:      "Lập đông",
		},
		{
			longitude: 240,
			name:      "Tiểu tuyết",
		},
		{
			longitude: 255,
			name:      "Đại tuyết",
		},
		{
			longitude: 270,
			name:      "Đông chí",
		},
		{
			longitude: 285,
			name:      "Tiểu hàn",
		},
		{
			longitude: 300,
			name:      "Đại hàn",
		},
		{
			longitude: 315,
			name:      "Lập xuân",
		},
		{
			longitude: 330,
			name:      "Vũ Thủy",
		},
		{
			longitude: 345,
			name:      "Kinh trập",
		},
	}
	LUCKY_HOURS       = [6]string{"110100101100", "001101001011", "110011010010", "101100110100", "001011001101", "010010110011"}
	VIETNAM_TIME_ZONE = "Asia/Ho_Chi_Minh"
	YEARLY_EVENTS     = []YearEvent{
		{
			day:   1,
			month: 1,
			name:  "Tết Nguyên Đán",
			kind:  LUNAR,
		},
		{
			day:   15,
			month: 1,
			name:  "Rằm tháng Giêng",
			kind:  LUNAR,
		},
		{
			day:   10,
			month: 3,
			name:  "Giỗ Tổ Hùng Vương",
			kind:  LUNAR,
		},
		{
			day:   15,
			month: 4,
			name:  "Phật Đản",
			kind:  LUNAR,
		},
		{
			day:   5,
			month: 5,
			name:  "Lễ Đoan Ngọ",
			kind:  LUNAR,
		},
		{
			day:   15,
			month: 7,
			name:  "Vu Lan",
			kind:  LUNAR,
		},
		{
			day:   15,
			month: 8,
			name:  "Tết Trung Thu",
			kind:  LUNAR,
		},
		{
			day:   23,
			month: 12,
			name:  "Ông Táo chầu trời",
			kind:  LUNAR,
		},
	}
)
