package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	if num <= 9 {
		fmt.Println(num)
	} else if num <= 99 {
		fmt.Println(num / 10)
	} else if num <= 999 {
		fmt.Println(num / 100)
	} else if num <= 9999 {
		fmt.Println(num / 1000)
	} else {
		fmt.Println(num / 10000)
	}

	// Opinion 2, Recursion
	fmt.Println(getFirstNumber(num))

	// Opinion 3, use switch
	var n int
	_, err = fmt.Scan(&n)
	if err != nil {
		return
	}

	switch {
	case n < 10:
		fmt.Println(n)
	case n < 100:
		fmt.Println(n / 10)
	case n < 1000:
		fmt.Println(n / 100)
	case n < 10000:
		fmt.Println(n / 1000)
	case n < 100000:
		fmt.Println(n / 10000)
	}

	// Opinion 4, use string
	var s string
	_, err = fmt.Scan(&s)
	if err != nil {
		return
	}
	fmt.Print(string(s[0]))

	// Opinion 5, use for, universal
	var x int
	_, err = fmt.Scan(&x)
	if err != nil {
		return
	}
	for x > 9 {
		x = x / 10
	}
	fmt.Print(x)

	// Opinion 6, smart
	var in float64
	_, err = fmt.Scan(&in)
	if err != nil {
		return
	}
	fmt.Print(int(in / math.Pow10(int(math.Log10(in)))))
}

// Opinion 2, Recursion
func getFirstNumber(a int) int {
	if a < 10 {
		return a
	} else {
		return getFirstNumber(a / 10)
	}
}
