// https://app.codility.com/programmers/lessons/4-counting_elements/frog_river_one/
package frogriverone

import (
	"fmt"
	"strings"
)

func Do() {
	test(5, []int{1, 3, 1, 4, 2, 3, 5, 4}, 6)
	test(3, []int{1, 1, 1, 2, 1, 1, 3}, 7)
}

func test(x int, a []int, exp int) {
	actual := solution(x, a)

	err := ""
	if actual != exp {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("actual=%v, expected=%v, arr=%v %v\n", actual, exp, a, err)

}

func solution(x int, a []int) int {
	m := map[int]bool{}
	for i, v := range a {
		m[v] = true
		if len(m) == x {
			return i
		}

	}
	return -1
}
