package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHanoiOfLength1(t *testing.T) {
	assert := assert.New(t)

	s1 := NewStack(1)
	var s2 Stack
	var s3 Stack

	Hanoi(&s1, &s2, &s3, 1)
	assert.Equal(0, len(s1), "s1 should be empty")
	assert.Equal(0, len(s2), "s2 should be empty")
	assert.Equal(1, s3.Pop())
}

func TestHanoiOfLength3(t *testing.T) {
	assert := assert.New(t)

	s1 := NewStack(3, 2, 1)
	var s2 Stack
	var s3 Stack

	Hanoi(&s1, &s2, &s3, 3)
	assert.Equal(0, len(s1), "s1 should be empty")
	assert.Equal(0, len(s2), "s2 should be empty")
	assert.Equal(1, s3.Pop())
	assert.Equal(2, s3.Pop())
	assert.Equal(3, s3.Pop())
}

func TestMove(t *testing.T) {
	assert := assert.New(t)

	s1 := NewStack(1, 2, 3)
	var s2 Stack

	move(&s1, &s2)
	assert.Equal(3, s2.Pop())
	move(&s1, &s2)
	assert.Equal(2, s2.Pop())
	assert.Equal(1, s1.Pop())
}
