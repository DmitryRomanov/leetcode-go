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
