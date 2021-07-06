package pathsum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPathSumExample1(t *testing.T) {
	assert := assert.New(t)
	result := hasPathSum(&TreeNode{}, 22)
	assert.True(result)
}
