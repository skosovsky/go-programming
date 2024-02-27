package main

import "fmt"

func main() {
	var n, c, d int
	_, err := fmt.Scan(&n, &c, &d)
	if err != nil {
		return
	}

	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}

	// Opinion 2, smart
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}

	// Opinion 3, very smart
	if c%d != 0 && n >= c {
		fmt.Print(c)
	}
}
