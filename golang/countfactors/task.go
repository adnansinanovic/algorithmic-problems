// https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/count_factors/
package countfactors

import (
	"fmt"
)

func Do() {
	test(0, 0)
	test(1, 1)
	test(9, 3)
	test(24, 8)
	test(25, 3)

	test(32, 6)
	test(144, 15)
	test(1001, 8)
}

func test(s int, expected int) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func solution(a int) int {
	if a == 1 {
		return 1
	}
	count := 0
	i := 1
	for i*i < a {
		if a%i == 0 {
			count += 2
		}

		i++
		if i*i == a {
			count++
		}
	}

	return count
}

func solution2(a int) int {
	// brute force
	count := 0
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			count++
		}
	}
	return count
}
