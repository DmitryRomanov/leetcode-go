package firstbad

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	badVersion = 4
	assert := assert.New(t)
	result := firstBadVersion(5)
	assert.Equal(badVersion, result)
}
