package main

import "fmt"

func main() {
	var a, b, sum int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	for ; a <= b; a++ {
		sum += a
	}

	fmt.Println(sum)

	// Opinion 2, smart
	fmt.Println((b - a + 1) * (a + b) / 2)
}
