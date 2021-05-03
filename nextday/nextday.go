package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("the day tomorrow is", t.Weekday()+1)
}
