package main

import "fmt"

func arithmetic(a int, b int, operator string) (result int) {
	switch operator {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		result = a / b
	}
	return result
}
func main() {
	fmt.Println(arithmetic(1, 2, "multiply"))
}
