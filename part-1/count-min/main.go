package main

import (
	"fmt"
	"math"
)

func main() {
	var count, numMin, countMin int

	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	numMin = math.MaxInt

	for count != 0 {
		var num int

		_, err := fmt.Scan(&num)
		if err != nil {
			return
		}

		if num < numMin {
			numMin = num
			countMin = 1
		} else if num == numMin {
			countMin++
		}

		count--
	}
	fmt.Print(countMin)
}
