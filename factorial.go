package practice

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func FactorialLoop(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
