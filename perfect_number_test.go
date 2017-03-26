package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPerfectNumber(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(false, IsPerfectNumber(1), "1")
	assert.Equal(false, IsPerfectNumber(2), "2")
	assert.Equal(false, IsPerfectNumber(3), "3")
	assert.Equal(false, IsPerfectNumber(4), "4")
	assert.Equal(false, IsPerfectNumber(5), "5")
	assert.Equal(true, IsPerfectNumber(6), "6")
	assert.Equal(false, IsPerfectNumber(7), "7")
	assert.Equal(false, IsPerfectNumber(8), "8")
	assert.Equal(false, IsPerfectNumber(9), "9")
	assert.Equal(false, IsPerfectNumber(10), "10")
	assert.Equal(false, IsPerfectNumber(11), "11")
	assert.Equal(false, IsPerfectNumber(12), "12")
	assert.Equal(false, IsPerfectNumber(13), "13")
	assert.Equal(false, IsPerfectNumber(14), "14")
	assert.Equal(false, IsPerfectNumber(20), "20")
	assert.Equal(false, IsPerfectNumber(26), "26")
	assert.Equal(true, IsPerfectNumber(28), "28")
	assert.Equal(false, IsPerfectNumber(29), "29")
	assert.Equal(false, IsPerfectNumber(30), "30")
	assert.Equal(false, IsPerfectNumber(42), "42")
	assert.Equal(false, IsPerfectNumber(100), "100")
	assert.Equal(false, IsPerfectNumber(200), "200")
	assert.Equal(false, IsPerfectNumber(400), "400")
	assert.Equal(true, IsPerfectNumber(496), "496")
	assert.Equal(true, IsPerfectNumber(8128), "8128")
}

func TestPerfectNumbers(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		[]int{6, 28, 496, 8128},
		PerfectNumbers(4),
	)
}
