package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		return
	}

	legsSquared := a*a + b*b
	hypotenuseSquared := c * c

	if legsSquared == hypotenuseSquared {
		fmt.Print("Прямоугольный")
	} else {
		fmt.Print("Непрямоугольный")
	}
}
