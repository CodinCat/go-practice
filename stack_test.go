package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	assert := assert.New(t)
	stack := NewStack(1, 2, 3)
	assert.Equal(3, stack.Pop())
	assert.Equal(2, stack.Pop())
	assert.Equal(1, stack.Pop())
	stack.Push(10)
	assert.Equal(10, stack.Pop())
}

func TestStack(t *testing.T) {
	assert := assert.New(t)
	var stack Stack
	stack.Push(5)
	stack.Push(6)
	assert.Equal(6, stack.Pop())
	assert.Equal(5, stack.Pop())
	stack.Push(1)
	assert.Equal(1, stack.Pop())
}
