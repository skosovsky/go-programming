package main

import (
	"fmt"
)

func main() {
	var nums, num, sum int

	_, err := fmt.Scan(&nums) // а можно так fmt.Scanf("%1d%1d%1d",&a,&b,&c)
	if err != nil {
		return
	}

	num = nums
	for num != 0 {
		sum += num % 10
		num /= 10
	}

	fmt.Print(sum)
}
