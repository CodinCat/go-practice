package practice

func Prime(length int) []int {
	var result []int
	n := 2
	for len(result) != length {
		if OptimizedIsPrime(n) {
			result = append(result, n)
		}
		n++
	}
	return result
}

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	limit := n/2 + 1
	for divisor := 2; divisor < limit; divisor++ {
		if n%divisor == 0 {
			return false
		}
	}
	return true
}

func OptimizedIsPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	limit := n/2 + 1
	for divisor := 5; divisor < limit; divisor += 6 {
		if n%divisor == 0 || n%(divisor+2) == 0 {
			return false
		}
	}
	return true
}
