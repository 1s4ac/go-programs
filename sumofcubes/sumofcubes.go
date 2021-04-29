package main

import "fmt"

func sumCubes(n int) (result int) {
	for i := 1; i <= n; i++ {
		result += (i) * (i) * (i)
	}
	return result
}

func main() {
	fmt.Println(sumCubes(3))
	fmt.Println(sumCubes(4))
}
