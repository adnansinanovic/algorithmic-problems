// https://app.codility.com/programmers/lessons/4-counting_elements/max_counters/

package maxcounters

import (
	"fmt"
	"strings"
)

func Do() {

	test(5, []int{3, 4, 4, 6, 1, 4, 4}, []int{3, 2, 2, 4, 2})
	test(5, []int{3, 4, 4, 6, 1, 4, 4, 7, 2}, []int{4, 5, 4, 4, 4})

	test(3, []int{1, 3, 5, 2, 1, 1, 1, 2}, []int{4, 3, 1})
	test(3, []int{1, 3, 5, 2, 1, 5, 1, 3}, []int{3, 2, 3})
}

func test(n int, a []int, exp []int) {
	actual := solution(n, a)

	err := ""
	if len(actual) != len(exp) {
		err = strings.Repeat("!", 30) + "gths different len"
	} else {
		for i := 0; i < len(actual); i++ {
			if actual[i] != exp[i] {
				err = strings.Repeat("!", 30) + " mismatch values"
			}

		}
	}

	fmt.Printf("actual=%v, expected=%v, arr=%v %v\n", actual, exp, a, err)

}

func solution(n int, a []int) []int {
	counters := make([]int, n)
	max := 0
	maxUpdate := 0

	for _, v := range a {
		index := v - 1

		if v >= 1 && v <= n {
			if counters[index] < maxUpdate {
				counters[index] = maxUpdate + 1
			} else {
				counters[index]++
			}

			if max < counters[index] {
				max = counters[index]
			}
		} else {
			maxUpdate = max
		}

		// fmt.Printf("%v max=%v\n", counters, max)
	}

	for i := 0; i < n; i++ {
		if counters[i] < maxUpdate {
			counters[i] = maxUpdate
		}
	}

	return counters
}

func solution2(n int, a []int) []int {
	counters := make([]int, n)

	max := 0
	for _, v := range a {
		if v >= 1 && v <= n {
			counters[v-1]++

			if max < counters[v-1] {
				max = counters[v-1]
			}
		} else {
			for ii := 0; ii < n; ii++ {
				counters[ii] = max
			}
		}

		fmt.Printf("%v\n", counters)
	}

	return counters
}
