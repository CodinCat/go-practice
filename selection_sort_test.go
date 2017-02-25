package practice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	assert := assert.New(t)
	var testCases = []struct {
		expected,
		input []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{100, 200, 300, 500}, []int{100, 200, 500, 300}},
		{[]int{-100, -1, 0, 0, 2}, []int{0, 2, -1, -100, 0}},
	}
	for _, testCase := range testCases {
		msg := fmt.Sprintf("Input: %v", testCase.input)
		SelectionSort(testCase.input)
		assert.Equal(testCase.expected, testCase.input, msg)
	}
}

func TestSelectMin(t *testing.T) {
	assert := assert.New(t)
	var testCases = []struct {
		expected   int
		inputSlice []int
		inputFrom  int
	}{
		{0, []int{1}, 0},
		{0, []int{1, 9}, 0},
		{3, []int{9, 8, 1, 0}, 0},
		{2, []int{70, 81, 3, 22}, 0},
		{6, []int{70, 81, 1, 22, 99, 100, 3, 8}, 3},
	}
	for _, testCase := range testCases {
		assert.Equal(testCase.expected, selectMin(testCase.inputSlice, testCase.inputFrom))
	}
}
