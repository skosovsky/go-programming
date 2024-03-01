package main

import "fmt"

func main() {
	var a, b, average float64

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	average = (a + b) / 2
	fmt.Print(average)
}
