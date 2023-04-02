package lc278

func FirstBadVersion(n int) int {
	low := 0
	high := n + 1
	for low <= high {
		middle := (low + high) / 2
		if !isBadVersion(middle) && isBadVersion(middle+1) {
			return middle + 1
		} else if isBadVersion(middle) && isBadVersion(middle+1) {
			if middle == 1 {
				return middle
			} else {
				high = middle
			}
		} else {
			low = middle + 1
		}
	}
	return low
}

func isBadVersion(version int) bool {
	return version == 5
}
