package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	num1 := num / 100000
	num2 := num / 10000 % 10
	num3 := num / 1000 % 10
	num4 := num / 100 % 10
	num5 := num / 10 % 10
	num6 := num / 1 % 10

	sum1 := num1 + num2 + num3
	sum2 := num4 + num5 + num6

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// Opinion 2, use code from string
	var s string
	_, err = fmt.Scan(&s)
	if err != nil {
		return
	}
	if s[0]+s[1]+s[2] == s[3]+s[4]+s[5] {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
