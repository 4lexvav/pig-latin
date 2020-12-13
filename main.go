package main

import (
	"fmt"
	"os"
	"strings"
)

type PLString string

var plSuffix string = "ay"

func isVowel(char rune) bool {
	for _, v := range "aeiouAEIOU" {
		if char == v {
			return true
		}
	}

	return false
}

func (s PLString) toPigLatin() string {
	var result strings.Builder

	for _, word := range strings.Fields(string(s)) {
		var pl strings.Builder
		var suffix strings.Builder

		isStop := false

		for _, char := range word {
			if isStop || isVowel(char) {
				isStop = true
				pl.WriteRune(char)
			} else {
				suffix.WriteRune(char)
			}
		}

		result.WriteString(pl.String() + suffix.String() + plSuffix + " ")
	}

	return result.String()
}

func main() {
	input := "pig latin"

	if len(os.Args) >= 2 {
		input = strings.Join(os.Args[1:], " ")
	}

	fmt.Printf("Original: %s\n", input)
	fmt.Printf("Pig Latin: %s\n", PLString(input).toPigLatin())
}
