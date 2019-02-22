// https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/min_perimeter_rectangle/
package minperimeterrectangle

import (
	"fmt"
	"math"
)

func Do() {
	test(1, 4)
	test(30, 22)
}

func test(s int, expected int) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func solution(n int) int {
	min := math.MaxInt64
	i := 1
	for i*i <= n {
		if n%i == 0 {
			a := i
			b := n / i
			perimeter := 2 * (a + b)

			if perimeter < min {
				min = perimeter
			}
		}

		i++
	}

	return min
}
