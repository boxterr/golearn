package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	x := mySum(17001, 42000, 500000)
	pl("17,000 * 42,000 =", x)
}

func mySum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
		sum++
	}
	return sum
}
