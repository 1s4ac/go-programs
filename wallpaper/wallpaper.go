package main

import (
	"fmt"
)

func wallPaper(l, w, h float64) string {
	rolls := (((l*h)*2 + (h*w)*2) / 5.2) * 1.15
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	if l == 0.0 || w == 0 || h == 0 {
		return numbers[0]
	}
	return numbers[int(rolls+1)]
}

func main() {
	fmt.Println(wallPaper(6.3, 4.5, 3.29))
	fmt.Println(wallPaper(4.0, 3.5, 3.0))
	fmt.Println(wallPaper(0.0, 3.5, 3.0))
}
//Easy wallpaper - codewars
