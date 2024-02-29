package main

import "fmt"

func main() {
	var num, powerNum int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	switch {
	case num > 0:
		for powerNum*2 <= num {
			if powerNum == 0 {
				powerNum = 1
				fmt.Print(powerNum, " ")
			}

			powerNum *= 2
			if powerNum <= num {
				fmt.Print(powerNum, " ")
			}
		}
	case num < 0:
		for powerNum*2 > num {
			if powerNum == 0 {
				powerNum = -1
				fmt.Print(powerNum, " ")
			}

			powerNum *= 2
			if powerNum > num {
				fmt.Print(powerNum, " ")
			}
		}
	default:
		fmt.Print(powerNum, " ")
	}
}
