package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)
	var stack Stack
	stack.Push(5)
	stack.Push(6)
	assert.Equal(stack.Pop(), 6)
	assert.Equal(stack.Pop(), 5)
	stack.Push(1)
	assert.Equal(stack.Pop(), 1)
}
