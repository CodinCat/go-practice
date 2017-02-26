package practice

func Fibonacci(n int) []int {
	a := 0
	b := 1
	fib := []int{b}
	for i := 1; i < n; i++ {
		next := a + b
		fib = append(fib, next)
		a, b = b, next
	}
	return fib
}
