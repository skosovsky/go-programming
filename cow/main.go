package main

import "fmt"

func main() {
	var num, lastNum int
	var description string
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	lastNum = num % 10

	if num <= 20 {
		switch num {
		case 0:
			description = "korov"
		case 1:
			description = "korova"
		case 2, 3, 4:
			description = "korovi"
		default:
			description = "korov"
		}
	} else {
		switch lastNum {
		case 0:
			description = "korov"
		case 1:
			description = "korova"
		case 2, 3, 4:
			description = "korovi"
		default:
			description = "korov"
		}
	}

	fmt.Print(num, " ", description)
}
