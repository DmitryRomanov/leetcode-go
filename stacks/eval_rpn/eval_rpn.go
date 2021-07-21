package eval_rpn

import (
	"strconv"
)

var (
	operations = []string{"/", "*", "+", "-"}
)

func evalRPN(tokens []string) int {
	var stack []int
	var operand1, operand2 int
	for i := range tokens {
		if isOperation(tokens[i]) {
			operand2, stack = pop(stack)
			operand1, stack = pop(stack)

			if tokens[i] == "/" {
				stack = append(stack, operand1/operand2)
			} else if tokens[i] == "*" {
				stack = append(stack, operand1*operand2)
			} else if tokens[i] == "+" {
				stack = append(stack, operand1+operand2)
			} else if tokens[i] == "-" {
				stack = append(stack, operand1-operand2)
			}
		} else {
			value, _ := strconv.Atoi(tokens[i])
			stack = append(stack, value)
		}
	}

	return stack[0]
}

func pop(tokens []int) (int, []int) {
	return tokens[len(tokens)-1], tokens[:len(tokens)-1]
}

func isOperation(token string) bool {
	for _, v := range operations {
		if v == token {
			return true
		}
	}

	return false
}
