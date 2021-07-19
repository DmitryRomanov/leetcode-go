package valid_parentless

func isValid(s string) bool {
	chars := []rune(s)
	var stack []rune
	for i := range chars {
		if len(stack) > 0 {
			diffChars := chars[i] - stack[len(stack)-1]
			if diffChars > 0 && diffChars < 3 {
				_, stack = stack[len(stack)-1], stack[:len(stack)-1]
			} else {
				stack = append(stack, chars[i])
			}
		} else {
			stack = append(stack, chars[i])
		}
	}

	return len(stack) == 0
}
