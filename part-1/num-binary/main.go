package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	var binaryReverse, binary string
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	for num != 0 {
		strconv.Itoa(num)
		binaryReverse += strconv.Itoa(num % 2)
		num /= 2
	}

	for i := len(binaryReverse) - 1; i >= 0; i-- {
		binary += string(binaryReverse[i])
	}

	fmt.Print(binary) // или fmt.Printf("%b", num)
}
