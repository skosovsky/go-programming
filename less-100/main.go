package main

import "fmt"

func main() {
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		if num < 10 {
			continue
		} else if num > 100 {
			break
		}

		fmt.Println(num)
	}

	// Opinion 2, with Scan in for
	var n int
	for fmt.Scan(&n); n <= 100; fmt.Scan(&n) {
		if n >= 10 {
			fmt.Println(n)
		}
	}

	// Opinion 3, with instruction in switch
	var i int
	for {
		switch fmt.Scan(&i); {
		case i > 100:
			return
		case i >= 10:
			fmt.Println(i)
		}
	}
}
