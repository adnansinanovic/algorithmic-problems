// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
package jumpingontheclouds

import (
	"fmt"
)

func Do() {
	test([]int32{0, 0, 1, 0, 0, 1, 0}, 4)
	test([]int32{0, 0, 0, 0, 1, 0}, 3)
}

func test(s []int32, expected int32) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func solution(c []int32) int32 {
	n := int32(len(c))
	start := int32(0)
	if c[start] == 1 {
		start = 1
	}

	current := start
	jumps := int32(0)
	for current < n {
		if current+2 < n && c[current+2] != 1 {
			jumps++
			current = current + 2

		} else {
			jumps++
			current++
		}

	}

	return jumps - 1
}
