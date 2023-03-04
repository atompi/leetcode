package twofourfourfour

func CountSubarrays(nums []int, minK int, maxK int) int64 {
	ans := 0
	j1, j2, k := -1, -1, -1
	for i, v := range nums {
		if v < minK || v > maxK {
			k = i
		}
		if v == minK {
			j1 = i
		}
		if v == maxK {
			j2 = i
		}
		ans += max(0, min(j1, j2)-k)
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ChatGPT wrong
func CountSubarrays2(nums []int, minK int, maxK int) int64 {
	n := len(nums)
	left, right := 0, 0
	var count int64 = 0

	// Move the left pointer to the first occurrence of minK
	for left < n && nums[left] != minK {
		left++
	}

	// Iterate over all possible right pointers
	for right < n {
		// If the right pointer points to a value outside the range [minK, maxK],
		// move the left pointer to the next occurrence of minK
		if nums[right] < minK || nums[right] > maxK {
			left = right + 1
		} else if nums[right] == maxK {
			// If the right pointer points to maxK, count the number of subarrays
			// that can be formed by fixing maxK and varying the position of minK
			j := right
			for j >= left && nums[j] == maxK {
				j--
			}
			count += int64(right-j) * int64(j-left+1)
			left = j + 1
		}
		right++
	}

	return count
}
