package main

import (
	"fmt"
	"time"
)

func main() {
	countdown(30, 0)
}

func countdown(start, end int) {
	for i := start; i >= end; i-- {
		if i%2 == 0 {
			fmt.Printf("%d: even\n", i)
		} else {
			fmt.Printf("%d: odd\n", i)
		}
		time.Sleep(500 * time.Millisecond)
	}

	if end > start {
		fmt.Println("starting number greater than ending number")
	}
}
