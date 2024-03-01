package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		iPower := i * i
		fmt.Println(iPower) // или math.Pow(i,2)
	}
}
