package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := range workArray {
		_, err := fmt.Scan(&workArray[i])
		if err != nil {
			return
		}
	}

	for range 3 { // nolint:typecheck
		var a, b uint8
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			return
		}

		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	for _, v := range workArray {
		fmt.Printf("%d ", v)
	}
}
