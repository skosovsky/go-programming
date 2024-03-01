package main

import "fmt"

func main() {
	var x, y, z int
	_, err := fmt.Scan(&x, &y, &z)
	if err != nil {
		return
	}

	fmt.Print(vote(x, y, z))
}

func vote(x int, y int, z int) int {
	if x+y+z <= 1 {
		return 0
	}

	return 1 // или (x + y + z) / 2 или x & y | x & z | y & z
}
