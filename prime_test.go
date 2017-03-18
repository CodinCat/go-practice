package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type isPrimeFunc func(n int) bool

func TestPrime(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{2}, Prime(1))
	assert.Equal([]int{2, 3, 5}, Prime(3))
	assert.Equal(
		[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		Prime(15),
	)
}

func TestIsPrimeNaive(t *testing.T) {
	testIsPrime(t, IsPrime)
}

func TestOptimizedIsPrime(t *testing.T) {
	testIsPrime(t, OptimizedIsPrime)
}

func testIsPrime(t *testing.T, fn isPrimeFunc) {
	assert := assert.New(t)
	assert.Equal(false, fn(1), "1")
	assert.Equal(true, fn(2), "2")
	assert.Equal(true, fn(3), "3")
	assert.Equal(false, fn(4), "4")
	assert.Equal(true, fn(5), "5")
	assert.Equal(false, fn(6), "6")
	assert.Equal(true, fn(7), "7")
	assert.Equal(true, fn(11), "11")
	assert.Equal(true, fn(13), "13")
	assert.Equal(false, fn(15), "15")
	assert.Equal(true, fn(17), "17")
	assert.Equal(false, fn(21), "21")
	assert.Equal(false, fn(22), "22")
	assert.Equal(true, fn(23), "23")
	assert.Equal(false, fn(35), "35")
	assert.Equal(false, fn(50), "50")
	assert.Equal(false, fn(51), "51")
	assert.Equal(false, fn(52), "52")
	assert.Equal(true, fn(53), "53")
	assert.Equal(false, fn(55), "55")
	assert.Equal(true, fn(89), "89")
	assert.Equal(true, fn(97), "97")
	assert.Equal(true, fn(131), "131")
	assert.Equal(false, fn(133), "133")
	assert.Equal(true, fn(137), "137")
	assert.Equal(true, fn(139), "139")
	assert.Equal(false, fn(143), "143")
	assert.Equal(false, fn(161), "161")
	assert.Equal(false, fn(169), "169")
	assert.Equal(false, fn(187), "187")
}
