package main

import "fmt"

func main() {
	var a, b int

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			fmt.Print(i)
			return
		}
	}
	fmt.Print("NO")
}
