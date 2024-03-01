package main

import "fmt"

func main() {
	var num, remainder, sum, root int

	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	remainder = num
	for remainder != 0 {
		sum += remainder % 10
		remainder /= 10
	}

	remainder = sum
	for remainder != 0 {
		root += remainder % 10
		remainder /= 10
	}

	fmt.Print(root)

	digitalRoot(num)

	digitalRootFor(num)

	digitalRootFormula(num)
}

func digitalRoot(num int) {
	var root int
	for num != 0 {
		root = root + num%10
		num = num / 10
	}
	if root < 10 {
		fmt.Print(root)
	} else {
		digitalRoot(root)
	}
}

func digitalRootFor(num int) {
	var a int
	for num > 9 {
		a = num % 10
		num /= 10
		num += a
	}
	fmt.Print(num)
}

func digitalRootFormula(num int) {
	fmt.Print((num-1)%9 + 1)
}
