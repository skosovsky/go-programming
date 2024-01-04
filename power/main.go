package main

import "fmt"

func main() {
	var num, result int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	result = num * num // math.Pow(num, 2)
	fmt.Println(result)
}
