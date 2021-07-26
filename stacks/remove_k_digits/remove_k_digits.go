package remove_k_digits

func removeKdigits(num string, k int) string {
	nums := []rune(num)
	var stack, result []rune

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
			} else {
				break
			}
		}

		if add {
			stack = append(stack, nums[i])
		}

		if i+1 == len(nums) && k > 0 {
			for k > 0 && len(stack) > 0 {
				_, stack = stack[len(stack)-1], stack[:len(stack)-1]
				k--
			}
		}

	}

	// empty string
	if len(stack) == 0 {
		return "0"
	}

	zeroBegins := true

	for i := range stack {
		if zeroBegins && string(stack[i]) == "0" && i+1 != len(stack) {
			continue
		} else {
			zeroBegins = false
		}
		result = append(result, stack[i])
	}

	return string(result)
}
