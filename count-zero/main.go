package main

import "fmt"

func main() {
	var count, countZero int

	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	for count != 0 {
		var num int

		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		if num == 0 {
			countZero++
		}

		count--
	}

	fmt.Print(countZero)
}
