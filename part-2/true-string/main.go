package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err == io.EOF {
		text = strings.TrimSuffix(text, "\n")
	}

	textRune := []rune(text)
	firstSymbol := string(textRune[0])
	lastSymbol := string(textRune[len(textRune)-1])

	if firstSymbol != strings.ToUpper(firstSymbol) || lastSymbol != "." { // или unicode.IsUpper
		fmt.Print("Wrong")
		return
	}

	fmt.Print("Right")
}
