package main

import (
	"fmt"
)

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&a)
		if err != nil {
			return
		}
		array[i] = a
	}

	maxNum := array[0] // или math.MinInt
	for i := range array {
		if array[i] > maxNum {
			maxNum = array[i]
		}
	}

	fmt.Println(maxNum)
}
