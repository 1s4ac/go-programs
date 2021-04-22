package main

import (
	"fmt"
	"strings"
)

func abbrevname(name string) string {
	split := strings.Split(name, " ")

	first, last := split[0][0], split[1][0]
	return fmt.Sprintf("%s.%s", strings.ToUpper(string(first)), strings.ToUpper(string(last)))
}

func main() {
	fmt.Println(abbrevname("Joe Mama"))
	fmt.Println(abbrevname("Chris Jones"))
	fmt.Println(abbrevname("Joe King"))
	fmt.Println(abbrevname("Banana Man"))
	fmt.Println(abbrevname("lower case"))
}
