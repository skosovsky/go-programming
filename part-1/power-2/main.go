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

	powerSimple(num)
	powerSimpleBit(num)
	powerSimple2(num)
}

func powerSimple(mas int) {
	for i := 1; i <= mas; i *= 2 { // или for i := 1; i <= mas; i += i {
		fmt.Print(i, " ")
	}
}

func powerSimpleBit(n int) {
	for d := 1; d <= n; d <<= 1 {
		fmt.Print(d, " ")
	}
}

func powerSimple2(n int) {
	var a int = 1
	for a <= n {
		fmt.Print(a, " ")
		a *= 2
	}
}
