package crawler_log_files

func minOperations(logs []string) int {
	var stack []string
	for i := range logs {
		if logs[i] == "../" {
			if len(stack) > 0 {
				_, stack = stack[len(stack)-1], stack[:len(stack)-1]
			}
		} else if logs[i] == "./" {
			// nothing
		} else {
			stack = append(stack, logs[i])
		}
	}
	return len(stack)
}
