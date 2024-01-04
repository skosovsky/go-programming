package main

import "fmt"

func main() {
	const fullHour int = 60
	const minutesRatio int = 720 / 360
	var degrees, hours, minutes int
	_, err := fmt.Scan(&degrees)
	if err != nil {
		return
	}

	hours = degrees * minutesRatio / fullHour
	minutes = degrees * minutesRatio % fullHour
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
