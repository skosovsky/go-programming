package main

import "fmt"

func main() {
	a, b := 2, 4
	test(&a, &b)
}

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Print(*x1, *x2)
}
