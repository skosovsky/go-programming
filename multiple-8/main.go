package main

import "fmt"

func main() {
	var n, sum int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	for ; n > 0; n-- {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		if num >= 10 && num < 100 && num%8 == 0 {
			sum += num
		}
	}

	fmt.Println(sum)

	// Opinion 2, with Scan in for
	var n2, sum2 int
	for fmt.Scan(&n2); n2 > 0; n2-- {
		var x int
		if fmt.Scan(&x); x > 9 && x < 100 && x%8 == 0 {
			sum2 += x
		}
	}
	fmt.Println(sum2)
}
