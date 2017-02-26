package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fib = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233}

func TestFibonacci(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(fib[:1], Fibonacci(1))
	assert.Equal(fib[:2], Fibonacci(2))
	assert.Equal(fib[:5], Fibonacci(5))
	assert.Equal(fib[:13], Fibonacci(13))
}
