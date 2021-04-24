package main

import "fmt"

func findOdd(seq []int) int {
	values := make(map[int]int)

	for _, num := range seq {
		values[num] = values[num] + 1

	}
	for k, v := range values {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(findOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
	fmt.Println(findOdd([]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}))
}
