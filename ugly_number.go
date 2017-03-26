package practice

func IsUglyNumber(n int) bool {
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

func UglyNumbers(length int) []int {
	uglyNumbers := make([]int, 0, length)
	for i := 1; len(uglyNumbers) != length; i++ {
		if IsUglyNumber(i) {
			uglyNumbers = append(uglyNumbers, i)
		}
	}
	return uglyNumbers
}
