// https://app.codility.com/programmers/lessons/9-maximum_slice_problem/maxSlice_sum/
package maxslicesum

import (
	"fmt"
)

func Do() {
	test([]int{3, 2, -6, 4, 0}, 5)
	test([]int{-2, 1}, 1)
	test([]int{4, -3, 5, -4, 2}, 6)
	test([]int{1, -3, 5, 2, -3}, 7)
	test([]int{3, 4, 2, -5, -1, -7, 9, 8, -3, -5, -1, -5}, 17)

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
	maxEnding := a[0]
	maxSlice := a[0]

	for i := 1; i < len(a); i++ {
		maxEnding = max(a[i], maxEnding+a[i])
		maxSlice = max(maxEnding, maxSlice)

		// println(i, maxEnding, maxSlice)
	}

	return maxSlice

}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
