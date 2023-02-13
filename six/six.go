package six

func Convert(s string, numRows int) string {
	convMatrix, down, up := make([][]byte, numRows), 0, numRows-2
	for i := 0; i < len(s); {
		if down < numRows {
			convMatrix[down] = append(convMatrix[down], s[i])
			down++
			i++
		} else if up > 0 {
			convMatrix[up] = append(convMatrix[up], s[i])
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	result := make([]byte, 0, len(s))
	for _, row := range convMatrix {
		result = append(result, row...)
	}
	return string(result)
}
