package main

import "fmt"

func main() {
	var x, p, y, year int
	_, err := fmt.Scan(&x, &p, &y)
	if err != nil {
		return
	}

	for year = 0; x <= y; year++ {
		percent := x * p / 100
		x += percent
	}

	fmt.Println(year)

	// Opinion 2, with strange for
	count := 0
	for ; x < y; x += x * p / 100 {
		count++
	}
	fmt.Println(count)
}
