package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	text, err := ioutil.ReadFile("sentence.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(text))

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
