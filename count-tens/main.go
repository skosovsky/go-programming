package main

import "fmt"

func main() {
	var num, result int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	result = num / 10 % 10
	fmt.Println(result)

	// Opinion 2, with string
	var a string
	_, err = fmt.Scan(&a)
	if err != nil {
		return
	}

	fmt.Print(string(a[len(a)-2]))
}
