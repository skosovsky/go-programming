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
}
