package thaiid

import (
	"regexp"
	"strconv"
)

// IsValid ตรวจสอบว่าเลขบัตรประชาชนไทยถูกต้องตามรูปแบบและ checksum
func IsValid(id string) bool {
	// 1. ตรวจสอบรูปแบบ: ต้องเป็นตัวเลข 13 หลัก
	if !regexp.MustCompile(`^\d{13}$`).MatchString(id) {
		return false
	}

	// 2. แปลงแต่ละตัวเป็นเลข
	digits := make([]int, 13)
	for i, r := range id {
		d, _ := strconv.Atoi(string(r))
		digits[i] = d
	}

	// 3. หลัก 1 ต้องอยู่ในช่วง 1-8 และหลัก 2-5 ต้องไม่เป็น 0000 (ระบุจังหวัดอย่างหลวม)
	if digits[0] < 1 || digits[0] > 8 {
		return false
	}
	provinceCode := digits[1]*1000 + digits[2]*100 + digits[3]*10 + digits[4]
	if provinceCode == 0 {
		return false
	}

	// 4. น้ำหนักสำหรับหลักที่ 1-12
	weights := []int{13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	// 4. คำนวณ sum
	sum := 0
	for i := 0; i < 12; i++ {
		sum += digits[i] * weights[i]
	}

	// 5. คำนวณ check digit
	mod := sum % 11
	check := (11 - mod) % 10

	// 6. เปรียบเทียบกับ digit ที่ 13
	return check == digits[12]
}

// Normalize ลบช่องว่าง/ยัติภังค์ แล้วคืน 13 หลัก (ใช้ร่วมกับ IsValid)
func Normalize(id string) string {
	re := regexp.MustCompile(`[^\d]`)
	return re.ReplaceAllString(id, "")
}

// IsValidWithNormalize ตรวจสอบหลังลบสัญลักษณ์รบกวน
func IsValidWithNormalize(id string) bool {
	return IsValid(Normalize(id))
}
