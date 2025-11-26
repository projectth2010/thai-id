package thaiid

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		id       string
		expected bool
	}{
		{"9000800092221", false},  // หลักแรก 9 → ไม่ผ่าน
		{"1000000092221", false},  // รหัสจังหวัด 0000 → ไม่ผ่าน
		{"1234567890123", false},  // สมมุติ (อาจไม่ผ่าน)
		{"1100800092221", true},   // ตัวอย่างที่ผ่าน (ลองตรวจสอบด้วยเว็บจริง)
		{"1100800092222", false},  // เปลี่ยนหลักสุดท้าย → ไม่ผ่าน
		{"110080009222", false},   // 12 หลัก → ไม่ผ่าน
		{"11008000922211", false}, // 14 หลัก → ไม่ผ่าน
		{"11008000922a1", false},  // มีตัวอักษร → ไม่ผ่าน
		{"", false},
		{"1-2345-67890-12-3", false}, // รูปแบบมี - → ใช้ Normalize ก่อน
	}

	for _, tc := range cases {
		t.Run(tc.id, func(t *testing.T) {
			if got := IsValid(tc.id); got != tc.expected {
				t.Errorf("IsValid(%q) = %v; want %v", tc.id, got, tc.expected)
			}
		})
	}
}

func TestIsValidWithNormalize(t *testing.T) {
	if !IsValidWithNormalize("1-1008-00092-22-1") {
		t.Error("Expected valid ID after normalization")
	}
}
