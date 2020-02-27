package vector

func plusOne(digits []int) []int {
	var carry int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i], carry = (digits[i]+carry)%10, (digits[i]+carry)/10
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
