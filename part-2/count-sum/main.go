package main

import "fmt"

func main() {
	a, b := sumInt(1, 2, 3, 4, 5)

	fmt.Println(a, b)
}

func sumInt(args ...int) (int, int) {
	var sum int

	for _, v := range args {
		sum += v
	}

	return len(args), sum
}
