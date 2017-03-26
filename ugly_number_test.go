package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUglyNumber(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, IsUglyNumber(1), "1")
	assert.Equal(true, IsUglyNumber(2), "2")
	assert.Equal(true, IsUglyNumber(3), "3")
	assert.Equal(true, IsUglyNumber(4), "4")
	assert.Equal(true, IsUglyNumber(5), "5")
	assert.Equal(true, IsUglyNumber(6), "6")
	assert.Equal(false, IsUglyNumber(7), "7")
	assert.Equal(true, IsUglyNumber(8), "8")
	assert.Equal(true, IsUglyNumber(9), "9")
	assert.Equal(true, IsUglyNumber(10), "10")
	assert.Equal(false, IsUglyNumber(11), "11")
	assert.Equal(true, IsUglyNumber(12), "12")
	assert.Equal(false, IsUglyNumber(13), "13")
	assert.Equal(false, IsUglyNumber(14), "14")
	assert.Equal(true, IsUglyNumber(15), "15")
	assert.Equal(true, IsUglyNumber(16), "16")
	assert.Equal(false, IsUglyNumber(17), "17")
	assert.Equal(true, IsUglyNumber(18), "18")
	assert.Equal(false, IsUglyNumber(19), "19")
	assert.Equal(true, IsUglyNumber(20), "20")
	assert.Equal(false, IsUglyNumber(21), "21")
	assert.Equal(false, IsUglyNumber(22), "22")
	assert.Equal(false, IsUglyNumber(23), "23")
	assert.Equal(true, IsUglyNumber(24), "24")
	assert.Equal(true, IsUglyNumber(25), "25")
	assert.Equal(false, IsUglyNumber(26), "26")
}

func TestUglyNumbers(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		[]int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18},
		UglyNumbers(13),
	)
}
