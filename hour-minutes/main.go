package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		secInHour   = 3600
		secInMinute = 60
	)

	var seconds, hours, minutes int
	_, err := fmt.Scan(&seconds)
	if err != nil {
		return
	}

	hours = seconds / secInHour
	minutes = (seconds - hours*secInHour) / secInMinute

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)

	// Вариант с time
	s := time.Date(0, 0, 0, 0, 0, seconds, 0, time.UTC)
	fmt.Printf("It is %d hours %d minutes.", s.Hour(), s.Minute())
}
