package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	fmt.Print(fibonacci(num))
	fmt.Print(fibonacciSmart(num))
	fmt.Print(fibonacciRecursiveSlow(num))
	fmt.Print(fibonacciRecursiveEnd(num))
}

func fibonacci(n int) int {
	switch n {
	case 1, 2:
		return 1
	case 3:
		return 2
	default:
		a, b, c, count := 1, 1, 2, 3 //nolint:ineffassign //a used

		for n != count {
			a, b = b, c
			c = a + b
			count++
		}

		return c
	}
}

func fibonacciSmart(n int) int {
	first, second := 1, 1
	for i := 1; i < n; i++ {
		first, second = second, first+second
	}
	return first
}

func fibonacciRecursiveSlow(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacciRecursiveSlow(n-2) + fibonacciRecursiveSlow(n-1)
}

func fibonacciRecursiveEnd(n int) int {
	return helper(1, 0, n)
}
func helper(a, b, n int) int {
	if n == 1 {
		return a
	}

	return helper(a+b, a, n-1)
}
