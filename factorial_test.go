package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, Factorial(1))
	assert.Equal(2, Factorial(2))
	assert.Equal(6, Factorial(3))
	assert.Equal(24, Factorial(4))
	assert.Equal(120, Factorial(5))
}

func TestFactorialLoop(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, FactorialLoop(1))
	assert.Equal(2, FactorialLoop(2))
	assert.Equal(6, FactorialLoop(3))
	assert.Equal(24, FactorialLoop(4))
	assert.Equal(120, FactorialLoop(5))
}
