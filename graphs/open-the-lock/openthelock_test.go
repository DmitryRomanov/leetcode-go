package open_the_lock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenLockExample1(t *testing.T) {
	assert := assert.New(t)
	result := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	assert.Equal(6, result)
}

func TestOpenLockExample2(t *testing.T) {
	assert := assert.New(t)
	result := openLock([]string{"8888"}, "0009")
	assert.Equal(1, result)
}

func TestOpenLockExample3(t *testing.T) {
	assert := assert.New(t)
	result := openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	assert.Equal(-1, result)
}

func TestOpenLockExample4(t *testing.T) {
	assert := assert.New(t)
	result := openLock([]string{"0000"}, "8888")
	assert.Equal(-1, result)
}
