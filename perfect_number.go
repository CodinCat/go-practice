package practice

func IsPerfectNumber(n int) bool {
	factorSum := 0
	limit := n/2 + 1
	for i := 1; i <= limit; i++ {
		if n%i == 0 {
			factorSum += i
		}
	}
	return factorSum == n && n != 1
}

func PerfectNumbers(length int) []int {
	perfectNumbers := make([]int, 0, length)
	for i := 1; len(perfectNumbers) < length; i++ {
		if IsPerfectNumber(i) {
			perfectNumbers = append(perfectNumbers, i)
		}
	}
	return perfectNumbers
}
