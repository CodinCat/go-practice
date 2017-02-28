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
