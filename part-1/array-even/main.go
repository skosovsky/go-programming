package main

import "fmt"

func main() {
	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	nums := make([]int, count)

	for i := 0; i < count; i++ { // можно в этом же цикле
		_, err := fmt.Scan(&nums[i])
		if err != nil {
			return
		}
	}

	for i := range nums { // или for i:=0 ; i<n ; i +=2
		if i%2 == 0 {
			fmt.Print(nums[i], " ")
		}
	}
}
