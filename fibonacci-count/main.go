package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	a, b, c, count := 0, 1, 1, 2

	for num >= c {
		a, b = b, c
		c = a + b
		count++

		if c == num {
			fmt.Print(count)
			return
		}
	}

	fmt.Print(-1)
}
