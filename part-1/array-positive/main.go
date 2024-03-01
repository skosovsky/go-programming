package main

import "fmt"

func main() {
	var count int
	var countPositiveNum int

	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	for i := 0; i < count; i++ {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		if num > 0 {
			countPositiveNum++
		}
	}

	fmt.Print(countPositiveNum)
}
