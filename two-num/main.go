package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	findSameIfs(a, b)
	findSameFors(a, b)
	findSameRecursive(a, b)
	findSameFors2(a, b)
	findSameStrings()
	findSameFors3(a, b)
}

func findSameIfs(a, b int) {
	var aNum1, aNum2, aNum3, aNum4 int
	var bNum1, bNum2, bNum3, bNum4 int
	var result string

	if a < 10 {
		aNum1 = a
	} else if a < 100 {
		aNum2 = a % 10
		aNum1 = a / 10
	} else if a < 1000 {
		aNum3 = a % 10
		aNum2 = a / 10 % 10
		aNum1 = a / 100
	} else if a < 10000 {
		aNum4 = a % 10
		aNum3 = a / 10 % 10
		aNum2 = a / 100 % 10
		aNum1 = a / 1000
	}

	if b < 10 {
		bNum1 = b
	} else if b < 100 {
		bNum2 = b % 10
		bNum1 = b / 10
	} else if b < 1000 {
		bNum3 = b % 10
		bNum2 = b / 10 % 10
		bNum1 = b / 100
	} else if b < 10000 {
		bNum4 = b % 10
		bNum3 = b / 10 % 10
		bNum2 = b / 100 % 10
		bNum1 = b / 1000
	}

	if a < 10000 && (aNum1 == bNum1 || aNum1 == bNum2 || aNum1 == bNum3 || aNum1 == bNum4) {
		if result == "" {
			result = fmt.Sprintf("%d", aNum1)
		} else {
			result += fmt.Sprintf(" %d", aNum1)
		}
	}
	if a >= 10 && a < 10000 && (aNum2 == bNum1 || aNum2 == bNum2 || aNum2 == bNum3 || aNum2 == bNum4) {
		if result == "" {
			result = fmt.Sprintf("%d", aNum2)
		} else {
			result += fmt.Sprintf(" %d", aNum2)
		}
	}
	if a >= 100 && a < 10000 && (aNum3 == bNum1 || aNum3 == bNum2 || aNum3 == bNum3 || aNum3 == bNum4) {
		if result == "" {
			result = fmt.Sprintf("%d", aNum3)
		} else {
			result += fmt.Sprintf(" %d", aNum3)
		}
	}
	if a >= 1000 && a < 10000 && (aNum4 == bNum1 || aNum4 == bNum2 || aNum4 == bNum3 || aNum4 == bNum4) {
		if result == "" {
			result = fmt.Sprintf("%d", aNum4)
		} else {
			result += fmt.Sprintf(" %d", aNum4)
		}
	}

	fmt.Printf("%s", result)
}

func findSameFors(a, b int) {
	var an, bn int

	for i := 1000; i >= 1 && a > 0; i = i / 10 {
		if a/i == 0 {
			continue
		}
		an = a / i % 10
		for i := 1; i < 10000 && b > 0; i = i * 10 {
			if b < i {
				break
			}
			bn = b / i % 10
			if bn == an {
				fmt.Print(an, " ")
				break
			}
		}
	}
}

func findSameRecursive(x, y int) {
	if x > 10 {
		findSameRecursive(x/10, y)
		findSameRecursive(x%10, y)
	} else if y > 10 {
		findSameRecursive(x, y/10)
		findSameRecursive(x, y%10)
	} else if x == y {
		fmt.Print(x, " ")
	}
}

func findSameFors2(x, y int) {
	var yTemp, rev int
	for ; x > 0; x /= 10 {
		for yTemp = y; yTemp > 0; yTemp /= 10 {
			if x%10 == yTemp%10 {
				rev = rev*10 + x%10
			}
		}
	}
	for ; rev > 0; rev /= 10 {
		fmt.Print(rev%10, " ")
	}
}

func findSameStrings() {
	var first, second string
	_, err := fmt.Scan(&first, &second)
	if err != nil {
		return
	}

	for i := 0; i < len(first); i++ {
		if strings.Contains(second, string(first[i])) {
			fmt.Print(string(first[i]), " ")
		}
	}
}

func findSameFors3(a, b int) {
	for i := 1000; i > 0; i /= 10 {
		if (a / i) != 0 {
			n := (a / i) % 10
			for j := 1000; j > 0; j /= 10 {
				if (b / j) != 0 {
					k := (b / j) % 10
					if n == k {
						fmt.Print(n, " ")
						break
					}
				}
			}
		}
	}
}
