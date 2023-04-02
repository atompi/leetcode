package lc605

func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 {
		return n+flowerbed[0] <= 1
	}
	for i := 0; i < len(flowerbed); {
		sum := 0
		if i == 0 {
			sum = flowerbed[i] + flowerbed[i+1]
		} else if i == len(flowerbed)-1 {
			sum = flowerbed[i-1] + flowerbed[i]
		} else {
			sum = flowerbed[i-1] + flowerbed[i] + flowerbed[i+1]
		}
		if sum == 0 {
			n--
			i += 2
		} else {
			i++
		}
		if n == 0 {
			break
		}
	}
	return n == 0
}
