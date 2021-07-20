package crawler_log_files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinOperationsExample1(t *testing.T) {
	assert := assert.New(t)
	result := minOperations([]string{"d1/", "d2/", "../", "d21/", "./"})
	assert.Equal(2, result)
}

func TestMinOperationsExample2(t *testing.T) {
	assert := assert.New(t)
	result := minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"})
	assert.Equal(3, result)
}

func TestMinOperationsExample3(t *testing.T) {
	assert := assert.New(t)
	result := minOperations([]string{"d1/", "../", "../", "../"})
	assert.Equal(0, result)
}
