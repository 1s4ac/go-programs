package main

import "fmt"

func positiveSum(numbers []int) (sum int) {
	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}
	return sum
}

func main() {
	fmt.Println(positiveSum([]int{1, 2, 3, 4, -5, 6, -7}))
}
