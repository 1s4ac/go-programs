package main

import (
	"fmt"
	"time"
)

func leapYear(year int) string {
	if time.Date(year, time.Month(12), 31, 0, 0, 0, 0, time.Local).YearDay() > 365 {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Println(leapYear(2005))
	fmt.Println(leapYear(2004))
	fmt.Println(leapYear(1948))
	fmt.Println(leapYear(2020))
	fmt.Println(leapYear(2021))
}
