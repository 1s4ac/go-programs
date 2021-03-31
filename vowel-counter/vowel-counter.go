package main

import (
	"fmt"
)

func main() {
	text := "the quick brown fox jumped over the lazy dog"
	fmt.Println(text)

	vowel, consonant := 0, 0
	for _, i := range text {
		switch i {
		case 'a', 'e', 'i', 'o', 'u':
			vowel++
		case ' ':
			continue
		default:
			consonant++
		}
	}
	fmt.Println("vowel count:", vowel)
	fmt.Println("consonant count:", consonant)
}
