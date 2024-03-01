package main

import "fmt"

func main() {
	var num, deletedNum, withoutNum int
	var numSlcReverse []int

	_, err := fmt.Scan(&num, &deletedNum)
	if err != nil {
		return
	}

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
