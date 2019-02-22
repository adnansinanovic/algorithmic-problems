// https://www.hackerrank.com/challenges/ctci-big-o/problem
package primality

import (
	"fmt"
	"math"
)

func Do() {
	Test(12, "Not prime")
	Test(2, "Prime")
	Test(5, "Prime")
	Test(7, "Prime")
	Test(31, "Prime")
	Test(33, "Not prime")
	Test(1982, "Not prime")
	Test(9891, "Not prime")
	Test(14582734, "Not prime")
}

func Test(n int32, expected string) {
	actual := primality(n)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, n, err)
}

func primality(n int32) string {
	if n < 2 {
		return "Not prime"
	}

	if n <= 3 {
		return "Prime"
	}

	sqrt := int32(math.Sqrt(float64(n))) + 1
	for i := int32(2); i < sqrt; i++ {
		if n%i == 0 {
			return "Not prime"
		}
	}

	return "Prime"
}
