package main

import "fmt"

func main() {
	var num, deletedNum, withoutNum int
	var numSlcReverse []int

	_, err := fmt.Scan(&num, &deletedNum)
	if err != nil {
		return
	}

	deleteNumSmart(num, deletedNum) // Smart

	for num != 0 {
		numSlcReverse = append(numSlcReverse, num%10)
		num /= 10
	}

	for i := len(numSlcReverse) - 1; i >= 0; i-- {
		if numSlcReverse[i] != deletedNum {
			withoutNum = withoutNum*10 + numSlcReverse[i]
		}
	}

	fmt.Print(withoutNum)
}

func deleteNumSmart(a, b int) {
	result := 0
	for i := 1; a > 0; {
		if a%10 != b {
			result += a % 10 * i
			i *= 10
		}
		a /= 10
	}
	fmt.Println(result)
}
