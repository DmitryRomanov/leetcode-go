package make_string_great

import "unicode"

func makeGood(s string) string {
	chars := []rune(s)
	var stack []rune

	for i := range chars {
		if len(stack) == 0 {
			stack = append(stack, chars[i])
		} else {
			if unicode.ToLower(chars[i]) == unicode.ToLower(stack[len(stack)-1]) && chars[i] != stack[len(stack)-1] {
				_, stack = stack[len(stack)-1], stack[:len(stack)-1]
			} else {
				stack = append(stack, chars[i])
			}
		}
	}

	return string(stack)
}
