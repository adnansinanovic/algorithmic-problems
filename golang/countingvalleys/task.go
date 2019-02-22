// https://www.hackerrank.com/challenges/counting-valleys/problem
package countingvalleys

import (
	"fmt"
)

func Do() {
	test(7, "DDUUUUDD", 1)
	test(8, "UDDDUDUU", 1)
}

func test(n int32, s string, expected int32) {
	actual := countingValleys(n, s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func countingValleys(n int32, s string) int32 {
	d := byte("D"[0])

	level := 0
	valleyCount := int32(0)
	for i := int32(0); i < n; i++ {
		if s[i] == d {
			level--
		} else {
			level++

			if level == 0 {
				valleyCount++
			}
		}
	}

	return valleyCount
}
