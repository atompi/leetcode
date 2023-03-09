package lc989

func AddToArrayForm(num []int, k int) []int {
	kArray := []int{}
	for k > 0 {
		kArray = append(kArray, k%10)
		k = int(k / 10)
	}
	for i, j := 0, len(kArray)-1; i < j; i, j = i+1, j-1 {
		kArray[i], kArray[j] = kArray[j], kArray[i]
	}

	ans := []int{}
	i := len(num) - 1
	j := len(kArray) - 1
	carry := 0

	for {
		if j >= 0 && i >= 0 {
			c := num[i] + kArray[j] + carry
			carry = c / 10
			ans = append([]int{c % 10}, ans...)
			i--
			j--
		} else if i < 0 && j >= 0 {
			c := kArray[j] + carry
			carry = c / 10
			ans = append([]int{c % 10}, ans...)
			j--
		} else if i >= 0 && j < 0 {
			c := num[i] + carry
			carry = c / 10
			ans = append([]int{c % 10}, ans...)
			i--
		} else {
			if carry == 1 {
				ans = append([]int{carry}, ans...)
			}
			break
		}
	}

	return ans
}
