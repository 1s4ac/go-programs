package main

import (
	"fmt"
)

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func main() {
	fmt.Println(isTriangle(1, 2, 3))
	fmt.Println(isTriangle(10, 2, 9))
}
