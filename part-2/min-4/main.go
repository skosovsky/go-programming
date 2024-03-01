package main

import "fmt"

func main() {
	minimumFromFour()
}

func minimumFromFour() int {
	var a, b, c, d, minNum int
	_, err := fmt.Scan(&a, &b, &c, &d)
	if err != nil {
		return 0
	}

	minNum = a

	if b < minNum {
		minNum = b
	}
	if c < minNum {
		minNum = c
	}
	if d < minNum {
		minNum = d
	}

	return minNum
}
