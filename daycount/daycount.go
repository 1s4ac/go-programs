package main

import (
	"fmt"
	"time"
)

func days(year int) (counter int) {
	for month := 1; month <= 12; month++ {
		if time.Date(year, time.Month(month), 13, 0, 0, 0, 0, time.Local).Weekday().String() == "Friday" {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(days(2005))
	fmt.Println(days(2015))
}
