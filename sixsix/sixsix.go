package sixsix

func PlusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if carry == 1 {
			tmp := digits[i] + 1
			carry = tmp / 10
			digits[i] = tmp % 10
		} else {
			continue
		}
	}

	if carry == 1 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}
