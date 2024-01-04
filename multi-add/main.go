package main

import "fmt"

func main() {
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		return
	}

	result := a*2 + 100
	resultSmart := a<<1 + 100
	fmt.Println(result, resultSmart)
}
