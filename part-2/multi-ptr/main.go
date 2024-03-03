package main

import "fmt"

func main() {
	a, b := 2, 2
	test(&a, &b)
}

func test(x1 *int, x2 *int) {
	result := *x1 * *x2
	fmt.Print(result)
}
