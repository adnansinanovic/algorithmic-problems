// https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_profit/
package maxprofit

import (
	"fmt"
)

func Do() {
	test([]int{23171, 21011, 21123, 21366, 21013, 21367}, 356)
	test([]int{100, 90, 92, 93, 91, 94}, 4)

}

func test(s []int, expected int) {
	actual := solution(s)

	err := ""
	if actual != expected {
		err = "!!!!!!!!!!!!!"
	}

	fmt.Printf("actual=%-10v expected=%-10v n=%v %v\n", actual, expected, s, err)
}

func solution(a []int) int {
	if len(a) == 0 {
		return 0
	}
	diffs := make([]int, len(a))
	diffs[0] = 0
	for i := 1; i < len(a); i++ {
		diffs[i] = a[i] - a[i-1]
	}

	maxEnding := diffs[0]
	maxSlice := diffs[0]
	for i := 1; i < len(diffs); i++ {
		maxEnding = max(0, maxEnding+(diffs[i]))
		maxSlice = max(maxSlice, maxEnding)
	}

	return maxSlice

}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
