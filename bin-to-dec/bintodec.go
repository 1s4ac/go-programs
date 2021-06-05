package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func binToDec(bin []int) (dec int) {
	for a, b := range bin {
		if b == 1 {
			dec1 := math.Pow(2, float64(a))
			dec += int(dec1)
		} else {
			continue
		}
	}
	return dec
}

func intToArr(i int) (a []int) {
	i = reverseInt(i)

	bin3 := strconv.Itoa(i)
	splitArray := strings.Split(bin3, "")
	bin2 := []int{}
	for _, y := range splitArray {
		x, err := strconv.Atoi(y)
		if err != nil {
			panic(err)
		}
		bin2 = append(bin2, x)
	}
	return bin2
}

func reverseInt(n int) (r int) {
	for n > 0 {
		remainder := n % 10
		r *= 10
		r += remainder
		n /= 10
	}
	return r
}

func main() {
	fmt.Println(binToDec(intToArr(111001)))
	fmt.Println(binToDec(intToArr(100011)))
	fmt.Println(binToDec(intToArr(1111)))
	fmt.Println(binToDec(intToArr(1001)))
	fmt.Println(binToDec(intToArr(1110)))
	fmt.Println(binToDec(intToArr(1011)))
}
//unsure why this doesn't work for some binary numbers
