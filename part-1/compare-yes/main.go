package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	num1 := num / 100
	num2 := num / 10 % 10
	num3 := num % 10

	if num1 != num2 && num1 != num3 && num2 != num3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// Opinion 2, with string
	var a string
	_, err = fmt.Scan(&a)
	if err != nil {
		return
	}

	if a[0] == a[1] || a[0] == a[2] || a[1] == a[2] {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
