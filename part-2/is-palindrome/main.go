package main

import (
	"fmt"
)

func main() {
	var text string

	_, err := fmt.Scan(&text)
	if err != nil {
		return
	}

	textRune := []rune(text)
	var textRuneReverse []rune

	for i := len(textRune) - 1; i >= 0; i-- {
		textRuneReverse = append(textRuneReverse, textRune[i])
	}

	if string(textRune) != string(textRuneReverse) {
		fmt.Print("Нет")
		return
	}

	fmt.Print("Палиндром")
}
