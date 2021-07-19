package remove_dublicates

func removeDuplicates(s string) string {
	chars := []rune(s)
	var stack []rune

	for i := range chars {
		if len(stack) > 0 {
			if stack[len(stack)-1] == chars[i] {
				_, stack = stack[len(stack)-1], stack[:len(stack)-1]
			} else {
				stack = append(stack, chars[i])
			}
		} else {
			stack = append(stack, chars[i])
		}
	}
	return string(stack)
}
