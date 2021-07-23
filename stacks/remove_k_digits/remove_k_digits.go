package remove_k_digits

func removeKdigits(num string, k int) string {
	nums := []rune(num)
	var stack []rune

	if len(nums) == 1 {
		return "0"
	}

	for i := range nums {
		add := true
		for k > 0 {
			if len(stack) > 0 && stack[len(stack)-1] > nums[i] {
				// remove (i-1)
				_, stack = stack[len(stack)-1], stack[:len(stack)-1]
				k--
			} else if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
				add = false
				k--
				break
			} else {
				break
			}
		}

		if i+1 == len(nums) && k > 0 {
			for k > 0 && len(stack) > 0 {
				_, stack = stack[len(stack)-1], stack[:len(stack)-1]
				k--
			}
		}

		// remove leading zero
		if add && (len(stack) > 0 || string(nums[i]) != "0") {
			stack = append(stack, nums[i])
		}
	}

	// empty string
	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
