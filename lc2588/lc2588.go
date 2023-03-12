package lc2588

func BeautifulSubarrays(nums []int) int64 {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] ^ nums[i-1]
	}
	count := make(map[int]int)
	count[0] = 1
	var res int64 = 0
	for i := 1; i <= n; i++ {
		res += int64(count[prefix[i]])
		count[prefix[i]]++
	}
	return res
}
