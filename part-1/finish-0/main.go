package main

import "fmt"

func main() {
	var num, maxNum, count int

	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		if num == 0 {
			break
		}

		if num > maxNum {
			maxNum = num
			count = 1
		} else if num == maxNum {
			count++
		}
	}

	fmt.Println(count)

	//Opinion 2, with Scan in for
	var a, m, counta int

	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		switch {
		case a > m:
			m, counta = a, 1
		case a == m:
			counta++
		}
	}
	fmt.Println(counta)

	// Opinion 3, without break
	var a2, b2, c2 = -1, 0, 0
	for a2 != 0 {
		fmt.Scan(&a)
		if a2 > b2 {
			b2, c2 = a2, 1
		} else if a2 == b2 {
			c2++
		}
	}
	fmt.Println(c2)
}
