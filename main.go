package learningGo

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	a := 17001
	b := 42001
	c := 50001
	x := mySum(a, b, c)
	pf("%v + %v + %v = %v\n", a, b, c, x)
}

func mySum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
