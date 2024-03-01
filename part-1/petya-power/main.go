package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b) // считаем переменную 'a' и 'b' с консоли
	if err != nil {
		return
	}

	a *= a
	b *= b
	c = a + b
	fmt.Println(c)

	// Opinion 2, with Pow
	var a2, b2, c2 float64
	_, err = fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	c2 = math.Pow(a2, 2) + math.Pow(b2, 2)
	fmt.Println(int(c2))
}
