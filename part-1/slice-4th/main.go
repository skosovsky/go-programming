package main

import "fmt"

func main() {
	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	nums := make([]int, count)

	for i := 0; i < count; i++ {
		_, err := fmt.Scan(&nums[i])
		if err != nil {
			return
		}
	}

	fmt.Print(nums[3])
}
