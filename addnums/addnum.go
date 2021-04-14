package main

import (
	"fmt"
)

func sum(nums []int) int {
	sum := 0
	for _, i := range nums { //if _ is removed, it will add assign the index (0,1,2,3,4)
		sum += i //in this example
	}
	return sum
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
}
