package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := n; i >= 1; i-- {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))
}
