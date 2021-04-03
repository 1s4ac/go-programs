package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "banana"
	for i, letter := range word {
		if i%2 == 0 {
			letterUpper := strings.ToUpper(string(letter))
			fmt.Printf(letterUpper)
		} else {
			letterLower := strings.ToLower(string(letter))
			fmt.Printf(letterLower)
		}
	}
}
