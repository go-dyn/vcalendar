package vcalendar

var (
	UNKNOWN   = 0
	SOLAR     = 1
	LUNAR     = 2
	PI        = 3.141592653589793
	VI        = "vi"
	EN        = "en"
	CAN       = [10]string{"Giáp", "Ất", "Bính", "Đinh", "Mậu", "Kỷ", "Canh", "Tân", "Nhâm", "Quý"}
	CHI       = [12]string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}
	WEEKS     = [7]string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"}
	MONTHS    = [12]string{"Giêng", "Hai", "Ba", "Tư", "Năm", "Sáu", "Bảy", "Tám", "Chín", "Mười", "Một", "Chạp"}
	LANGUAGES = []Locale{{
		Code: VI,
		Name: "Việt Nam",
	}, {
		Code: EN,
		Name: "English",
	}}
	SOLAR_TERMS = []SolarTerm{
		{
			Longitude: 0,
			Name:      "Xuân phân",
		},
		{
			Longitude: 15,
			Name:      "Thanh minh",
		},
		{
			Longitude: 30,
			Name:      "Cốc vũ",
		},
		{
			Longitude: 45,
			Name:      "Lập hạ",
		},
		{
			Longitude: 60,
			Name:      "Tiểu mãn",
		},
		{
			Longitude: 75,
			Name:      "Mang chủng",
		},
		{
			Longitude: 90,
			Name:      "Hạ chí",
		},
		{
			Longitude: 105,
			Name:      "Tiểu thử",
		},
		{
			Longitude: 120,
			Name:      "Đại thử",
		},
		{
			Longitude: 135,
			Name:      "Lập thu",
		},
		{
			Longitude: 150,
			Name:      "Xử thử",
		},
		{
			Longitude: 165,
			Name:      "Bạch lộ",
		},
		{
			Longitude: 180,
			Name:      "Thu phân",
		},
		{
			Longitude: 195,
			Name:      "Hàn lộ",
		},
		{
			Longitude: 210,
			Name:      "Sương giáng",
		},
		{
			Longitude: 225,
			Name:      "Lập đông",
		},
		{
			Longitude: 240,
			Name:      "Tiểu tuyết",
		},
		{
			Longitude: 255,
			Name:      "Đại tuyết",
		},
		{
			Longitude: 270,
			Name:      "Đông chí",
		},
		{
			Longitude: 285,
			Name:      "Tiểu hàn",
		},
		{
			Longitude: 300,
			Name:      "Đại hàn",
		},
		{
			Longitude: 315,
			Name:      "Lập xuân",
		},
		{
			Longitude: 330,
			Name:      "Vũ Thủy",
		},
		{
			Longitude: 345,
			Name:      "Kinh trập",
		},
	}
	LUCKY_HOURS       = [6]string{"110100101100", "001101001011", "110011010010", "101100110100", "001011001101", "010010110011"}
	VIETNAM_TIME_ZONE = "Asia/Ho_Chi_Minh"
	YEARLY_EVENTS     = []YearEvent{
		{
			Day:   1,
			Month: 1,
			Name:  "Tết Nguyên Đán",
			Type:  LUNAR,
		},
		{
			Day:   15,
			Month: 1,
			Name:  "Rằm tháng Giêng",
			Type:  LUNAR,
		},
		{
			Day:   10,
			Month: 3,
			Name:  "Giỗ Tổ Hùng Vương",
			Type:  LUNAR,
		},
		{
			Day:   15,
			Month: 4,
			Name:  "Phật Đản",
			Type:  LUNAR,
		},
		{
			Day:   5,
			Month: 5,
			Name:  "Lễ Đoan Ngọ",
			Type:  LUNAR,
		},
		{
			Day:   15,
			Month: 7,
			Name:  "Vu Lan",
			Type:  LUNAR,
		},
		{
			Day:   15,
			Month: 8,
			Name:  "Tết Trung Thu",
			Type:  LUNAR,
		},
		{
			Day:   23,
			Month: 12,
			Name:  "Ông Táo chầu trời",
			Type:  LUNAR,
		},
	}
)
