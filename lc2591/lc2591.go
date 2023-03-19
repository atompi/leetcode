package clc6323

func DistMoney(money int, children int) int {
	if money < children {
		return -1
	}

	ans := 0

	if money/8 > children {
		return children - 1
	} else if money/8 == children {
		if money%8 == 0 {
			return children
		} else {
			return children - 1
		}
	} else {
		after := money - children
		for after >= 7 {
			after -= 7
			ans += 1
		}
		if children-ans == 1 {
			if after == 3 {
				ans -= 1
			}
		}
	}

	return ans
}
