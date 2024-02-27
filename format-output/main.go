package main

import "fmt"

func main() {
	var num float64
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	formatIfs(num)
	formatCases(num)
	formatCasesString(num)
}

func formatIfs(num float64) {
	numPower := num * num

	if num <= 0 {
		fmt.Printf("число %2.2f не подходит", num)
		return
	}

	if num > 10000 {
		fmt.Printf("%e", num)
		return
	}

	fmt.Printf("%.4f", numPower)
}

func formatCases(in float64) {
	switch {
	case in <= 0:
		fmt.Printf("число %2.2f не подходит", in)
	case in > 10000:
		fmt.Printf("%e", in)
	default:
		fmt.Printf("%.4f", in*in)
	}
}

func formatCasesString(a float64) {
	var f string

	switch {
	case a <= 0:
		f = "число %2.2f не подходит"
	case a > 10_000:
		f = "%e"
	default:
		a *= a
		f = "%.4f"
	}
	fmt.Printf(f, a)
}
