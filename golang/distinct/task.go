// https://app.codility.com/programmers/lessons/6-sorting/distinct/
package distinct

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int{2, 1, 1, 2, 3, 1}, 3)
}

func test(arr []int, expected int) {
	actual := solution(arr)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 30)
	}

	fmt.Printf("act=%v exp=%v arr=%v %v\n", actual, expected, arr, err)
}

func solution(a []int) int {
	ln := len(a)
	m := map[int]int{}
	for _, v := range a {
		if _, ok := m[v]; ok {
			ln--
		}

		m[v]++
	}

	return ln
}
