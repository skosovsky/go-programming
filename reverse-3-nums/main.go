package main

import (
	"fmt"
	"math"
)

func main() {
	var nums, num, reverseNums, lenNums int
	_, err := fmt.Scan(&nums)
	if err != nil {
		return
	}

	lenNums = int(math.Floor(math.Log10(float64(nums)) + 1))
	num = nums
	for i := lenNums; num != 0; i-- {
		reverseNums += num % 10 * int(math.Pow(10, float64(i-1))) // или reverseNums = reverseNums * 10 + num % 10
		num /= 10
	}

	fmt.Print(reverseNums)
}
