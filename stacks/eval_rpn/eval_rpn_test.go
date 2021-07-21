package eval_rpn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalRPNExample1(t *testing.T) {
	assert := assert.New(t)
	result := evalRPN([]string{"2", "1", "+", "3", "*"})
	assert.Equal(9, result)
}

func TestEvalRPNExample2(t *testing.T) {
	assert := assert.New(t)
	result := evalRPN([]string{"4", "13", "5", "/", "+"})
	assert.Equal(6, result)
}

func TestEvalRPNExample3(t *testing.T) {
	assert := assert.New(t)
	result := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	assert.Equal(22, result)
}
