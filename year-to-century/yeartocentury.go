package main

import "fmt"

func century(year int) int {
	if year%100 != 0 {
		return year/100 + 1
	}
	return year / 100

}

func main() {

	fmt.Println(century(1990))
	fmt.Println(century(1705))
	fmt.Println(century(1900))
	fmt.Println(century(1601))
	fmt.Println(century(2000))
	fmt.Println(century(89))
	fmt.Println(century(2021))
}
