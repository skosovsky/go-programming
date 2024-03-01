package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	if num > 0 {
		fmt.Println("Число положительное")
	} else if num < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}

	// Opinion 2, with case

	switch {
	case num > 0:
		fmt.Println("Число положительное")
	case num < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Ноль")
	}
}
