// https://app.codility.com/programmers/lessons/12-euclidean_algorithm/chocolates_by_numbers/
package chocolatesbynumber

import (
	"fmt"
)

func Do() {
	test(10, 4, 5)
	test(947853, 4453, 947853)
	test(10, 3, 10)

}

func test(n, m int, expected int) {
	actual := solution(n, m)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v m=%v %v\n", actual, expected, n, m, err)
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}

	return gcd(b, a%b)
}

func solution(N int, M int) int {
	lcm := N * M / gcd(N, M)
	return lcm / M
}
