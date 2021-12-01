package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) == 0 {
		return ""
	}

	openBracketIndex := strings.Index(str, "(")
	if openBracketIndex < 0 {
		return ""
	}

	closeBracketIndex := strings.Index(str, ")")
	if closeBracketIndex < 0 {
		return ""
	}

	return str[openBracketIndex+1 : closeBracketIndex]
}

func main() {
	fmt.Println(findFirstStringInBracket("Hai (Stockbit) Team!"))
}
