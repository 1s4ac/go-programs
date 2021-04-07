package main

import "fmt"

func factorial(n int) int {
	x := 1
	for i := 1; i <= n; i++ {
  //i will increment by 1 until it equals n, 6 & 9 in this example
		x *= i //multiply and assign i
	}
	return x
}

func main() {
	fmt.Println(factorial(6))
	fmt.Println(factorial(9))
}
