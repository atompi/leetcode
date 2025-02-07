package lc20

func IsValid(s string) bool {
	if len(s) < 2 || len(s)%2 != 0 {
		return false
	}
	m := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := make([]rune, 0)
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')', '}', ']':
			if len(stack) == 0 || m[string(v)] != string(stack[len(stack)-1]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
