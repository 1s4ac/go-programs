package main

import "fmt"

func printerError(s string) string {
	counter := 0
	for _, char := range s {
		if string(char) > "m" {
			counter++
		}
	}
	return fmt.Sprintf("%d/%d", counter, len(s))
}

func main() {
	fmt.Println(printerError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))

}
