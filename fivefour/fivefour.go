package fivefour

func SpiralOrder(matrix [][]int) []int {
	elementsCount := len(matrix) * len(matrix[0])
	i, j := 0, 0
	ans := make([]int, 0)
	xMax, yMax, xMin, yMin := len(matrix[0]), len(matrix), 0, 1

	for elementsCount > 0 {
		for j < xMax && elementsCount > 0 {
			ans = append(ans, matrix[i][j])
			j++
			elementsCount--
		}
		xMax = j - 1
		j -= 1
		i += 1
		for i < yMax && elementsCount > 0 {
			ans = append(ans, matrix[i][j])
			i++
			elementsCount--
		}
		yMax = i - 1
		i -= 1
		j -= 1
		for j >= xMin && elementsCount > 0 {
			ans = append(ans, matrix[i][j])
			j--
			elementsCount--
		}
		xMin += 1
		j += 1
		i -= 1
		for i > yMin && elementsCount > 0 {
			ans = append(ans, matrix[i][j])
			i--
			elementsCount--
		}
		yMin += 1
	}

	return ans
}
